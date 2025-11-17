package routes

import (
	"github.com/gin-gonic/gin"
	"skillup-backend/controllers"
	"skillup-backend/middleware"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes (no auth required)
	r.POST("/api/auth/signup", controllers.Signup)
	r.POST("/api/auth/login", controllers.Login)

	// Protected routes (auth required)
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	// goals & topics
	api.GET("/goals", controllers.GetGoals)
	api.POST("/goals", controllers.CreateGoal)

	api.GET("/topics", controllers.GetTopics)
	api.POST("/topics", controllers.CreateTopic)

	// tasks -> handled as study activities or topics/tasks inside frontend; we keep activity endpoints
	api.POST("/activity", controllers.CreateActivity)
	api.GET("/activity", controllers.GetActivity)

	// documents & PDF ingestion
	api.POST("/documents/upload", controllers.UploadDocument)
	api.GET("/documents", controllers.GetDocuments)

	// chat (RAG)
	api.POST("/chat/query", controllers.ChatQuery)

	// quizzes
	api.POST("/quizzes", controllers.CreateQuiz)
	api.GET("/quizzes", controllers.GetQuizzes)
}
