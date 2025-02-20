package main

import (
	"engage-and-earn/db"
	"engage-and-earn/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database := db.InitDB()

	// Handlers
	userHandler := handlers.UserHandler{DB: database}
	taskHandler := handlers.TaskHandler{DB: database}
	checkInHandler := handlers.CheckInHandler{DB: database}
	rewardHandler := handlers.RewardHandler{DB: database}

	// User Routes
	r.POST("/users", userHandler.Register)
	r.GET("/users/:id", userHandler.GetUser)

	// Task Routes
	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.GetTasks)
	r.POST("/tasks/:id/complete", taskHandler.CompleteTask)

	// Check-In Routes
	r.POST("/checkin/:user_id", checkInHandler.CheckIn)

	// Reward Routes
	r.GET("/rewards", rewardHandler.GetRewards)
	r.POST("/rewards/:user_id/:reward_id", rewardHandler.RedeemReward)

	// Seed some rewards
	database.Create(&handlers.Reward{Name: "Badge", PointsCost: 25, Description: "Virtual Badge"})
	database.Create(&handlers.Reward{Name: "Discount", PointsCost: 200, Description: "10% off"})

	r.Run(":8080")
}
