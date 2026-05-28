// checking if the changes will be pushed or not

package main

import (
	"log"
	"todo-api/internal/config"
	"todo-api/internal/database"
	"todo-api/internal/handlers"
	"todo-api/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	//get info of db from config
	var cfg *config.Config
	var err error
	cfg, err = config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	//create database pool
	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	defer pool.Close() // defer keyword schedules the close method once the main function exists (closes connection pool for us)

	var router *gin.Engine = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {

		//gin.H is a map with key as a string and value is any datatype
		c.JSON(200, gin.H{
			"message":  "Todo API is running",
			"status":   "Success",
			"database": "Connected",
		})
	})

	router.POST("/auth/register", handlers.CreateUserHandler(pool))
	router.POST("/auth/login", handlers.LoginHandler(pool, cfg))

	protected := router.Group("/todos")
	protected.Use(middleware.AuthMiddleware(cfg))

	protected.POST("", handlers.CreateTodoHandler(pool))
	protected.GET("", handlers.GetAllTodosHandler(pool))
	protected.GET("/:id", handlers.GetTodosByIdHandler(pool))
	protected.PUT("/:id", handlers.UpdateTodoHandler(pool))
	protected.DELETE("/:id", handlers.DeleteTodoHandler(pool))

	router.Run(":" + cfg.Port)
}
