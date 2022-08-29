build:
	go build -o ./dist/server

test:
	go test ./...

cover:
	go test -coverprofile=c.out
	go tool cover -html=c.out

vet:
	go vet ./...

clean:
	rm -rf ./dist