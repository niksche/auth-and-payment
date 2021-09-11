package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	authdelivery "github.com/niksche/flex/app/auth/delivery/http"
	authpostgres "github.com/niksche/flex/app/auth/repository/postgres"
	authusecase "github.com/niksche/flex/app/auth/usecase"
	"github.com/niksche/flex/app/utils/config"

	paymentdelivery "github.com/niksche/flex/app/payment/delivery/http"
	paymentpostgres "github.com/niksche/flex/app/payment/repository/postgres"
	paymentusecase "github.com/niksche/flex/app/payment/usecase"
)

func StartNew() {
	db := initDB()
	userRepo := authpostgres.NewUserRepository(db)
	userUC := authusecase.NewAuthUseCase(userRepo)
	userHandlers := authdelivery.NewHandler(userUC)

	paymentRepo := paymentpostgres.NewPaymentRepository(db)
	paymentUC := paymentusecase.NewPaymentUseCase(paymentRepo)
	paymentHandlers := paymentdelivery.NewHandler(paymentUC)

	router := mux.NewRouter()
	router.HandleFunc("/signup", userHandlers.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/login", userHandlers.LogIn).Methods(http.MethodPost)
	router.HandleFunc("/logout", userHandlers.LogOut).Methods(http.MethodGet)
	router.HandleFunc("/payment", paymentHandlers.Payment).Methods(http.MethodGet)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("cannot start service:", err)
	}
	fmt.Printf("started server at :8080")
}
func initDB() *pgxpool.Pool {
	dbConnPool, err := ConnectToDB(config.DbConfig)
	if err != nil {
		return nil
	}
	return dbConnPool
}

func ConnectToDB(db config.DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s pool_max_conns=%d",
		db.DbHost, db.DbName, db.DbUser, db.DbPassword, db.DbMaxConns,
	)

	return pgxpool.Connect(context.Background(), connStr)
}
