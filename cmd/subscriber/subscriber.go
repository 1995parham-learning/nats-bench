package subscriber

import (
	"sync"
	"time"

	"github.com/1995parham/nats-bench/config"
	"github.com/1995parham/nats-bench/internal/subscriber"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const TimeoutFlag = "timeout"

func main(cfg config.Config, timeout time.Duration) {
	var donewg sync.WaitGroup

	donewg.Add(cfg.Instances)

	for i := 0; i < cfg.Instances; i++ {
		go func() {
			s := subscriber.Run(cfg.URL, cfg.Subject, cfg.MessageCount, timeout)

			logrus.Infof("%+v", s)

			donewg.Done()
		}()
	}

	donewg.Wait()
}

// Register producer command.
func Register(root *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "subscriber",
		Short: "Subscribe messages from NATS",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.New(cmd.Flags())

			timeout, err := cmd.Flags().GetDuration(TimeoutFlag)
			if err != nil {
				logrus.Fatal(err)
			}

			main(cfg, timeout)
		},
	}

	cmd.Flags().Duration(TimeoutFlag, time.Hour, "subscription timeout")

	root.AddCommand(cmd)
}
