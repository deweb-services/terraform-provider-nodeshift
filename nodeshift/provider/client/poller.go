package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const pollerSecondsTick = 10 * time.Second

//go:generate mockgen -source poller.go -destination=./poller_mocks.go -package=client

type requester interface {
	DoSignedRequest(ctx context.Context, method string, endpoint string, body io.ReadSeeker) ([]byte, error)
}

func poll[T any](ctx context.Context, r requester, url string, getStatus func(*T) string) (*T, error) {
	ticker := time.NewTicker(pollerSecondsTick)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("polling canceled or timed out: %w", ctx.Err())

		case <-ticker.C:
			tflog.Debug(ctx, "polling: "+url)

			b, err := r.DoSignedRequest(ctx, http.MethodGet, url, nil)
			if err != nil {
				return nil, fmt.Errorf("failed to do signed request: %w", err)
			}

			tflog.Debug(ctx, "response from "+url, map[string]interface{}{
				"body": string(b),
			})

			var resp T
			if err := json.Unmarshal(b, &resp); err != nil {
				return nil, fmt.Errorf("failed to unmarshal response: %w", err)
			}

			switch strings.ToLower(getStatus(&resp)) {
			case "error":
				return nil, fmt.Errorf("failed to poll entity: %w", errStatusFailed)
			case "running":
				return &resp, nil
			default:
				continue
			}
		}
	}
}
