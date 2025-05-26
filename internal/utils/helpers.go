package utils

import (
	"browser-mmo-backend/internal/constants"
	"time"
)

// GetCurrentDate returns the current date in the format YYYY-MM-DD.
func GetCurrentDate() string {
	currentTime := time.Now().UTC()
	return currentTime.Format(constants.TimeFormatJustDate)
}
