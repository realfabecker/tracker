package cache

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/realfabecker/wallet/internal/core/ports"
)

// FileCacheHandler
type FileCacheHandler struct{}

// FileCacheHandler
func NewFileCache() ports.CacheHandler {
	return &FileCacheHandler{}
}

// Get
func (c *FileCacheHandler) Get(key string) ([]byte, error) {
	d, err := os.ReadFile(filepath.Join(os.TempDir(), key+".json"))
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	return d, nil
}

// Set
func (c *FileCacheHandler) Set(key string, data []byte) error {
	return os.WriteFile(filepath.Join(os.TempDir(), key+".json"), data, 0644)
}
