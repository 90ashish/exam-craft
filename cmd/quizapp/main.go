package main

import (
	"log"
	"os"

	"exam-craft/internal/handlers"
	"exam-craft/internal/repositories"
	"exam-craft/internal/routes"
	"exam-craft/internal/services"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	cache := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	userRepo := repositories.NewUserRepository(db)
	quizRepo := repositories.NewQuizRepository(db, cache)
	questionRepo := repositories.NewQuestionRepository(db)
	studentResponseRepo := repositories.NewStudentResponseRepository(db)

	userService := services.NewUserService(userRepo)
	quizService := services.NewQuizService(quizRepo)
	questionService := services.NewQuestionService(questionRepo)
	studentService := services.NewStudentService(studentResponseRepo, quizRepo)

	userHandler := handlers.NewUserHandler(userService)
	quizHandler := handlers.NewQuizHandler(quizService, questionService)
	studentHandler := handlers.NewStudentHandler(studentService)

	r := routes.SetupRouter(userHandler, quizHandler, studentHandler)
	r.Run(":8080")
}
