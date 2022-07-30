tests:
	go test ./...

dev:
	env POSTGRES_URL="postgres://nfe:p@ssword@localhost/nfe" \
	env JWT_SECRET="weaksecret" \
	go run main.go

run_prd:
	git pull
	docker-compose -f ./docker-compose.yml up -d --build