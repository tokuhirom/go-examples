package main
import (
    "net";
    "os";
    "fmt";
    "strings";
)

func Err(format string, v ...) {
    fmt.Fprintf(os.Stderr, format + "\n", v)
}

func main() {
    psock, e := net.Listen("tcp", "127.0.0.1:1978");
    if e != nil {
        Err("an error occured(%s)", e);
        return
    }

    // message queue
    queue := make(chan string);
    // channel map
    chan_for := make(map [net.Conn] chan string);

    // deliver thread
    go func() {
        for {
            msg := <-queue;
            for _, q := range chan_for {
                q <- msg;
            }
        }
    }();

    for {
        conn, e := psock.Accept();
        if e != nil {
            Err("an error occured(%s)", e);
            return
        }

        // reader thread
        go func() {
            defer conn.Close();

            buffer := make([]byte, 24);
            for {
                l, e := conn.Read(buffer);
                switch {
                case e == nil:
                    queue <- string(buffer[0:l]);
                case e == os.EOF:
                    return;
                case e != os.EAGAIN:
                    Err("Err on receiving a header (%s)", e);
                    return;
                }
            }
        }();

        w_ch := make(chan string);
        chan_for[conn] = w_ch;

        // writer thread
        go func() {
            for {
                msg := <-w_ch;
                _, err := conn.Write(strings.Bytes(msg));
                if err != nil {
                    break;
                }
            }
        }();
    }
}

