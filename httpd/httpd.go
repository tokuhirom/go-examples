package main

import (
	"http";
	"io";
    "fmt";
)

// hello world, the web server
func HelloServer(c *http.Conn, req *http.Request) {
	io.WriteString(c, "hello, world!\n");
}

func main() {
    fmt.Printf("http://localhost:1978/hello\n");
	http.Handle("/hello", http.HandlerFunc(HelloServer));
	err := http.ListenAndServe(":1978", nil);
	if err != nil {
		panic("ListenAndServe: ", err.String())
	}
}

