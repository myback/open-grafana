//go:build memcached
// +build memcached

package remotecache

import (
	"testing"

	"github.com/myback/open-grafana/pkg/setting"
)

func TestMemcachedCacheStorage(t *testing.T) {
	opts := &setting.RemoteCacheOptions{Name: memcachedCacheType, ConnStr: "localhost:11211"}
	client := createTestClient(t, opts, nil)
	runTestsForClient(t, client)
}
