package handlers

import (
	"check-in-and-quiz/internal/database"
	"check-in-and-quiz/internal/models"
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"time"
)

// HomeHandler serves the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

// CheckInHandler serves the check-in page
func CheckInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/checkin.html"))
		tmpl.Execute(w, nil)
		return
	}
}

// QuizHandler serves the quiz page
func QuizHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/quiz.html"))
		tmpl.Execute(w, nil)
		return
	}
}

// ResultHandler serves the result page
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	resultType := r.URL.Query().Get("type")
	name := r.URL.Query().Get("name")
	
	personality := getPersonalityByType(resultType)
	
	data := map[string]interface{}{
		"Name":        name,
		"Personality": personality,
	}
	
	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	tmpl.Execute(w, data)
}

// GetCheckInCountHandler returns the total number of check-ins
func GetCheckInCountHandler(w http.ResponseWriter, r *http.Request) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM checkins").Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"count": count})
}

// SubmitCheckInHandler handles check-in submission
func SubmitCheckInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CheckInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Insert into database
	query := `INSERT INTO checkins (user_type, name, staff_id, faculty, position, occupation, checkin_time) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`
	
	var lastInsertId int64
	var execErr error
	var result sql.Result
	switch req.UserType {
	case "student":
		result, execErr = database.DB.Exec(query, req.UserType, req.Name, req.StaffID, req.Faculty, "", "", time.Now())
	case "staff":
		result, execErr = database.DB.Exec(query, req.UserType, req.Name, req.StaffID, req.Faculty, req.Position, "", time.Now())
	case "external":
		result, execErr = database.DB.Exec(query, req.UserType, req.Name, "", "", "", req.Occupation, time.Now())
	default:
		http.Error(w, "Invalid user type", http.StatusBadRequest)
		return
	}

	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
		return
	}

	lastInsertId, _ = result.LastInsertId()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"id":      lastInsertId,
		"message": "เช็คชื่อสำเร็จ! ยินดีต้อนรับสู่จักรวาลของเรา",
	})
}

// SubmitQuizHandler handles quiz submission
func SubmitQuizHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.QuizRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate result
	resultType := calculateQuizResult(req.Answers)
	personality := getPersonalityByType(resultType)

	// Save to database
	answersJSON, _ := json.Marshal(req.Answers)
	query := `INSERT INTO quiz_results (name, answers, result_type, result_description) VALUES (?, ?, ?, ?)`
	_, err := database.DB.Exec(query, req.Name, string(answersJSON), resultType, personality.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"type":    resultType,
		"name":    req.Name,
	})
}

// calculateQuizResult calculates the personality type based on answers
func calculateQuizResult(answers map[string]int) string {
	// Simple scoring system
	// A: Venus (นักประนีประนอม)
	// B: Mars (นักรบกล้าหาญ)
	// C: Neptune (ช่างจินตนาการ)
	// D: Jupiter (ผู้นำแห่งปัญญา)
	// E: Saturn (ผู้มีวินัย)
	
	scores := map[string]int{
		"venus":     0,
		"mars":      0,
		"neptune":   0,
		"jupiter":   0,
		"saturn":    0,
	}

	for _, answer := range answers {
		switch answer {
		case 1:
			scores["venus"]++
		case 2:
			scores["mars"]++
		case 3:
			scores["neptune"]++
		case 4:
			scores["jupiter"]++
		case 5:
			scores["saturn"]++
		}
	}

	// Find highest score
	maxScore := 0
	result := "venus"
	for key, score := range scores {
		if score > maxScore {
			maxScore = score
			result = key
		}
	}

	return result
}

// getPersonalityByType returns the personality details
func getPersonalityByType(personalityType string) models.StarPersonality {
	personalities := map[string]models.StarPersonality{
		"venus": {
			Type:        "venus",
			Name:        "ดาวศุกร์ (Venus)",
			Description: "นักประนีประนอมผู้มีความอ่อนโยน คุณคือแสงสว่างที่อบอุ่นในจักรวาล",
			Traits:      []string{"อ่อนโยน", "มีศิลปะ", "รักความสงบ", "เห็นอกเห็นใจ"},
			Color:       "#FF69B4",
			Icon:        "💕",
		},
		"mars": {
			Type:        "mars",
			Name:        "ดาวอังคาร (Mars)",
			Description: "นักรบผู้กล้าหาญ พลังงานล้นเหลือ ไม่ยอมแพ้ต่ออุปสรรค",
			Traits:      []string{"กล้าหาญ", "กระตือรือร้น", "มุ่งมั่น", "เป็นผู้นำ"},
			Color:       "#FF4500",
			Icon:        "⚔️",
		},
		"neptune": {
			Type:        "neptune",
			Name:        "ดาวเนปจูน (Neptune)",
			Description: "ผู้ช่างฝันที่มีโลกส่วนตัวสูง จินตนาการไร้ขีดจำกัด",
			Traits:      []string{"ช่างจินตนาการ", "สร้างสรรค์", "ลึกลับ", "มีสัญชาตญาณ"},
			Color:       "#4169E1",
			Icon:        "🌊",
		},
		"jupiter": {
			Type:        "jupiter",
			Name:        "ดาวพฤหัสบดี (Jupiter)",
			Description: "ผู้นำแห่งปัญญา ความรู้ และความเจริญรุ่งเรือง",
			Traits:      []string{"ฉลาด", "มองการณ์ไกล", "เอื้อเฟื้อ", "มีคุณธรรม"},
			Color:       "#FFD700",
			Icon:        "🪐",
		},
		"saturn": {
			Type:        "saturn",
			Name:        "ดาวเสาร์ (Saturn)",
			Description: "ผู้มีวินัยและความรับผิดชอบ มั่นคงดั่งขุนเขา",
			Traits:      []string{"มีวินัย", "รับผิดชอบ", "อดทน", "จริงจัง"},
			Color:       "#C0C0C0",
			Icon:        "💍",
		},
	}

	return personalities[strings.ToLower(personalityType)]
}
