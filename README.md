# Check-In-And-Quiz Web Application

Web Application สำหรับระบบเช็คชื่อและแบบทดสอบบุคลิกภาพดวงดาว ด้วยธีมจักรวาลและดวงดาว

## 🌟 ฟีเจอร์หลัก

### 1. ระบบเช็คชื่อเข้างาน (Attendance System)
- รองรับผู้ใช้ 3 ประเภท: นิสิต, บุคลากร, และบุคคลภายนอก
- เก็บข้อมูลตามประเภทผู้ใช้:
  - **นิสิต**: ชื่อ-นามสกุล, รหัสนิสิต, คณะ
  - **บุคลากร**: ชื่อ-นามสกุล, รหัสพนักงาน, ตำแหน่ง, คณะ/หน่วยงาน
  - **บุคคลภายนอก**: ชื่อ-นามสกุล, อาชีพ
- แสดงจำนวนผู้เช็คชื่อแบบ Real-time
- มี QR Code สำหรับสแกนเช็คอิน

### 2. แบบทดสอบบุคลิกภาพดวงดาว (Star Personality Quiz)
- แบบทดสอบ 5 ข้อ แนว MBTI ย่อยง่าย
- ผลลัพธ์ 5 แบบ:
  - 💕 **ดาวศุกร์ (Venus)**: นักประนีประนอม อ่อนโยน มีศิลปะ
  - ⚔️ **ดาวอังคาร (Mars)**: นักรบกล้าหาญ พลังงานล้นเหลือ
  - 🌊 **ดาวเนปจูน (Neptune)**: ช่างจินตนาการ สร้างสรรค์ ลึกลับ
  - 🪐 **ดาวพฤหัสบดี (Jupiter)**: ผู้นำแห่งปัญญา ฉลาด มองการณ์ไกล
  - 💍 **ดาวเสาร์ (Saturn)**:ผู้มีวินัย รับผิดชอบ มั่นคง
- หน้าผลลัพธ์สวยงาม เหมาะสำหรับแคปหน้าจอแชร์

## 🛠️ Tech Stack

- **Frontend**: HTML, CSS, JavaScript (Tailwind CSS style)
- **Backend**: Golang
- **Database**: MySQL
- **Theme**: Fairytale, Celestial, Galaxy (โทนสีน้ำเงินเข้ม, ม่วง, ประกายดาว)

## 📁 โครงสร้างโปรเจกต์

```
check-in-and-quiz/
├── cmd/
│   └── server/
│       └── main.go              # Entry point ของ application
├── internal/
│   ├── database/
│   │   └── database.go          # Database connection & migration
│   ├── handlers/
│   │   └── handlers.go          # HTTP handlers
│   └── models/
│       └── models.go            # Data models
├── static/
│   ├── css/
│   │   └── styles.css           # Custom CSS styles
│   └── js/
│       └── app.js               # Frontend JavaScript
├── templates/
│   ├── home.html                # หน้าแรก
│   ├── checkin.html             # หน้าเช็คชื่อ
│   ├── quiz.html                # หน้าแบบทดสอบ
│   └── result.html              # หน้าผลลัพธ์
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies checksum
└── server                       # Compiled binary
```

## 🗄️ Database Schema

### ตาราง `checkins`
```sql
CREATE TABLE checkins (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_type ENUM('student', 'staff', 'external') NOT NULL,
    name VARCHAR(255) NOT NULL,
    staff_id VARCHAR(100),
    faculty VARCHAR(255),
    position VARCHAR(255),
    occupation VARCHAR(255),
    checkin_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### ตาราง `quiz_results`
```sql
CREATE TABLE quiz_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    answers JSON NOT NULL,
    result_type VARCHAR(100) NOT NULL,
    result_description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🚀 การติดตั้งและรันโปรแกรม

### ความต้องการระบบ
- Go 1.19 หรือสูงกว่า
- MySQL 5.7 หรือสูงกว่า
- Web Browser ที่ทันสมัย

### 1. ติดตั้ง Dependencies
```bash
cd /workspace
go mod tidy
```

### 2. ตั้งค่า Database
สร้าง Database ใน MySQL:
```sql
CREATE DATABASE checkin_quiz CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. ตั้งค่า Environment Variables
```bash
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=your_password
export DB_NAME=checkin_quiz
export PORT=8080
```

### 4. รันโปรแกรม
```bash
# Build
go build -o server ./cmd/server

# Run
./server
```

หรือใช้คำสั่งเดียว:
```bash
go run ./cmd/server/main.go
```

### 5. เข้าใช้งาน
เปิด Web Browser ไปที่: `http://localhost:8080`

## 🎨 Design Features

### Theme & Design
- **พื้นหลัง**: Gradient สีน้ำเงินเข้ม-ม่วง พร้อมดาวระยิบระยับ
- **Animation**: ดาวกระพริบ, ดาวตกไหลผ่าน
- **Glassmorphism**: การ์ดโปร่งใสเบลอพื้นหลัง
- **Colors**: 
  - Primary: #667eea → #764ba2 (Purple gradient)
  - Secondary: #f093fb → #f5576c (Pink gradient)
  - Background: #0f0c29 → #302b63 → #24243e (Dark space gradient)

### UI Components
- ปุ่ม Gradient สวยงามพร้อม Hover effect
- ฟอร์ม Input สไตล์ Glassmorphism
- Counter แสดงจำนวน Real-time
- Result Card พร้อม Icon และ Traits badges
- Responsive Design รองรับ Mobile

## 📱 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | หน้าแรก |
| GET | `/checkin` | หน้าเช็คชื่อ |
| POST | `/api/checkin` | ส่งข้อมูลเช็คชื่อ |
| GET | `/api/checkin/count` | ดึงจำนวนผู้เช็คชื่อ |
| GET | `/quiz` | หน้าแบบทดสอบ |
| POST | `/api/quiz` | ส่งคำตอบแบบทดสอบ |
| GET | `/result?type=&name=` | หน้าผลลัพธ์ |

## 🌟 ตัวอย่างการใช้งาน

### 1. เช็คชื่อเข้างาน
```javascript
POST /api/checkin
Content-Type: application/json

{
    "user_type": "student",
    "name": "สมชาย ใจดี",
    "staff_id": "6510001",
    "faculty": "วิศวกรรมศาสตร์"
}

Response:
{
    "success": true,
    "id": 1,
    "message": "เช็คชื่อสำเร็จ! ยินดีต้อนรับสู่จักรวาลของเรา"
}
```

### 2. ทำแบบทดสอบ
```javascript
POST /api/quiz
Content-Type: application/json

{
    "name": "สมหญิง รักเรียน",
    "answers": {
        "q1": 1,
        "q2": 3,
        "q3": 4,
        "q4": 2,
        "q5": 5
    }
}

Response:
{
    "success": true,
    "type": "neptune",
    "name": "สมหญิง รักเรียน"
}
```

## 🔧 การพัฒนาต่อ

### เพิ่มฟีเจอร์แนะนำ
- [ ] QR Code Generator จริงสำหรับเช็คอิน
- [ ] Export ผลลัพธ์เป็นรูปภาพ (html2canvas)
- [ ] ระบบ Login/User Management
- [ ] Dashboard สำหรับ Admin ดูสถิติ
- [ ] Email Notification
- [ ] Social Media Sharing ที่สมบูรณ์

### การปรับปรุง UI/UX
- [ ] Loading Skeleton
- [ ] Toast Notifications ที่ดีขึ้น
- [ ] Animation เพิ่มเติม
- [ ] Dark/Light Mode

## 📝 License

MIT License

## 👨‍💻 Author

Full-Stack Developer

---

🌟 **ยินดีต้อนรับสู่จักรวาลดวงดาวของคุณ!** ✨
