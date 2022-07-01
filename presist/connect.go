package presist

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var DB DbStore

func Connect() DbStore {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:145963201@localhost:5432/key_val")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	DB.Client = conn
	fmt.Println("DATABASE CONNECTED")
	return DB
}
