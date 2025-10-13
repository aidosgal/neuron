build:
	@gcc -I/usr/local/include -I/usr/include/cjson -L/usr/local/lib -L/usr/lib \
    -o main src/main.c -lclips -lcjson -lm

run: build
	@./main
