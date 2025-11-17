package controllers

import (
	"net/http"
	"skillup-backend/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateQuiz(c *gin.Context) {
	userId := c.GetString("user_id")
	var body struct {
		TopicID        *string `json:"topic_id"`
		QuizType       string  `json:"quiz_type"`
		Questions      string  `json:"questions"` // JSON
		Score          float64 `json:"score"`
		TotalQuestions int     `json:"total_questions"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	now := time.Now()
	q := db.Quiz{
		ID:             uuid.NewString(),
		UserID:         userId,
		TopicID:        body.TopicID,
		QuizType:       body.QuizType,
		Questions:      body.Questions,
		Score:          body.Score,
		TotalQuestions: body.TotalQuestions,
		AttemptedAt:    &now,
	}
	db.DB.Create(&q)
	c.JSON(http.StatusOK, q)
}

func GetQuizzes(c *gin.Context) {
	userId := c.GetString("user_id")
	var qs []db.Quiz
	db.DB.Where("user_id = ?", userId).Order("attempted_at desc").Limit(50).Find(&qs)
	c.JSON(http.StatusOK, qs)
}
