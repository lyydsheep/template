package domainService

import (
	"context"
	"your-module-name/logic/domain"
)

// DemoDomainService 保持依赖注入风格
type DemoDomainService interface {
	GetDemos(context.Context) ([]domain.DemoOrder, error)
	CreateDemoOrder(context.Context, *domain.DemoOrder) (*domain.DemoOrder, error)
}
