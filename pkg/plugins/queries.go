package plugins

import (
	"github.com/myback/open-grafana/pkg/bus"
	"github.com/myback/open-grafana/pkg/models"
)

func GetPluginSettings(orgId int64) (map[string]*models.PluginSettingInfoDTO, error) {
	query := models.GetPluginSettingsQuery{OrgId: orgId}

	if err := bus.Dispatch(&query); err != nil {
		return nil, err
	}

	pluginMap := make(map[string]*models.PluginSettingInfoDTO)
	for _, plug := range query.Result {
		pluginMap[plug.PluginId] = plug
	}

	for _, pluginDef := range Plugins {
		// ignore entries that exists
		if _, ok := pluginMap[pluginDef.Id]; ok {
			continue
		}

		// default to enabled true
		opt := &models.PluginSettingInfoDTO{
			PluginId: pluginDef.Id,
			OrgId:    orgId,
			Enabled:  true,
		}

		// apps are disabled by default unless autoEnabled: true
		if app, exists := Apps[pluginDef.Id]; exists {
			opt.Enabled = app.AutoEnabled
			opt.Pinned = app.AutoEnabled
		}

		// if it's included in app check app settings
		if pluginDef.IncludedInAppId != "" {
			// app components are by default disabled
			opt.Enabled = false

			if appSettings, ok := pluginMap[pluginDef.IncludedInAppId]; ok {
				opt.Enabled = appSettings.Enabled
			}
		}

		pluginMap[pluginDef.Id] = opt
	}

	return pluginMap, nil
}

func GetEnabledPlugins(orgId int64) (*EnabledPlugins, error) {
	enabledPlugins := NewEnabledPlugins()
	pluginSettingMap, err := GetPluginSettings(orgId)
	if err != nil {
		return nil, err
	}

	isPluginEnabled := func(pluginId string) bool {
		_, ok := pluginSettingMap[pluginId]
		return ok
	}

	for pluginId, app := range Apps {
		if b, ok := pluginSettingMap[pluginId]; ok {
			app.Pinned = b.Pinned
			enabledPlugins.Apps = append(enabledPlugins.Apps, app)
		}
	}

	// add all plugins that are not part of an App.
	for dsId, ds := range DataSources {
		if isPluginEnabled(ds.Id) {
			enabledPlugins.DataSources[dsId] = ds
		}
	}

	for _, panel := range Panels {
		if isPluginEnabled(panel.Id) {
			enabledPlugins.Panels = append(enabledPlugins.Panels, panel)
		}
	}

	return &enabledPlugins, nil
}

// IsAppInstalled checks if an app plugin with provided plugin ID is installed.
func IsAppInstalled(pluginID string) bool {
	_, exists := Apps[pluginID]
	return exists
}
