package entity

type UserGroup struct {
	ID             string  `json:"id"`
	Name           string  `json:"name,omitempty"`
	RingID         string  `json:"ringId,omitempty"`
	UserCount      int64   `json:"userCount"`
	Icon           string  `json:"icon,omitempty"`
	AllUsersGroup  bool    `json:"allUsersGroup"`
	TeamForProject Project `json:"teamForProject,omitempty"`
}
