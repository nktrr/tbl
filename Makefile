loadTest:
	go build -o load_tester load_tester/cmd/main.go
	./load_tester/main
	rm ./load_tester/main
