NAME	=	sums

TARGET	=	./bin

all: build-macos build-windows build-linux

build-macos:
	GOOS=darwin GOARCH=amd64 go build -o $(TARGET)/$(NAME)_darwin_amd64

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(TARGET)/$(NAME)_win_amd64.exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(TARGET)/$(NAME)_linux_amd64

test:
	go test -v ./...

clean:
	rm -rf $(TARGET)

$(TARGET)/:
	mkdir -p $(TARGET)
