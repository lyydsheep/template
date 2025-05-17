package dao

import (
	"context"
	"your-module-name/dal/model"
	"your-module-name/logic/domain"
)

type DemoDAO interface {
	FindAllDemo(c context.Context) ([]model.DemoOrder, error)
	CreateDemoOrder(c context.Context, order *domain.DemoOrder) (*model.DemoOrder, error)
}
