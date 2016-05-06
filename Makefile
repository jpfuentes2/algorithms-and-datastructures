doc:
	godoc -http=:6060 -index

lint:
	golint ./datastructures/... ./algorithms/...

test:
	go test -v ./datastructures/... ./algorithms/...
