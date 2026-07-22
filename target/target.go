package target

import (
	"context"
	"time"
)

// Target interface defines the methods of target
type Target interface {
	// Connect connects to the target databse
	Connect(
		ctx context.Context,
		timeout time.Duration,
	) error
}
