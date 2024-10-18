package models

import "time"

type TransferRecord struct {
	Sender     string
	Recipient  string
	Amount     float32
	CreateTime time.Time
}
