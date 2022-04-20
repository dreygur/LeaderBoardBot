run:
	go run bin/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o LeaderBoardBot bin/main.go