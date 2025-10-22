package dto

type AssignRolesRequest struct {
	RoleIDs []uint `json:"role_ids" validate:"required,min=1,dive,gt=0"`
}
