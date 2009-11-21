package main
import (
    "rpc";
    "log";
    "fmt";
);
import . "arith";

func main() {
    client, err := rpc.DialHTTP("tcp", "localhost:1234");
    if err != nil {
        log.Exit("dialing:", err);
    }
    args := &Args{7,8};
    reply := new(Reply);
    err = client.Call("Arith.Add", args, reply);
    if err != nil {
        log.Exit("arith error:", err);
    }
    fmt.Printf("Arith: %d+%d=%d\n", args.A, args.B, reply.C);
}
