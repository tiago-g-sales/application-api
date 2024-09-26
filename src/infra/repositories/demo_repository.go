//go:generate mockgen -source demo_repository.go -destination mock/demo_repository_mock.go -package mock
package repositories

import (
	"github.com/example/application-api/src/domain/models"
	"context"
	"fmt"
)

var (
	memoryList = []models.Demo{
		{ID: 1, Name: "Demo 1"},
		{ID: 2, Name: "Demo 2"},
		{ID: 3, Name: "Demo 3"},
	}
)

type DemoRepository interface {
	FindAll(ctx context.Context) ([]models.Demo, error)
	FindById(ctx context.Context, id int) (*models.Demo, error)
}

type DemoDBRepository struct{}

func NewDemoMemoryRepository() *DemoDBRepository {
	return &DemoDBRepository{}
}

func (r *DemoDBRepository) FindAll(_ context.Context) ([]models.Demo, error) {
	return memoryList, nil
}

func (r *DemoDBRepository) FindById(_ context.Context, id int) (*models.Demo, error) {
	for _, demo := range memoryList {
		if demo.ID == id {
			return &demo, nil
		}
	}
	return nil, fmt.Errorf("demo with id %d not found", id)
}
