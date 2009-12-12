include $(GOROOT)/src/Make.$(GOARCH)

TARG=rpc

all: server client

TARG=$proj
GOFILES=$proj.go

server: server.8
	$(LD) -o server server.$(O)

server.8: server.go arith.$(O)
	$(GC) -I. -o server.$(O) server.go

client: client.8
	$(LD) -o client client.$(O)

client.8: client.go arith.$(O)
	$(GC) -I. -o client.$(O) client.go

arith.8: arith.go
	$(GC) -I. -o arith.$(O) arith.go

clean:
	rm -f *.$(O) *.out server client

