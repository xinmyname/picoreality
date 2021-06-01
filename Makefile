cartgen: ./cmd/cartgen.go
	go build -o ./bin/cartgen ./cmd/cartgen.go

picoreality: cartgen
	./bin/cartgen

.PHONY: clean

clean:
	rm -f ./bin/*
	rm picoreality.p8
