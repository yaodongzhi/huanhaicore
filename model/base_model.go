package model

type BaseModel struct {
	Id        uint64 `gorm:"primaryKey;type:bigint unsigned;comment:这是一个自增的索引" column:"id" json:"id"`
	Status    uint8  `gorm:"column:status;default:1;NOT NULL;comment:1状态正常，0失效" json:"status"`
	UpdatedAt string `gorm:"column:updated_at;NOT NULL;type:datetime(3);default:CURRENT_TIMESTAMP(3);onUpdate:CURRENT_TIMESTAMP(3);comment:更新时间" json:"updated_at"`
	CreatedAt string `gorm:"column:created_at;NOT NULL;type:datetime(3);default:CURRENT_TIMESTAMP(3);comment:创建时间" json:"created_at"`
}
