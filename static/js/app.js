// Create animated stars background
function createStars() {
    const starsBg = document.querySelector('.stars-bg');
    if (!starsBg) return;

    // Create regular stars
    for (let i = 0; i < 100; i++) {
        const star = document.createElement('div');
        star.className = 'star';
        
        const size = Math.random() * 3 + 1;
        star.style.width = `${size}px`;
        star.style.height = `${size}px`;
        star.style.left = `${Math.random() * 100}%`;
        star.style.top = `${Math.random() * 100}%`;
        star.style.setProperty('--duration', `${Math.random() * 3 + 2}s`);
        star.style.setProperty('--opacity', `${Math.random() * 0.7 + 0.3}`);
        star.style.animationDelay = `${Math.random() * 2}s`;
        
        starsBg.appendChild(star);
    }

    // Create shooting stars
    for (let i = 0; i < 3; i++) {
        const shootingStar = document.createElement('div');
        shootingStar.className = 'shooting-star';
        shootingStar.style.left = `${Math.random() * 100}%`;
        shootingStar.style.top = `${Math.random() * 50}%`;
        shootingStar.style.animationDelay = `${Math.random() * 5 + i * 3}s`;
        
        starsBg.appendChild(shootingStar);
    }
}

// Update check-in counter in real-time
async function updateCheckInCount() {
    try {
        const response = await fetch('/api/checkin/count');
        const data = await response.json();
        const countElement = document.getElementById('checkin-count');
        if (countElement) {
            countElement.textContent = data.count.toLocaleString();
        }
    } catch (error) {
        console.error('Error fetching check-in count:', error);
    }
}

// Initialize page
document.addEventListener('DOMContentLoaded', () => {
    createStars();
    updateCheckInCount();
    toggleUserTypeFields();
    
    // Update counter every 5 seconds
    setInterval(updateCheckInCount, 5000);
});

// Form submission handler for check-in
async function submitCheckIn(event) {
    event.preventDefault();
    
    const form = event.target;
    const submitBtn = form.querySelector('button[type="submit"]');
    const originalText = submitBtn.innerHTML;

    const data = collectCheckInData(form);
    const validationError = validateCheckInData(data);
    if (validationError) {
        showNotification(validationError, 'error');
        return;
    }
    
    // Show loading state
    submitBtn.disabled = true;
    submitBtn.innerHTML = '<span class="spinner"></span> กำลังบันทึก...';
    
    try {
        const response = await fetch('/api/checkin', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        });
        
        let result = {};
        try {
            result = await response.json();
        } catch (error) {
            result = {};
        }
        
        if (response.ok && result.success) {
            showNotification(result.message, 'success');
            form.reset();
            toggleUserTypeFields();
            updateCheckInCount();
        } else {
            showNotification(result.message || result.error || 'เกิดข้อผิดพลาด กรุณาลองใหม่อีกครั้ง', 'error');
        }
    } catch (error) {
        showNotification('เกิดข้อผิดพลาดในการเชื่อมต่อ กรุณาลองใหม่อีกครั้ง', 'error');
    } finally {
        submitBtn.disabled = false;
        submitBtn.innerHTML = originalText;
    }
}

// Form submission handler for quiz
async function submitQuiz(event) {
    event.preventDefault();
    
    const form = event.target;
    const submitBtn = form.querySelector('button[type="submit"]');
    const originalText = submitBtn.innerHTML;
    
    // Validate all questions are answered
    const selectedOptions = form.querySelectorAll('.quiz-option.selected');
    if (selectedOptions.length < 5) {
        showNotification('กรุณาตอบคำถามให้ครบทุกข้อ', 'error');
        return;
    }
    
    // Show loading state
    submitBtn.disabled = true;
    submitBtn.innerHTML = '<span class="spinner"></span> กำลังคำนวณผล...';
    
    const answers = {};
    form.querySelectorAll('.quiz-question').forEach((question, index) => {
        const selected = question.querySelector('.quiz-option.selected');
        if (selected) {
            answers[`q${index + 1}`] = parseInt(selected.dataset.value);
        }
    });
    
    const name = form.querySelector('input[name="name"]').value;
    
    try {
        const response = await fetch('/api/quiz', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, answers }),
        });
        
        const result = await response.json();
        
        if (response.ok && result.success) {
            window.location.href = `/result?type=${result.type}&name=${encodeURIComponent(result.name)}`;
        } else {
            showNotification(result.message || 'เกิดข้อผิดพลาด กรุณาลองใหม่อีกครั้ง', 'error');
        }
    } catch (error) {
        showNotification('เกิดข้อผิดพลาดในการเชื่อมต่อ กรุณาลองใหม่อีกครั้ง', 'error');
    } finally {
        submitBtn.disabled = false;
        submitBtn.innerHTML = originalText;
    }
}

// Show notification
function showNotification(message, type) {
    const alertDiv = document.createElement('div');
    alertDiv.className = `alert alert-${type}`;
    alertDiv.textContent = message;
    
    const container = document.querySelector('.container');
    const firstElement = container.firstChild;
    container.insertBefore(alertDiv, firstElement);
    
    setTimeout(() => {
        alertDiv.remove();
    }, 5000);
}

// Quiz option selection
function selectOption(option) {
    const parent = option.parentElement;
    parent.querySelectorAll('.quiz-option').forEach(opt => opt.classList.remove('selected'));
    option.classList.add('selected');
}

// Toggle user type fields
function toggleUserTypeFields() {
    const userType = document.getElementById('user-type');
    if (!userType) return;

    const sections = {
        student: document.getElementById('student-fields'),
        staff: document.getElementById('staff-fields'),
        external: document.getElementById('external-fields'),
    };

    const requiredByType = {
        student: ['staff_id_student', 'faculty_student'],
        staff: ['position_staff', 'faculty_staff'],
        external: ['occupation_external'],
    };

    Object.entries(sections).forEach(([type, section]) => {
        if (!section) return;

        const isActive = userType.value === type;
        section.classList.toggle('hidden', !isActive);

        section.querySelectorAll('input, select, textarea').forEach((input) => {
            input.disabled = !isActive;
            input.required = false;
        });
    });

    (requiredByType[userType.value] || []).forEach((id) => {
        const input = document.getElementById(id);
        if (input) {
            input.required = true;
        }
    });
}

function collectCheckInData(form) {
    const userType = form.querySelector('[name="user_type"]').value;
    const data = {
        user_type: userType,
        name: form.querySelector('[name="name"]').value.trim(),
    };

    if (userType === 'student') {
        data.staff_id = form.querySelector('#staff_id_student').value.trim();
        data.faculty = form.querySelector('#faculty_student').value.trim();
    } else if (userType === 'staff') {
        data.staff_id = form.querySelector('#staff_id_staff').value.trim();
        data.position = form.querySelector('#position_staff').value.trim();
        data.faculty = form.querySelector('#faculty_staff').value.trim();
    } else if (userType === 'external') {
        data.occupation = form.querySelector('#occupation_external').value.trim();
    }

    return data;
}

function validateCheckInData(data) {
    if (!data.user_type) {
        return 'กรุณาเลือกประเภทผู้เข้าร่วม';
    }
    if (!data.name) {
        return 'กรุณากรอกชื่อ-นามสกุล';
    }

    if (data.user_type === 'student') {
        if (!data.staff_id) return 'กรุณากรอกรหัสนิสิต';
        if (!data.faculty) return 'กรุณากรอกคณะ';
    } else if (data.user_type === 'staff') {
        if (!data.position) return 'กรุณากรอกตำแหน่ง';
        if (!data.faculty) return 'กรุณากรอกคณะ/หน่วยงาน';
    } else if (data.user_type === 'external') {
        if (!data.occupation) return 'กรุณากรอกอาชีพ';
    } else {
        return 'ประเภทผู้เข้าร่วมไม่ถูกต้อง';
    }

    return null;
}

// Share result
function shareResult() {
    if (navigator.share) {
        navigator.share({
            title: 'ผลลัพธ์ดวงดาวของฉัน',
            text: `ฉันคือ ${document.querySelector('.result-title')?.textContent}`,
            url: window.location.href,
        }).catch(console.error);
    } else {
        // Fallback: copy URL to clipboard
        navigator.clipboard.writeText(window.location.href).then(() => {
            showNotification('คัดลอกลิงก์แล้ว!', 'success');
        });
    }
}

// Download result as image (simple implementation)
function downloadResult() {
    showNotification('กำลังดาวน์โหลด...', 'success');
    // In a real implementation, you would use html2canvas or similar library
    setTimeout(() => {
        showNotification('แคปหน้าจอเพื่อบันทึกผลลัพธ์!', 'success');
    }, 1000);
}
