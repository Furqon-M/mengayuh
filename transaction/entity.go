package transaction

import "time"

type Transaction struct {
	ID        int
	Campaign  int
	userID    int
	Amount    int
	Status    string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
