package librarypanels

import (
	"github.com/myback/open-grafana/pkg/models"
	"github.com/myback/open-grafana/pkg/services/dashboards"
	"github.com/myback/open-grafana/pkg/services/guardian"
)

func isGeneralFolder(folderID int64) bool {
	return folderID == 0
}

func requirePermissionsOnFolder(user *models.SignedInUser, folderID int64) error {
	if isGeneralFolder(folderID) && user.HasRole(models.ROLE_EDITOR) {
		return nil
	}

	if isGeneralFolder(folderID) && user.HasRole(models.ROLE_VIEWER) {
		return models.ErrFolderAccessDenied
	}

	s := dashboards.NewFolderService(user.OrgId, user)
	folder, err := s.GetFolderByID(folderID)
	if err != nil {
		return err
	}

	g := guardian.New(folder.Id, user.OrgId, user)

	canEdit, err := g.CanEdit()
	if err != nil {
		return err
	}
	if !canEdit {
		return models.ErrFolderAccessDenied
	}

	return nil
}
