package conf

import "os"
import "github.com/joho/godotenv"

func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()
	if os.Getenv("ACTIVE_ENV") == "DEV" {
		_ = godotenv.Load(".env.dev")
	} else if os.Getenv("ACTIVE_ENV") == "PROD" {
		_ = godotenv.Load(".env.prod")
	}
}
