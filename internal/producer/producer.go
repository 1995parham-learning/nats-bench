package producer

import (
	"time"

	"github.com/nats-ir/nats-bench/internal/common"
	"github.com/nats-ir/nats-bench/message"
	"github.com/sirupsen/logrus"
)

// MsgsPerClient divides the number of messages by
// the number of clients and tries to distribute them as evenly as possible.
func MsgsPerClient(nMsgs, nClients int) []int {
	var counts []int

	if nClients == 0 || nMsgs == 0 {
		return counts
	}

	counts = make([]int, nClients)

	mc := nMsgs / nClients

	for i := 0; i < nClients; i++ {
		counts[i] = mc
	}

	extra := nMsgs % nClients
	for i := 0; i < extra; i++ {
		counts[i]++
	}

	return counts
}

func Run(url, subject string, msgCount, msgSize int) {
	c := common.Connect(url)
	defer c.Close()

	var msg message.Message

	if msgSize > 0 {
		msg.Payload = make([]byte, msgSize)
	}

	msg.Time = time.Now()

	for i := 0; i < msgCount; i++ {
		if err := c.Publish(subject, msg); err != nil {
			logrus.Error(err)
		}
	}

	c.Flush()
}
