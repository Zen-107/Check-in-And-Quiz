# 🐳 Docker Commands สำหรับ Check-In-And-Quiz

## รัน MySQL ด้วย Docker

### เริ่มทำงาน Database
```bash
docker-compose up -d
```

### ตรวจสอบสถานะ
```bash
docker-compose ps
```

### ดู Logs
```bash
# ดู logs ทั้งหมด
docker-compose logs

# ดู logs เฉพาะ database และตาม real-time
docker-compose logs -f db
```

### หยุดการทำงาน (ข้อมูลยังอยู่)
```bash
docker-compose down
```
> ⚠️ คำสั่งนี้ **ไม่ลบข้อมูล** เพราะข้อมูลถูกเก็บใน Volume

### ลบทุกอย่างรวมถึง Volume (ข้อมูลหาย!)
```bash
docker-compose down -v
```
> ⚠️ **คำเตือน:** คำสั่งนี้จะลบ Volume และข้อมูลทั้งหมด!

## การจัดการข้อมูล

### เข้าถึง MySQL ใน Container
```bash
docker exec -it checkin_quiz_db mysql -u quizuser -pquizpassword checkin_quiz
```

### Backup Database
```bash
docker exec checkin_quiz_db mysqldump -u quizuser -pquizpassword checkin_quiz > backup.sql
```

### Restore Database
```bash
cat backup.sql | docker exec -i checkin_quiz_db mysql -u quizuser -pquizpassword checkin_quiz
```

### ดูข้อมูลใน Volume
```bash
# ดูว่า Volume เก็บอยู่ที่ไหน
docker volume inspect checkin_quiz_mysql_data

# หรือดูผ่าน container
docker run --rm -v checkin_quiz_mysql_data:/var/lib/mysql alpine ls -la /var/lib/mysql
```

## Environment Variables ที่ใช้

| Variable | ค่าเริ่มต้น | คำอธิบาย |
|----------|-----------|----------|
| MYSQL_ROOT_PASSWORD | rootpassword | รหัสผ่าน root |
| MYSQL_DATABASE | checkin_quiz | ชื่อ database |
| MYSQL_USER | quizuser | ผู้ใช้งานสำหรับ app |
| MYSQL_PASSWORD | quizpassword | รหัสผ่านผู้ใช้ |

## เชื่อมต่อจากภายนอก

Database พร้อมใช้งานที่:
- **Host:** localhost
- **Port:** 3306
- **User:** quizuser
- **Password:** quizpassword
- **Database:** checkin_quiz

## Troubleshooting

### Port 3306 ถูกใช้งานอยู่แล้ว
แก้ไข `docker-compose.yml` เปลี่ยน port mapping:
```yaml
ports:
  - "3307:3306"  # ใช้ port 3307 แทน
```

แล้วอย่าลืมแก้ `DB_PORT` ใน environment variables เป็น 3307

### Container ไม่เริ่มทำงาน
```bash
# ดู logs เพื่อหาสาเหตุ
docker-compose logs db

# ลอง rebuild
docker-compose down
docker-compose up -d --build
```

### Reset Database ทั้งหมด
```bash
docker-compose down -v
docker-compose up -d
```
