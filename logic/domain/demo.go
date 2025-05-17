package domain

import "time"

// DemoOrder 仅保留有业务含义的字段
type DemoOrder struct {
	Id           int64     `json:"id"`
	UserId       string    `json:"userId"`
	BillMoney    int64     `json:"billMoney"`
	OrderId      string    `json:"orderNo"`
	OrderGoodsId int64     `json:"orderGoodsId"`
	State        int8      `json:"state"`
	PaidAt       time.Time `json:"paidAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
