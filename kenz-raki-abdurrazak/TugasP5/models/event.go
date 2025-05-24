package models

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SpeakerID   int    `json:"speaker_id"`
}
