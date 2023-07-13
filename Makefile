# Run the Docker containers
up:
	docker-compose up --build

# Run tests
test:
	go test

# Clean up the project
down:
	docker-compose down --volumes

# Access the MySQL terminal
mysql:
	docker-compose exec db mysql -h 127.0.0.1 -u'user' -p'password' transaction_routine

# Access the Go application terminal
go:
	docker-compose exec app bash

# Run go mod tidy to update module dependencies
install:
	go mod tidy