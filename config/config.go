package config

import (
	"log"
	"os"
)

// Configuration variables
var (
	JwtSecret      = []byte(os.Getenv("JWT_SECRET"))
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleSecret   = os.Getenv("GOOGLE_CLIENT_SECRET")
	RedirectURL    = os.Getenv("REDIRECT_URL")
	DbDSN          = os.Getenv("DB_DSN")
	RedisAddr      = os.Getenv("REDIS_ADDR")
)

func init() {
	if JwtSecret == nil || GoogleClientID == "" || GoogleSecret == "" || RedirectURL == "" || DbDSN == "" || RedisAddr == "" {
		log.Fatal("Missing environment variables")
	}
}
