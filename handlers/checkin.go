package handlers

import (
	"engage-and-earn/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CheckInHandler struct {
	DB *gorm.DB
}

func (h *CheckInHandler) CheckIn(c *gin.Context) {
	userID := c.Param("user_id")
	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	now := time.Now()
	lastCheckIn := user.LastCheckIn
	today := now.Truncate(24 * time.Hour)

	if lastCheckIn.Truncate(24 * time.Hour).Equal(today) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already checked in today"})
		return
	}

	points := 2 // Base check-in points
	if lastCheckIn.Truncate(24 * time.Hour).Add(24 * time.Hour).Equal(today) {
		user.Streak++
	} else {
		user.Streak = 1
	}

	// Streak bonuses
	if user.Streak == 5 {
		points += 5
	} else if user.Streak == 10 {
		points += 15
	}

	user.Points += points
	user.LastCheckIn = now
	h.DB.Save(&user)

	checkIn := models.CheckIn{
		UserID: user.ID,
		Date:   now,
		Points: points,
	}
	h.DB.Create(&checkIn)

	c.JSON(http.StatusOK, gin.H{
		"points":       points,
		"streak":       user.Streak,
		"total_points": user.Points,
	})
}
