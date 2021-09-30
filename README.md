# go_ecommerce

It's an API project to study repository pattern and table relationship in e-commerce using postgres as database

## Usage

Clone the repository

```bash
git clone https://github.com/victor-bern/go_ecommerce.git
```

Change connection string inside *src/database/database.go*

```go
func GetDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://USER:PASSWORD@localhost:PORT/DATABASENAME?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(10)

	return db
}

``` 

Start with ```go run main.go``` to run migrations

