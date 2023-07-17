//go:build inotify

package file

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestWatcher(t *testing.T) {
	dir, ok := os.LookupEnv("INOTIFY_TEST_DIR")
	if !ok {
		t.Fatal("INOTIFY_TEST_DIR is not setting")
	}
	watcher, err := NewWatcher(dir)
	require.NoError(t, err)
	err = watcher.Start(func(filename string) {

	})
	require.NoError(t, err)
	time.Sleep(2 * time.Second)
	watcher.Stop()
}
