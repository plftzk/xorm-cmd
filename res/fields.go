package res

type Fields struct {
	Id        int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Sn        string `json:"sn" xorm:"not null VARCHAR(32)"`
	UserSn    string `json:"user_sn" xorm:"not null CHAR(8)"`
	Name      string `json:"name" xorm:"not null VARCHAR(32)"`
	Icon      string `json:"icon" xorm:"not null default '' VARCHAR(128)"`
	Lat       string `json:"lat" xorm:"not null default '' VARCHAR(16)"`
	Lng       string `json:"lng" xorm:"not null default '' VARCHAR(16)"`
	Status    int    `json:"status" xorm:"not null default 0 TINYINT(1)"`
	DeletedAt int64  `json:"deleted_at" xorm:"not null default 0 INT(10)"`
	UpdatedAt int    `json:"updated_at" xorm:"not null INT(10)"`
	CreatedAt int    `json:"created_at" xorm:"not null INT(10)"`
}
