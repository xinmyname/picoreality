picoreality: cartgen
	./bin/cartgen
	

cartgen: ./cmd/cartgen.go
	go build -o ./bin/cartgen ./cmd/cartgen.go

.PHONY: clean

clean:
	rm -f ./bin/*
	rm -f picoreality.p8
