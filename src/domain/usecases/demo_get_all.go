//go:generate mockgen -source demo_get_all.go -destination mock/demo_get_all_mock.go -package mock
package usecases

import (
	"github.com/example/application-api/src/domain/models"
	"github.com/example/application-api/src/infra/repositories"
	"context"
)

type IDemoGetAll interface {
	Execute(ctx context.Context) ([]models.Demo, error)
}

func NewDemoGetAll() IDemoGetAll {
	return &DemoGetAll{Repo: repositories.NewDemoMemoryRepository()}
}

type DemoGetAll struct {
	Repo repositories.DemoRepository
}

func (uc *DemoGetAll) Execute(ctx context.Context) ([]models.Demo, error) {
	return uc.Repo.FindAll(ctx)
}
