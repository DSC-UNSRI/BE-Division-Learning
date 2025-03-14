package data

import (
	"os"
	"time"
)

// Fungsi untuk mencatat log ke dalam file
func CatatLog(data string) {
	file, err := os.OpenFile("data/log_iftar.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	logData := time.Now().Format("2006-01-02 15:04:05") + " - " + data + "\n"
	file.WriteString(logData)
}
