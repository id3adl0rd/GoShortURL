package interfaces

import "context"

type Database interface {
	InsertURL(context.Context, string) error
	GetByURL(context.Context, string) (string, error)
	GetAll(context.Context) (map[string]string, error)
}
