TARGET = heck

all: $(TARGET)

$(TARGET): deps $(wildcard *.go)
	go build .

%.peg.go: %.peg
	peg $<

check: deps
	go test -v ./...

clean:
	$(RM) $(TARGET) *~

deps:
	go get -v github.com/pointlander/peg
	go get -d -t -v ./...
