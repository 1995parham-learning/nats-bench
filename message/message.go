package message

import "time"

type Message struct {
	Time    time.Time
	Payload []byte
}
