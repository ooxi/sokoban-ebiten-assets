.PHONY: all
all:	clean	\
	build	\
	test


.PHONY: clean
clean:
	@if [ -d 'target' ]; then	\
		rm -rf 'target';	\
	fi

	@mkdir 'target'


.PHONY: build
build:
	go mod download
	go build -o 'target' ./...


.PHONY: test
test:
	go test ./...

