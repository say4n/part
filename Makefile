.PHONY: all

all: clean mondrian wave

mondrian: mondrian/main.go
	mkdir -p build
	go build -o build/mondrian mondrian/main.go

wave: wave/main.go
	mkdir -p build
	go build -o build/wave wave/main.go

clean:
	rm -rf build

