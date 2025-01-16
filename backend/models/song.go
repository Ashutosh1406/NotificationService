package models

import "github.com/Ashutosh1406/NotificationService/backend/handlers"

type Song struct {
	ID    string
	Title string
	Genre string
}

func NewSongUploaded(song Song) {
	handlers.NotifyUsers(song.Genre, song.Title)
}
