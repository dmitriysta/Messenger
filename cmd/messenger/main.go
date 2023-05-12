package messenger

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	os "os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "qwerty123456"
	dbname   = "messenger"
)

func dbConnection() {
	os.Setenv("host", "localhost")
	os.Setenv("port", "5432")
	os.Setenv("user", "admin")
	os.Setenv("password", "qwerty123456")
	os.Setenv("dbname", "messenger")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Fatal("Failed to open database connection: %v", zap.Error(err))
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		logger.Fatal("failed to ping database", zap.Error(err))
	}
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer logger.Sync()

	dbConnection()

	r := mux.NewRouter()
	userController := NewUserController(logger)
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	logger.Info("successfully connected to PostgreSQL database")

	port := ":80"
	logger.Info("server listening on port", zap.String("port", port))

	err = http.ListenAndServe(port, r)
	if err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}

}
