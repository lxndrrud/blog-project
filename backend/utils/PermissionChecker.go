package utils

import "github.com/lxndrrud/blog-project/models"

type IPermissionChecker interface {
	CanCreatePosts(permissions []models.Permission) bool
}

func NewPermissionChecker() IPermissionChecker {
	return &permissionChecker{}
}

type permissionChecker struct {
}

func (c *permissionChecker) CanCreatePosts(permissions []models.Permission) bool {
	for _, p := range permissions {
		if p.Code == "СОЗД_ПОСТЫ" {
			return true
		}
	}
	return false
}
