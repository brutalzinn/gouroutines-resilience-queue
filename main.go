package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	var wg sync.WaitGroup

	job_one := Job{
		Name:   "Job 1",
		Action: test,
	}

	job_two := Job{
		Name:   "Job 2",
		Action: test,
	}
	job_three := Job{
		Name:   "Job 3",
		Action: test,
	}
	job_four := Job{
		Name:   "Job 4",
		Action: error_test,
	}

	var jobs = []Job{
		job_one, job_two, job_three, job_four,
	}

	for _, item := range jobs {
		item.Execute(&wg)
		item.insert_log()
	}

}

func error_test() error {
	return errors.New("something happens")
}

func test() error {
	fmt.Println("test default %s", time.Now())
	return nil
}

func (job Job) insert_log() {
	database_url := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), database_url)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "INSERT INTO queue (name, priority, status) VALUES ($1, $2, $3)", job.Name, job.Priority, job.Status)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

}
