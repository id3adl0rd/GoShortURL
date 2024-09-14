package db

import "context"

type MockDB struct {
}

func (m *MockDB) InsertURL(ctx context.Context, i interface{}) error {
	return nil
}

func (m *MockDB) GetByURL(ctx context.Context, s string) (string, error) {
	return "", nil
}

func (m *MockDB) GetAll(ctx context.Context) (map[string]string, error) {
	return nil, nil
}
