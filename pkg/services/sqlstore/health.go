package sqlstore

import (
	"github.com/myback/open-grafana/pkg/bus"
	"github.com/myback/open-grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *models.GetDBHealthQuery) error {
	_, err := x.Exec("SELECT 1")
	return err
}
