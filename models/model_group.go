package models

type UserGroup struct {
	UserGroupId uint   `gorm:"primaryKey;column:user_group_id" json:"user_group_id"`
	UserGroup   string `json:"user_group"`
}

// Override nama tabel
func (UserGroup) TableName() string {
	return "t_user_group"
}
