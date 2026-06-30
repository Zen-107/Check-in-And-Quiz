-- init.sql: สร้างตารางสำหรับระบบ Check-In และ Quiz
-- ไฟล์นี้จะถูกทำงานอัตโนมัติเมื่อ MySQL Container เริ่มต้นครั้งแรก

USE checkin_quiz;

-- ตารางเก็บข้อมูลการเช็คชื่อ
CREATE TABLE IF NOT EXISTS attendances (
    id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL COMMENT 'ชื่อ-นามสกุล',
    user_type ENUM('student', 'staff', 'external') NOT NULL COMMENT 'ประเภท: student=นิสิต, staff=บุคลากร, external=บุคคลภายนอก',
    
    -- ข้อมูลสำหรับนิสิต
    student_id VARCHAR(50) NULL COMMENT 'รหัสนิสิต',
    faculty VARCHAR(255) NULL COMMENT 'คณะ',
    
    -- ข้อมูลสำหรับบุคลากร
    position VARCHAR(255) NULL COMMENT 'ตำแหน่ง',
    
    -- ข้อมูลสำหรับบุคคลภายนอก
    occupation VARCHAR(255) NULL COMMENT 'อาชีพ',
    
    check_in_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'เวลาเช็คอิน',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    INDEX idx_user_type (user_type),
    INDEX idx_check_in_time (check_in_time)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='ตารางเก็บข้อมูลการเช็คชื่อ';

-- ตารางเก็บข้อมูลผลแบบทดสอบดวงดาว
CREATE TABLE IF NOT EXISTS quiz_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL COMMENT 'ชื่อ-นามสกุลผู้ทำแบบทดสอบ',
    user_type ENUM('student', 'staff', 'external') NOT NULL COMMENT 'ประเภทผู้ใช้',
    
    -- ข้อมูลระบุตัวตน (แล้วแต่ประเภท)
    student_id VARCHAR(50) NULL COMMENT 'รหัสนิสิต',
    faculty VARCHAR(255) NULL COMMENT 'คณะ',
    position VARCHAR(255) NULL COMMENT 'ตำแหน่ง',
    occupation VARCHAR(255) NULL COMMENT 'อาชีพ',
    
    -- ผลแบบทดสอบ
    answers JSON NOT NULL COMMENT 'คำตอบทั้งหมด (JSON array)',
    result_type VARCHAR(100) NOT NULL COMMENT 'ประเภทดวงดาวที่ได้',
    result_description TEXT NULL COMMENT 'คำอธิบายผลลัพธ์',
    
    submitted_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT 'เวลาที่ส่งแบบทดสอบ',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    INDEX idx_result_type (result_type),
    INDEX idx_submitted_at (submitted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='ตารางเก็บข้อมูลผลแบบทดสอบดวงดาว';

-- แทรกข้อมูลตัวอย่าง (Optional)
-- INSERT INTO attendances (full_name, user_type, student_id, faculty) VALUES 
-- ('ทดสอบ ระบบ', 'student', '670000001', 'วิศวกรรมศาสตร์');

SELECT 'Database initialized successfully!' AS status;
