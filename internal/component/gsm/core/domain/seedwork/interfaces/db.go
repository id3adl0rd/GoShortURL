package interfaces

import "context"

type Database interface {
	InsertURL(context.Context, interface{}) error
	GetByURL(context.Context, string) (string, error)
	GetAll(context.Context) (map[string]string, error)
}
