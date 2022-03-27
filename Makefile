.PHONY: all
all:	clean	\
	build	\
	test


.PHONY: clean
clean:
	@if [ -d 'target' ]; then	\
		rm -rf 'target';	\
	fi


.PHONY: build
build:
	go mod download
	go build -o 'target/sokoban-ebiten-assets' ./...


.PHONY: test
test:
	go test ./...

