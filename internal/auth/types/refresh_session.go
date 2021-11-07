package authTypes

import "time"

type RefreshSession struct {
	IP	string
	UserAgent	string
	CreatedAt time.Time
}
