package handlers

import (
	"engage-and-earn/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RewardHandler struct {
	DB *gorm.DB
}

func (h *RewardHandler) GetRewards(c *gin.Context) {
	var rewards []models.Reward
	if err := h.DB.Find(&rewards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rewards"})
		return
	}
	c.JSON(http.StatusOK, rewards)
}

func (h *RewardHandler) RedeemReward(c *gin.Context) {
	userID := c.Param("user_id")
	rewardID := c.Param("reward_id")

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var reward models.Reward
	if err := h.DB.First(&reward, rewardID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reward not found"})
		return
	}

	if user.Points < reward.PointsCost {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient points"})
		return
	}

	user.Points -= reward.PointsCost
	h.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Reward redeemed", "remaining_points": user.Points})
}
