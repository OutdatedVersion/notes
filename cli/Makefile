clean:
	rm -rf build

build: clean
	mkdir build
	go build -o build/notes

installDependencies:
	go get -v github.com/mattn/go-sqlite3