package initalize

import (
	"context"
	"fmt"
	"github.com/Simnek/web-go/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
	"time"
)

func FastInitialize() {

	// Create a connection pool
	connConfig, err := pgxpool.ParseConfig(util.ConnectionString())
	if err != nil {
		fmt.Println("Error parsing connection string:", err)
		return
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), connConfig)
	if err != nil {
		fmt.Println("Error creating connection pool:", err)
		return
	}

	// Defer closing the connection pool
	defer pool.Close()
	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start time
	start := time.Now()

	// Data for table1
	dataTable1 := make([][]interface{}, 1000000)
	for i := 0; i < 1000000; i++ {
		dataTable1[i] = []interface{}{"John Doe", fmt.Sprintf("123 Main St %d", i)}
	}

	// Data for table2
	dataTable2 := make([][]interface{}, 1000000)
	for i := 0; i < 1000000; i++ {
		dataTable2[i] = []interface{}{"John Doe", fmt.Sprintf("email@email.com %d", i)}
	}

	// Execute CopyFrom concurrently in separate goroutines for each table
	wg.Add(2)
	go func() {
		defer wg.Done()

		// Get a connection from the pool
		conn, err := pool.Acquire(context.Background())
		if err != nil {
			fmt.Println("Unable to acquire connection from the pool for table1:", err)
			return
		}
		defer conn.Release()

		// Prepare the values for the query
		var values []interface{}
		for i := 0; i < 1000000; i += 10000 {
			// Prepare the SQL query
			query := `
        INSERT INTO customer (name, address)
        VALUES `

			batch := dataTable1[i : i+10000]
			for _, row := range batch {
				query += fmt.Sprintf("($%d, $%d),", len(values)+1, len(values)+2)
				values = append(values, row[0], row[1])
			}

			// Remove the last comma from the query
			query = query[:len(query)-1]

			// Execute the query
			_, err = conn.Exec(context.Background(), query, values...)
			if err != nil {
				fmt.Println("Error inserting rows into table1:", err)
				return
			}

			// Reset the values for the next batch
			values = nil
		}

		fmt.Println("Rows inserted into table1 successfully!")
	}()

	go func() {
		defer wg.Done()

		// Get a connection from the pool
		conn, err := pool.Acquire(context.Background())
		if err != nil {
			fmt.Println("Unable to acquire connection from the pool for table2:", err)
			return
		}
		defer conn.Release()

		// Prepare the values for the query
		var values []interface{}
		for i := 0; i < 1000000; i += 10000 {
			// Prepare the SQL query
			query := `
        INSERT INTO users (name, email)
        VALUES `

			batch := dataTable2[i : i+10000]
			for _, row := range batch {
				query += fmt.Sprintf("($%d, $%d),", len(values)+1, len(values)+2)
				values = append(values, row[0], row[1])
			}

			// Remove the last comma from the query
			query = query[:len(query)-1]

			// Execute the query
			_, err = conn.Exec(context.Background(), query, values...)
			if err != nil {
				fmt.Println("Error inserting rows into table2:", err)
				return
			}

			// Reset the values for the next batch
			values = nil
		}

		fmt.Println("Rows inserted into table2 successfully!")
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// End time
	end := time.Now()

	// Calculate the execution time
	executionTime := end.Sub(start)

	fmt.Println("Execution time:", executionTime)
}
