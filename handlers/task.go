package handlers

import (
	"engage-and-earn/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	DB *gorm.DB
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CompleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := h.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	if task.Completed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task already completed"})
		return
	}

	var user models.User
	if err := h.DB.First(&user, task.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	task.Completed = true
	user.Points += task.Points
	h.DB.Save(&task)
	h.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"task": task, "user_points": user.Points})
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID := c.Query("user_id")
	var tasks []models.Task
	if err := h.DB.Where("user_id = ? AND completed = ?", userID, false).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
