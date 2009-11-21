package main
import (
    "rpc";
    "http";
    "log";
    "net";
    "./arith";
);

func main() {
    arith := new(arith.Arith);
    rpc.Register(arith);
    rpc.HandleHTTP();
    l, e := net.Listen("tcp", ":1234");
    if e != nil {
        log.Exit("listen error: ", e);
    }
    http.Serve(l, nil);
}
