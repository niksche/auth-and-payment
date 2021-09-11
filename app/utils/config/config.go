package config

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	JwtKey = []byte("secretkeydonttellanyonepls")
)

type Claims struct {
	Username           string `json:"username"`
	jwt.StandardClaims ``
}

type DBConfig struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
	DbMaxConns int
}

var (
	DbConfig = DBConfig{
		DbHost:     "localhost",
		DbName:     "test1",
		DbUser:     "test3",
		DbPassword: "test3",
		DbMaxConns: 10,
	}
)

func InitDB() *pgxpool.Pool {
	dbConnPool, err := ConnectToDB(DbConfig)
	if err != nil {
		return nil
	}
	return dbConnPool
}

func ConnectToDB(db DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s pool_max_conns=%d",
		db.DbHost, db.DbName, db.DbUser, db.DbPassword, db.DbMaxConns,
	)

	return pgxpool.Connect(context.Background(), connStr)
}
