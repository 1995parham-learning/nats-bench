package producer

import (
	"sync"

	"github.com/nats-ir/nats-bench/config"
	"github.com/nats-ir/nats-bench/internal/producer"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	var donewg sync.WaitGroup

	donewg.Add(cfg.Instances)

	msgs := producer.MsgsPerClient(cfg.MessageCount, cfg.Instances)
	for _, nMsgs := range msgs {
		go func(nMsgs int) {
			producer.Run(cfg.URL, cfg.Subject, nMsgs, cfg.MessageSize)
			donewg.Done()
		}(nMsgs)
	}

	donewg.Wait()
}

// Register producer command.
func Register(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "producer",
		Short: "Produce messages to NATS",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.New(cmd.Flags())
			main(cfg)
		},
	}

	root.AddCommand(cmd)
}
