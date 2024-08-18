package middleware

import (
	"time"
)

func GetNowTime() *string {
	// Create the user
	currentTime := time.Now()

	// 時刻を"YYYY-MM-DD HH:MM:SS"の形式でフォーマット（datetime(0)形式）
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return &formattedTime
}
