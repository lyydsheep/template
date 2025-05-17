package reply

type DemoOrder struct {
	OrderId   string `json:"orderId"`
	UserId    string `json:"userId"`
	BillMoney int64  `json:"billMoney"`
	PaidAt    string `json:"paidAt"`
	CreatedAt string `json:"createAt"`
	UpdatedAt string `json:"updateAt"`
}
