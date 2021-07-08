package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cenkayla/userbalance/internal/apiserver"
	"github.com/cenkayla/userbalance/internal/db"
	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect database.")
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	s := db.New(conn)
	server := apiserver.NewServer(*s)
	http.ListenAndServe(":8080", server)
}
