package response

import "time"

type Notification struct {
	ID             int       `json:"id"`
	IDUser         int       `json:"id_user"`
	IDNotification int       `json:"id_notification"`
	Description    string    `json:"description"`
	Date           time.Time `json:"date"`
}
