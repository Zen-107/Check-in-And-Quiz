-- MySQL Schema for Star Exhibition System
-- Database: star_exhibition

CREATE DATABASE IF NOT EXISTS star_exhibition CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE star_exhibition;

-- Table: participants (ผู้เข้าร่วมงาน)
CREATE TABLE IF NOT EXISTS participants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    participant_type ENUM('student', 'staff', 'external') NOT NULL COMMENT 'ประเภท: นิสิต, บุคลากร, บุคคลภายนอก',
    
    -- ข้อมูลสำหรับนิสิต
    student_id VARCHAR(20) DEFAULT NULL COMMENT 'รหัสนิสิต',
    student_name VARCHAR(100) DEFAULT NULL COMMENT 'ชื่อนิสิต',
    faculty VARCHAR(100) DEFAULT NULL COMMENT 'คณะ',
    department VARCHAR(100) DEFAULT NULL COMMENT 'สาขาวิชา',
    year_level INT DEFAULT NULL COMMENT 'ชั้นปี',
    
    -- ข้อมูลสำหรับบุคลากร
    staff_id VARCHAR(20) DEFAULT NULL COMMENT 'รหัสบุคลากร',
    staff_name VARCHAR(100) DEFAULT NULL COMMENT 'ชื่อบุคลากร',
    division VARCHAR(100) DEFAULT NULL COMMENT 'หน่วยงาน/กอง',
    position VARCHAR(100) DEFAULT NULL COMMENT 'ตำแหน่ง',
    
    -- ข้อมูลสำหรับบุคคลภายนอก
    external_name VARCHAR(100) DEFAULT NULL COMMENT 'ชื่อ-นามสกุล',
    external_organization VARCHAR(100) DEFAULT NULL COMMENT 'สังกัด/องค์กร',
    external_phone VARCHAR(20) DEFAULT NULL COMMENT 'เบอร์โทรศัพท์',
    
    -- ข้อมูลกลาง
    checked_in_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'เวลาเช็คอิน',
    personality_result VARCHAR(50) DEFAULT NULL COMMENT 'ผลลัพธ์บุคลิกภาพดวงดาว',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_student_id (student_id),
    INDEX idx_staff_id (staff_id),
    INDEX idx_checked_in_at (checked_in_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: quiz_results (ผลแบบทดสอบบุคลิกภาพ)
CREATE TABLE IF NOT EXISTS quiz_results (
    id INT AUTO_INCREMENT PRIMARY KEY,
    participant_id INT NOT NULL,
    
    -- ผลลัพธ์ดวงดาว
    star_type VARCHAR(50) NOT NULL COMMENT 'ประเภทดวงดาว เช่น Red Giant, Blue Supergiant',
    star_identity VARCHAR(100) NOT NULL COMMENT 'ตัวตนดวงดาว เช่น The Leader, The Dreamer',
    star_description TEXT COMMENT 'คำอธิบายลักษณะนิสัย',
    star_traits JSON COMMENT 'คุณสมบัติเด่น (array)',
    
    -- คะแนนแต่ละด้าน (ถ้ามี)
    score_leadership INT DEFAULT 0,
    score_creativity INT DEFAULT 0,
    score_empathy INT DEFAULT 0,
    score_analytical INT DEFAULT 0,
    
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (participant_id) REFERENCES participants(id) ON DELETE CASCADE,
    INDEX idx_participant_id (participant_id),
    INDEX idx_star_type (star_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Table: statistics (สำหรับนับจำนวน Real-time)
CREATE TABLE IF NOT EXISTS statistics (
    id INT AUTO_INCREMENT PRIMARY KEY,
    stat_key VARCHAR(50) UNIQUE NOT NULL,
    stat_value INT DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert initial statistics
INSERT INTO statistics (stat_key, stat_value) VALUES 
('total_participants', 0),
('total_students', 0),
('total_staff', 0),
('total_external', 0)
ON DUPLICATE KEY UPDATE stat_value = stat_value;
