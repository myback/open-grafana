package api

import (
	"time"

	"github.com/myback/open-grafana/pkg/bus"
	"github.com/myback/open-grafana/pkg/models"
)

func (hs *HTTPServer) databaseHealthy() bool {
	const cacheKey = "db-healthy"

	if cached, found := hs.CacheService.Get(cacheKey); found {
		return cached.(bool)
	}

	healthy := bus.Dispatch(&models.GetDBHealthQuery{}) == nil

	hs.CacheService.Set(cacheKey, healthy, time.Second*5)
	return healthy
}
