package storage

import (
	"time"

	"github.com/cedy/private-virtual-phone/pkg/provider"
)

// Storage defines interface with storage layer.
type Storage interface {
	Save(provider.Message) error
	Get(id string) (provider.Message, error)
	GetMany(from time.Time, to time.Time, statuses []provider.Status) ([]provider.Message, error)
	GetWithStatus(provider.Status) error
}
