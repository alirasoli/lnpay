package data

import "context"

type Database interface {
	GetVersion(ctx context.Context) (int, error)
}
