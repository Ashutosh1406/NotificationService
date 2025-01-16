package handlers

import (
	"fmt"
)

func NotifyUsers(genreID string, genreName string) {
	message := fmt.Sprintf("A new song in the genre %s (%s) has been uploaded", genreName, genreID)
	broadcast <- message
}
