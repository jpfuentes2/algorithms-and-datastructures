doc:
	godoc -http=:6060 -index

lint:
	golint ./go/...

test:
	go test ./go/...
