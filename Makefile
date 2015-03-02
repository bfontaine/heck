all: heck.peg.go
	go build .

heck.peg.go: heck.peg
	peg $<
