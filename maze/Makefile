all: $(GOPATH)/bin/maze
	go fmt
	go test
	go vet
	go build -o $(GOPATH)/bin/maze cli/main.go

clean:
	rm -f *~

demo:
	echo done
