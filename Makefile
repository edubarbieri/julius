tests:
	go test ./...

run_dev:
	env POSTGRES_URL="postgres://nfe:p@ssword@localhost/nfe" \
	go run main.go

run_prd:
	git pull
	docker-compose -f ./docker-compose.yml up --build