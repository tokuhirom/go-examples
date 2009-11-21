package main
import (
    "rpc";
    "http";
    "log";
    "net";
);
import . "arith";

func main() {
    arith := new(Arith);
    rpc.Register(arith);
    rpc.HandleHTTP();
    l, e := net.Listen("tcp", ":1234");
    if e != nil {
        log.Exit("listen error: ", e);
    }
    http.Serve(l, nil);
}
/*
8g -o arith.8 arith.go
gopack grc arith.a arith.8
8g -I . rpc.go
*/
