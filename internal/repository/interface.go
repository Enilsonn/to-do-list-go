package repository

import (
	"context"
	"to-do-list/internal/models"
)

type TaskRepository interface {
	Create(ctx context.Context, t *models.Tarefa) (int64, error)
	List(ctx context.Context) ([]*models.Tarefa, error)
	GetByID(ctx context.Context, id int64) (*models.Tarefa, error)
	Update(ctx context.Context, id int64, t *models.Tarefa) error
	Delete(ctx context.Context, id int64) error
}
