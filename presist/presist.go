package presist

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
)

type DbStore struct {
	Client *pgx.Conn
}

func (dbs DbStore) Set(key string, value any) {
	_, err := dbs.Client.Exec(context.Background(), "INSERT INTO store (key, value) VALUES($1, $2)", key, value)
	if err != nil {
		log.Fatal(err)
	}
}

func (dbs DbStore) Get(key string) (any, error) {
	var value any
	err := dbs.Client.QueryRow(context.Background(), "SELECT value FROM store WHERE key=$1", key).Scan(&value)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return "", errors.New("Data not found")
		} else {
			log.Fatal(err)
		}
	}
	return value, nil
}

func (dbs DbStore) Delete(key string) {
	_, err := dbs.Client.Exec(context.Background(), "DELETE FROM store WHERE key=$1", key)
	if err != nil {
		log.Fatal(err)
	}
}
