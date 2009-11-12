package main
import (
    "net";
    "os";
    "fmt";
)

func Err(format string, v ...) {
    fmt.Fprintf(os.Stderr, format + "\n", v)
}

func Handler(conn net.Conn) {
    defer conn.Close();

    buffer := make([]byte, 24);
    for {
        l, e := conn.Read(buffer);
        switch {
        case e == nil:
            conn.Write(buffer[0:l])
        case e == os.EOF:
            return;
        case e != os.EAGAIN:
            Err("Err on receiving a header (%s)", e);
            return;
        }
    }
}

func main() {
    psock, e := net.Listen("tcp", "127.0.0.1:1978");
    if e != nil {
        Err("an error occured(%s)", e);
        return
    }

    for {
        conn, e := psock.Accept();
        if e != nil {
            Err("an error occured(%s)", e);
            return
        }
        go Handler(conn);
    }
}

