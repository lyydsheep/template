package appService

import (
	"context"
	"your-module-name/api/reply"
	"your-module-name/api/request"
)

type DemoAppService interface {
	GetAllIdentities(c context.Context) ([]int64, error)
	CreateDemoOrder(c context.Context, order *request.DemoOrderReq) (*reply.DemoOrder, error)
}
