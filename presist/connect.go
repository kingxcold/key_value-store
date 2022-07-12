package presist

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var DB DbStore

func Connect() DbStore {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println("Couldn't load env file")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DB.Client = conn
	fmt.Println("DATABASE CONNECTED")
	return DB
}
