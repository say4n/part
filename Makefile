.PHONY: all

all: clean mondrian

mondrian: mondrian/main.go
	mkdir -p build
	go build -o build/mondrian mondrian/main.go

clean:
	rm -rf build

