all: 
	go fmt
	go test
	go vet
	go build -o ${GOBIN}/maze cli/main.go

clean:
	rm -f *~ *.test

demo:
	echo done
