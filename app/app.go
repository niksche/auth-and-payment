package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"

	authdelivery "github.com/niksche/flex/app/auth/delivery/http"
	authpostgres "github.com/niksche/flex/app/auth/repository/postgres"
	authusecase "github.com/niksche/flex/app/auth/usecase"

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

	// router := mux.NewRouter()
	// router.HandleFunc("/g", userHandlers.SignUp).Methods(http.MethodPost)

	// fmt.Printf("hellow world")
	// if err := http.ListenAndServe(":5000", router); err != nil {
	// 	fmt.Printf("cannot start service:", err)
	// }

	http.HandleFunc("/about", Payment)
	http.HandleFunc("/signup", userHandlers.SignUp)
	http.HandleFunc("/login", userHandlers.LogIn)
	http.HandleFunc("/logout", userHandlers.LogOut)
	http.HandleFunc("/payment", paymentHandlers.Payment)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Payment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func initDB() *pgxpool.Pool {
	dbConnPool, err := ConnectToDB(
		"localhost",
		"test1",
		"test3",
		"test3",
		10,
	)
	if err != nil {
		return nil
	}
	return dbConnPool
}

func ConnectToDB(dbHost, dbName, dbUser, dbPassword string, dbMaxConns int) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s pool_max_conns=%d",
		dbHost, dbName, dbUser, dbPassword, dbMaxConns,
	)

	return pgxpool.Connect(context.Background(), connStr)
}
