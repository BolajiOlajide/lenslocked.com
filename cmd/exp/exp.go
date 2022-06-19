package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// PostgresConfig configuration information for database connection
type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

// ToString returns the postgres url string using the credentials passed
func (cfg PostgresConfig) ToString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.SSLMode,
	)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.ToString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

	// 	_, err = db.Exec(`INSERT INTO users (age, first_name, last_name, email)
	// VALUES (28, 'Bolaji', 'Olajide', 'b@olaji.de')
	// `)
	// 	row := db.QueryRow(`INSERT INTO users (age, first_name, last_name, email)
	// VALUES ($1, $2, $3, $4) RETURNING id;`, 33, "Tolu", "Duyile", "ricky@gmail.com")

	// 	var id int32
	// 	err = row.Scan(&id)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Printf("User created. ID => %d\n", id)
	// userID := 3
	// row := db.QueryRow(`SELECT first_name, age FROM users WHERE id = $1`, userID)
	// var name string
	// var age int32
	// err = row.Scan(&name, &age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		fmt.Printf("User with ID %d not found!\n", userID)
	// 	} else {
	// 		panic(err)
	// 	}
	// }

	// fmt.Printf("The first name is %s, he's %d years old.\n", name, age)

	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)

	// 	_, err := db.Exec(`
	// INSERT INTO orders (user_id, amount, description)
	// VALUES ($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Created fake orders.")
	// }

	type Order struct {
		ID          int32
		UserID      int32
		Amount      int32
		Description string
	}

	var orders []Order
	userID := 3
	rows, err := db.Query(`
	SELECT id, amount, description
	FROM orders
	WHERE user_id=$1`, userID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserID = int32(userID)
		err = rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Orders: %v", orders)
}
