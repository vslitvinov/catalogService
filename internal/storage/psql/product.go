package psql

import "github.com/jackc/pgx/v4/pgxpool"


type Product struct {
	pool  *pgxpool.Pool
}

// Create
// Update
// Delete
// FindByID
// FindByName
// FindByShop
// FindByCategory
// FindByTags ????
// FindByShopAndCategory ???
