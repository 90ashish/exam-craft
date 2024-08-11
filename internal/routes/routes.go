package routes

import (
	"exam-craft/internal/handlers"
	"exam-craft/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application
func SetupRouter(userHandler *handlers.UserHandler, quizHandler *handlers.QuizHandler, studentHandler *handlers.StudentHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.GET("/auth/google", userHandler.GoogleLogin)
	r.GET("/auth/google/callback", userHandler.GoogleCallback)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/quizzes", middleware.RoleMiddleware("instructor"), quizHandler.CreateQuiz)
		authorized.GET("/quizzes/:id", quizHandler.GetQuiz)
		authorized.POST("/quizzes/:id/questions", middleware.RoleMiddleware("instructor"), quizHandler.AddQuestion)

		authorized.GET("/student/quizzes", middleware.RoleMiddleware("student"), studentHandler.ListQuizzes)
		authorized.POST("/student/quizzes/:id/response", middleware.RoleMiddleware("student"), studentHandler.SubmitResponse)
	}

	return r
}
