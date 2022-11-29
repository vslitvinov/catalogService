package storage

import (
	"context"
	"fmt"
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
	Delete(ctx context.Context, cid string) error
}

type Category struct {
	db      CategoryPSQLStorage
	cache   sync.Map
	isCache bool
}

func NewCategoryStorage(db *pgxpool.Pool) *Category {
	c := Category{
		db:    psql.NewCategoryStorage(db),
		cache: sync.Map{},
	}

	c.Start()

	return &c
}

func (c *Category) Start() {
	// load category db
}

func (c *Category) Create(ctx context.Context, title, parentID string) (string, error) {

	var cid string

	cid, err := c.db.Create(ctx, title, parentID)
	if err != nil {
		return cid, fmt.Errorf("storage.category.db.Create %w", err)
	}

	if c.isCache {

		mc := models.Category{
			ID:     cid,
			Title:  title,
			Parent: parentID,
		}

		c.cache.Store(cid, mc)

	}

	return cid, err

}

func (c *Category) GetAll(ctx context.Context) ([]models.Category, error) {

	var mcs []models.Category

	if c.isCache {
		c.cache.Range(func(k any, value any) bool {

			mcs = append(mcs, value.(models.Category))

			return true
		})

		return mcs, nil

	} else {
		var err error
		mcs, err = c.db.GetAll(ctx)
		if err != nil {
			return mcs, fmt.Errorf("storage.category.db.GetAll %w", err)
		}
	}

	return mcs, nil

}

func (c *Category) FindByTitle(ctx context.Context, title string) (models.Category, error) {

	var mc models.Category
	if c.isCache {

		c.cache.Range(func(k any, value any) bool {

			if value.(models.Category).Title == title {
				mc = value.(models.Category)
				return true
			} else {
				return false
			}

		})

		if &mc == nil {
			return mc, fmt.Errorf("storage.category.db.FindByTitle.notnull %w", nil)
		}

		return mc, nil

	} else {

		mc, err := c.db.FindByTitle(ctx, title)
		if err != nil {
			return mc, fmt.Errorf("storage.category.db.FindByTitle %w", err)
		}

	}
	return mc, nil

}

func (c *Category) FindByID(ctx context.Context, cid string) (models.Category, error) {

	var mc models.Category
	if c.isCache {

		_, ok := c.cache.LoadAndDelete(cid)
		if !ok {
			return mc, fmt.Errorf("storage.category.cache.LoadAndDelete %w", nil)
		}

	} else {

		mc, err := c.db.FindByID(ctx, cid)
		if err != nil {
			return mc, fmt.Errorf("storage.category.db.FindByID %w", err)
		}

	}
	return mc, nil

}

func (c *Category) Delete(ctx context.Context, cid string) error {

	err := c.db.Delete(ctx, cid)
	if err != nil {
		return fmt.Errorf("storage.category.db.Delete %w", err)
	}

	if c.isCache {
		_, ok := c.cache.LoadAndDelete(cid)
		if !ok {
			return fmt.Errorf("storage.category.cache.LoadAndDelete %w", nil)
		}
	}
	return nil
}
