package main
import (
    "rpc";
    "log";
    "fmt";
    "./arith";
);

func main() {
    client, err := rpc.DialHTTP("tcp", "localhost:1234");
    if err != nil {
        log.Exit("dialing:", err);
    }
    args := &arith.Args{7,8};
    reply := new(arith.Reply);
    err = client.Call("Arith.Add", args, reply);
    if err != nil {
        log.Exit("arith error:", err);
    }
    fmt.Printf("Arith: %d+%d=%d\n", args.A, args.B, reply.C);
}
