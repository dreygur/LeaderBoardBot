run:
	go run .

build:
	GOOS=linux GOARCH=amd64 go build -o LeaderBoardBot .