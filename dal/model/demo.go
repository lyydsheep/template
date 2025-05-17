package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type DemoOrder struct {
	Id           int64                 `gorm:"column:id;type:bigint;primaryKey;autoIncrement"`
	UserId       string                `gorm:"column:user_id;type:varchar(64)"`
	OrderId      string                `gorm:"column:order_id;type:varchar(64)"`
	BillMoney    int64                 `gorm:"column:bill_money" json:"bill_money"`         //订单金额（分）
	OrderGoodsId int64                 `gorm:"column:order_goods_id" json:"order_goods_id"` //订单对应的商品ID
	State        int8                  `gorm:"column:state;default:1" json:"state"`         //1-待支付，2-支付成功，3-支付失败
	PaidAt       time.Time             `gorm:"column:paid_at;default:\"1970-01-01 00:00:00\"" json:"paid_at"`
	IsDel        soft_delete.DeletedAt `gorm:"softDelete:flag"`
	CreatedAt    time.Time             `gorm:"column:created_at;autoCreateTime" json:"created_at"` //创建时间
	UpdatedAt    time.Time             `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` //更新时间
}
