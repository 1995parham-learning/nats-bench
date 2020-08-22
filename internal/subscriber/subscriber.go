package subscriber

import (
	"log"
	"time"

	"github.com/nats-ir/nats-bench/internal/common"
	"github.com/nats-ir/nats-bench/message"
	"github.com/nats-ir/nats-bench/sample"
)

type metrics struct {
	tripTime time.Duration
	size     int
}

func Run(url, subject string, msgCount int, timeout time.Duration) sample.Sample {
	c := common.Connect(url)
	defer c.Close()

	ch := make(chan metrics)

	if _, err := c.Subscribe(subject, func(m *message.Message) {
		go func() {
			ch <- metrics{
				tripTime: time.Since(m.Time),
				size:     len(m.Payload),
			}
		}()
	}); err != nil {
		log.Fatal(err)
	}

	var att time.Duration

	var start time.Time

	counter := 0
	over := false
	asize := 0

	for {
		ticker := time.NewTicker(timeout)
		select {
		case m := <-ch:
			att += m.tripTime
			asize += m.size
			counter++

			if counter == 1 {
				start = time.Now()
			}
		case <-ticker.C:
			over = true
		}
		ticker.Stop()

		if counter >= msgCount {
			counter = msgCount

			break
		}

		if over {
			break
		}
	}

	return sample.Sample{
		TripTime:         att / time.Duration(counter),
		MessagePerSecond: float64(counter) / time.Since(start).Seconds(),
		BytesPerSecond:   float64(asize) / time.Since(start).Seconds(),
		Duration:         time.Since(start),
	}
}
