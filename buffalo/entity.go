package buffalo

import "time"

type Buffalo struct {
	ID         int
	Name       string
	Family     string
	Population int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
