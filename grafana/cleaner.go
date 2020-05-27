package grafana

import (
	"context"
	"time"

	"github.com/lahsivjar/grafonnet-playground/config"
)

// startCleaner starts a job to periodically delete any stale dashboards
func startCleaner(ctx context.Context, pq *PriorityQueue, gSvc Service, cfg *config.Config) {
	ifBeforeNow := func(i *Item) bool {
		return i.ProcessAt.Before(time.Now())
	}

	ticker := time.NewTicker(cfg.AutoCleanupInterval)

	go func() {
		for ; true; <-ticker.C {
			select {
			case <-ctx.Done():
				return
			default:
				item := pq.PopConditionally(ifBeforeNow)

				if item != nil {
					err := gSvc.DeleteDashboard(item.Key)

					if err != nil {
						item.RetryCount = item.RetryCount + 1
						item.ProcessAt.Add(getBackoff(cfg.AutoCleanupMinBackoff, cfg.AutoCleanupMaxBackoff, item.RetryCount))
						pq.Push(item)
					}
				}
			}
		}
	}()
}

func getBackoff(minBackoff, maxBackoff time.Duration, retryCount int) time.Duration {
	currentBackoff := time.Duration(retryCount) * minBackoff

	if currentBackoff < maxBackoff {
		return currentBackoff
	}
	return maxBackoff
}
