dc:
	go build -o bin/main main.go

run:
	go build -o app cmd/url-shortener-service/main.go && PORT=3000 ./app