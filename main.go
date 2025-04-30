package main

import (
    "fmt"
    "log"
)

func main() {
    
    conn, err := pgx.Connect(context.Background(), "your-sqlite-database-connection-string")
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }
    defer conn.Close(context.Background())

    q := db.New(conn)
    workouts, err := q.GetWorkout(context.Background(), 1) 
    if err != nil {
        log.Fatalf("Error fetching workouts: %v", err)
    }

    fmt.Println(workouts)
}
