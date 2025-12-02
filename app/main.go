package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var (
	db       *sql.DB
	rdb      *redis.Client
	useCache bool
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	var err error

	dsn := os.Getenv("MYSQL_DSN")
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	useCache = os.Getenv("ENABLE_CACHE") == "true"

	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	http.HandleFunc("/user", getUserHandler)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id := r.URL.Query().Get("id")

	cacheKey := "user:" + id

	if useCache {
		if data, err := rdb.Get(ctx, cacheKey).Result(); err == nil {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(data))
			return
		}
	}

	var u User
	err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	jsonData, _ := json.Marshal(u)

	if useCache {
		rdb.Set(ctx, cacheKey, jsonData, 60*time.Second)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}
