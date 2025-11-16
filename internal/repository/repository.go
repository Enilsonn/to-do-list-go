package repository

import (
	"context"
	"database/sql"
	"to-do-list/internal/models"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

func (r *SQLiteRepository) Create(ctx context.Context, t *models.Tarefa) (int64, error) {
	query := `INSERT INTO tasks (title, done)
			  VALUES (?, ?)`

	result, err := r.db.ExecContext(ctx, query,
		t.Title,
		t.Done)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SQLiteRepository) List(ctx context.Context) ([]*models.Tarefa, error) {
	query := `SELECT id, title, done
			  FROM tasks`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*models.Tarefa
	for rows.Next() {
		var t models.Tarefa
		if err := rows.Scan(&t.ID, &t.Title, &t.Done); err != nil {
			return nil, err
		}
		list = append(list, &t)
	}
	return list, nil
}

func (r *SQLiteRepository) GetByID(ctx context.Context, id int64) (*models.Tarefa, error) {
	query := `SELECT id, title, done
			  FROM tasks
			  WHERE id=?`

	var t models.Tarefa

	if err := r.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.Title, &t.Done); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *SQLiteRepository) Update(ctx context.Context, id int64, t *models.Tarefa) error {
	query := `UPDATE tasks
			  SET title=?, done=?
			  WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, t.Title, t.Done, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *SQLiteRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks
			  WHERE id=?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
