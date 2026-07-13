// centralizing database logic in one place
// creating a connection pooling thats going to maintain 10-20 connections in a pool
// request is going to borrow a conenction from the pool
// after queries are done, connection is going to return back to the pool

package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	// all of the requests have access to Pool

	//we are going to use a context for database operations (we can set timeouts regarding connection)
	var ctx context.Context = context.Background()

	//parsing connection string into config object
	var config *pgxpool.Config
	var err error

	config, err = pgxpool.ParseConfig(databaseURL) // contains all the info related to database (password, port, ssl enable or disabled etc.)

	if err != nil {
		log.Printf("Unable to parse database URL: %v", err)
		return nil, err
	}

	//using connection pool
	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		log.Printf("Unable to create connection pool: %v", err)
		return nil, err
	}

	const maxRetries = 10

	for i := 1; i <= maxRetries; i++ {
		err = pool.Ping(ctx)

		if err == nil {
			break
		}

		log.Printf("Database not ready (attempt %d/%d): %v", i, maxRetries, err)

		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Printf("Could not connect after %d attempts", maxRetries)
		pool.Close()
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL database")

	return pool, nil
}
