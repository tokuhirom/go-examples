package arith

import (
    "os"
);

type Args struct {
    A int;
    B int;
};
type Reply struct {
    C int;
};
type Arith int;
func (p*Arith) Add(args *Args, reply *Reply) os.Error {
    reply.C = args.A + args.B;
    return nil;
}

