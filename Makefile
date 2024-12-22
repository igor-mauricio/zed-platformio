hi:
	echo "Hello, World!"

clean:
	rm -f build/zed-platformio

build: clean
	go build -o build/zed-platformio .
