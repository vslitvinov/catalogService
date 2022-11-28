package storage

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vslitvinov/catalogService/internal/models"
	"github.com/vslitvinov/catalogService/internal/storage/psql"
)


type CategoryPSQLStorage interface {
	Create(ctx context.Context, title, parentID string) (string, error) 
	GetAll(ctx context.Context) ([]models.Category, error) 
	FindByID(ctx context.Context, cid string) (models.Category, error)
	FindByTitle(ctx context.Context, title string) (models.Category, error) 
	Delete(ctx context.Context, cid string)  error 
}

type Category struct {
	db CategoryPSQLStorage
	cache sync.Map
}


func NewCategoryStorage(db *pgxpool.Pool) *Category{
	c := Category{
		db:psql.NewCategoryStorage(db),
		cache: sync.Map{},
	}

	c.Start()

	return &c
}

func (c *Category) Start (){
	// load category db
}


func (c *Category) Create(ctx context.Context, title, parentID string) (string, error) {
	
}





