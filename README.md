# User balance service

- Written in Go
- Stores data in PostgreSQL with pgx

# Quickstart
```shell
git clone https://github.com/cenkayla/userbalance.git
cd userbalance
go run cmd/main.go
```

# Usage

To get balance
```shell
$ curl 'localhost:8080/users/balance/1'
```

To add balance
```shell
$ curl --request PUT \
	--url "localhost:8080/users/balance/add/1?balance=50" \
	--header "Content-Type: application/x-www-form-urlencoded" \
```

To reduce balance
```shell
curl --request PUT \
	--url "localhost:8080/users/balance/reduce/1?balance=50" \
	--header "Content-Type: application/x-www-form-urlencoded" \
```

To transfer balance
```shell
curl --request PUT \
	--url "localhost:8080/users/balance/transfer/1?receiver_id=0&balance=20" \
	--header "Content-Type: application/x-www-form-urlencoded" \
```
