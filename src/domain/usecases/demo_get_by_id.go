//go:generate mockgen -source demo_get_by_id.go -destination mock/demo_get_by_id_mock.go -package mock
package usecases

import (
	"github.com/example/application-api/src/domain/models"
	"github.com/example/application-api/src/infra/repositories"
	"context"
)

type IDemoGetByID interface {
	Execute(context.Context, int) (*models.Demo, error)
}

func NewDemoGetByID() IDemoGetByID {
	return &DemoGetByID{Repo: repositories.NewDemoMemoryRepository()}
}

type DemoGetByID struct {
	Repo repositories.DemoRepository
}

func (uc *DemoGetByID) Execute(ctx context.Context, id int) (*models.Demo, error) {
	return uc.Repo.FindById(ctx, id)
}
