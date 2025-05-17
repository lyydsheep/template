package request

type DemoOrderReq struct {
	BillMoney    int64  `json:"billMoney"`
	OrderGoodsId int64  `json:"orderGoodsId"`
	UserId       string `json:"userId"`
}
