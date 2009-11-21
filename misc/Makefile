include $(GOROOT)/src/Make.$(GOARCH)

TARG=rpc

all: server client

server: server.8 arith.a
	$(LD) -o server server.$(O)

server.8: server.go arith.a
	$(GC) -I. -o server.$(O) server.go

client: client.8 arith.a
	$(LD) -o client client.$(O)

client.8: client.go arith.a
	$(GC) -I. -o client.$(O) client.go

arith.a: arith.$(O)
	gopack grc arith.a arith.$(O)

arith.8: arith.go
	$(GC) -o arith.8 arith.go

