build:
	GOOS=windows GOARCH=amd64 go build -o bin/ksp_save_file_parser.exe main.go

run:
	go run main.go