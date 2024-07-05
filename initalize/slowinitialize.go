package initalize

import (
	"context"
	"fmt"
	"github.com/Simnek/web-go/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func SlowInitialize() {
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

	// Data for table1
	dataTable1 := make([][]interface{}, 10000)
	for i := 0; i < 10000; i++ {
		dataTable1[i] = []interface{}{"John Doe", fmt.Sprintf("123 Main St %d", i)}
	}

	// Data for table2
	dataTable2 := make([][]interface{}, 10000)
	for i := 0; i < 10000; i++ {
		dataTable2[i] = []interface{}{"John Doe", fmt.Sprintf("email@email.com %d", i)}
	}

	// Start time
	start := time.Now()

	// Insert data into table1
	query := `
    INSERT INTO customer (name, address)
    VALUES
`
	var values []interface{}
	for index, row := range dataTable1 {
		query += fmt.Sprintf("($%d, $%d),", 2*index+1, 2*index+2)
		values = append(values, row[0], row[1])
	}
	query = query[:len(query)-1]
	_, err = pool.Exec(context.Background(), query, values...)
	if err != nil {
		fmt.Println("Error inserting rows into table1:", err)
		return
	}
	fmt.Println("Rows inserted into table1 successfully!")

	// Insert data into table2
	query = `
    INSERT INTO users (name, email)
    VALUES
`
	values = []interface{}{}
	for index, row := range dataTable2 {
		query += fmt.Sprintf("($%d, $%d),", 2*index+1, 2*index+2)
		values = append(values, row[0], row[1])
	}
	query = query[:len(query)-1]
	_, err = pool.Exec(context.Background(), query, values...)
	if err != nil {
		fmt.Println("Error inserting rows into table2:", err)
		return
	}
	fmt.Println("Rows inserted into table2 successfully!")

	// End time
	end := time.Now()

	// Calculate the execution time
	executionTime := end.Sub(start)

	fmt.Println("Execution time:", executionTime)

	// Close the connection pool
	pool.Close()

	fmt.Println("Data insertion completed.")
}
