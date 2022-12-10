// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameShareBasic = "share_basic"

// ShareBasic mapped from table <share_basic>
type ShareBasic struct {
	ID                     int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Identity               string         `gorm:"column:identity" json:"identity"`
	UserIdentity           string         `gorm:"column:user_identity" json:"user_identity"`
	RepositoryIdentity     string         `gorm:"column:repository_identity" json:"repository_identity"`           // 公共池中的唯一标识
	UserRepositoryIdentity string         `gorm:"column:user_repository_identity" json:"user_repository_identity"` // 用户池子中的唯一标识
	ExpiredTime            int32          `gorm:"column:expired_time" json:"expired_time"`                         // 失效时间，单位秒, 【0-永不失效】
	ClickNum               int32          `gorm:"column:click_num" json:"click_num"`                               // 点击次数
	CreatedAt              time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ShareBasic's table name
func (*ShareBasic) TableName() string {
	return TableNameShareBasic
}
