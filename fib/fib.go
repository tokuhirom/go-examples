package main
import fmt "fmt"

func fib(i int) int {
    if i <= 2 {
        return 1;
    }
    return fib(i-1) + fib(i-2);
}

func main() {
    fmt.Printf("%d\n", fib(10));
}
