doc:
	godoc -http=:6060 -index

lint:
	golint ./...

test:
	go test -v ./...
