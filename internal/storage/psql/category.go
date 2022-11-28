package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/catalogService/internal/models"
)

type Category struct {
	pool  *pgxpool.Pool
}

func NewCategoryStorage(db *pgxpool.Pool) *Category {
	return &Category{db}
}

func (c *Category) Create(ctx context.Context, title, parentID string) (string, error) {

	var id string

	sql := `INSERT INTO category (title, parent)
	VALUES ($1,$2)
	RETURNING id `

	rows, err := c.pool.Query(ctx,sql,
		title,
		parentID,
	)
	if err != nil {
		return id, fmt.Errorf("psql.storage.category.Create.Query %w", err)
	}

	for rows.Next(){
		err := rows.Scan(&id)
		if err != nil {
			return id, fmt.Errorf("psql.storage.category.Create.rows.Scan %w", err)
		}
	}

	return id, nil

}

func (c *Category) GetAll(ctx context.Context) ([]models.Category, error) {

	sql := `SELECT * FROM category`

	var cs []models.Category

	rows, err := c.pool.Query(ctx,sql)
	if err != nil {
		return cs, fmt.Errorf("psql.storage.category.Create.Query %w", err)
	} 

	for rows.Next() {

		ca := models.Category{}

		err := rows.Scan(
			&ca.ID,
			&ca.Title,
			&ca.Parent,
		)
		if err != nil {
			return cs, fmt.Errorf("psql.storage.category.GetAll.Query %w", err)
		}

		cs = append(cs, ca)

	}

	return cs, nil

}

func (c *Category) FindByID(ctx context.Context, cid string) (models.Category, error) {

	sql := `SELECT * FROM category WHERE id = $1`

	var ca models.Category

	rows, err := c.pool.Query(ctx,sql,cid)
	if err != nil {
		return ca, fmt.Errorf("psql.storage.category.FindByID.Query %w", err)
	}

	for rows.Next() {
		err := rows.Scan(
			&ca.ID,
			&ca.Title,
			&ca.Parent,
		)
		if err != nil {
			return ca, fmt.Errorf("psql.storage.category.FindByID.rows.Scan %w", err)
		}
	}

	return ca, nil

}

func (c *Category) FindByTitle(ctx context.Context, title string) (models.Category, error)  {

	sql := `SELECT * FROM category WHERE title = $1`

	var ca models.Category

	rows, err := c.pool.Query(ctx,sql,title)
	if err != nil {
		return ca, fmt.Errorf("psql.storage.category.FindByTitle.Query %w", err)
	}

	for rows.Next() {
		err := rows.Scan(
			&ca.ID,
			&ca.Title,
			&ca.Parent,
		)
		if err != nil {
			return ca, fmt.Errorf("psql.storage.category.FindByTitle.rows.Scan %w", err)
		}
	}

	return ca, nil
}

func (c *Category) Delete(ctx context.Context, cid string)  error {

	sql := `DELETE FROM category WHERE id = $1`

	_, err := c.pool.Query(ctx, sql, cid)
	if err != nil {
		return fmt.Errorf("psql.storage.category.Delete.Query %w", err)
	}

	return nil

}
