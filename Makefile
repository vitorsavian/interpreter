build:
	go build -o ./dist/interpreter

run: build
	./dist/interpreter
