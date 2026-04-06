package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CabangID  int    `json:"cabang_id" gorm:"default:0"`
	GroupName string `json:"group_name"`
}

// Override nama tabel
func (User) TableName() string {
	return "t_user"
}
