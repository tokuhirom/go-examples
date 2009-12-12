package main
import (
    "os";
    "strconv";
    "log";
);

func fib(i int) int {
    if i <= 2 {
        return 1;
    }
    return fib(i-1) + fib(i-2);
}

func main() {
    if len(os.Args) != 2 {
        log.Exit("Usage: " + os.Args[0] + " n");
    }
    i, err := strconv.Atoi(os.Args[1]);
    if err != nil {
        log.Exit(err);
    }
    res := fib(i);
    println(res);
}

