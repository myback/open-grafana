package hooks

import (
	"github.com/myback/grafana/pkg/api/dtos"
	"github.com/myback/grafana/pkg/models"
	"github.com/myback/grafana/pkg/registry"
)

type IndexDataHook func(indexData *dtos.IndexViewData, req *models.ReqContext)

type LoginHook func(loginInfo *models.LoginInfo, req *models.ReqContext)

type HooksService struct {
	indexDataHooks []IndexDataHook
	loginHooks     []LoginHook
}

func init() {
	registry.RegisterService(&HooksService{})
}

func (srv *HooksService) Init() error {
	return nil
}

func (srv *HooksService) AddIndexDataHook(hook IndexDataHook) {
	srv.indexDataHooks = append(srv.indexDataHooks, hook)
}

func (srv *HooksService) RunIndexDataHooks(indexData *dtos.IndexViewData, req *models.ReqContext) {
	for _, hook := range srv.indexDataHooks {
		hook(indexData, req)
	}
}

func (srv *HooksService) AddLoginHook(hook LoginHook) {
	srv.loginHooks = append(srv.loginHooks, hook)
}

func (srv *HooksService) RunLoginHook(loginInfo *models.LoginInfo, req *models.ReqContext) {
	for _, hook := range srv.loginHooks {
		hook(loginInfo, req)
	}
}
