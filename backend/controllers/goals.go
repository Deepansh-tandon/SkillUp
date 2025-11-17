package controllers

import (
	"net/http"
	"skillup-backend/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetGoals(c *gin.Context) {
	userId := c.GetString("user_id")
	var goals []db.Goal
	db.DB.Where("user_id = ?", userId).Find(&goals)
	c.JSON(http.StatusOK, goals)
}

func CreateGoal(c *gin.Context) {
	userId := c.GetString("user_id")
	var body struct {
		Title      string     `json:"title"`
		TargetDate *time.Time `json:"target_date"`
		Status     string     `json:"status"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}
	goal := db.Goal{
		ID:         uuid.NewString(),
		UserID:     userId,
		Title:      body.Title,
		TargetDate: body.TargetDate,
		Status:     body.Status,
	}
	db.DB.Create(&goal)
	c.JSON(http.StatusOK, goal)
}
