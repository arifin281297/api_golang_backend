package models

type UserCabang struct {
	CabangID uint   `gorm:"primaryKey;column:cabang_id" json:"cabang_id"`
	Cabang   string `json:"cabang"`
	Alias    string `json:"alias"`
}

// Override nama tabel
func (UserCabang) TableName() string {
	return "t_user_cabang"
}
