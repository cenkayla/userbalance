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
![get balance](https://user-images.githubusercontent.com/16439581/124961467-30f9fa00-e026-11eb-8440-abe7b691c7ff.PNG)

To add balance
```shell
$ curl --request PUT \
	--url "localhost:8080/users/balance/add/1?balance=50" \
	--header "Content-Type: application/x-www-form-urlencoded" \
```
![user balance add](https://user-images.githubusercontent.com/16439581/124961600-5ab32100-e026-11eb-8172-ddbe0fa62d0d.PNG)

To reduce balance
```shell
curl --request PUT \
	--url "localhost:8080/users/balance/reduce/1?balance=50" \
	--header "Content-Type: application/x-www-form-urlencoded" \
```
![balance reduce](https://user-images.githubusercontent.com/16439581/124961680-70c0e180-e026-11eb-95d6-9c045f87e551.PNG)

To transfer balance
```shell
curl --request PUT \
	--url "localhost:8080/users/balance/transfer/1?receiver_id=0&balance=20" \
	--header "Content-Type: application/x-www-form-urlencoded" \
```
![transfer balance](https://user-images.githubusercontent.com/16439581/124961733-80d8c100-e026-11eb-9243-8aa38b57aa91.PNG)



