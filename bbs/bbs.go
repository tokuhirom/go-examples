package main
import (
    "fmt";
    "template";
    "http";
    "io";
    "strings";
    "flag";
    "log";
    "container/vector";
    "os";
);

var addr = flag.String("addr", "0.0.0.0:1718", "http service address")

type tmplParams struct {
    msgs string;
};

type stringIO struct {
    buf string;
};
func (p *stringIO) Write(f []byte)(n int, err os.Error) {
    p.buf += string(f);
    return len(f), nil;
}

func html_escape(src string) string {
    sio := new(stringIO);
    template.HTMLEscape(sio, strings.Bytes(src));
    return sio.buf;
}

func main() {
    flag.Parse();
    templ := func () *template.Template {
        templateStr, err := io.ReadFile("tmpl/top.tmpl");
        if err != nil {
            log.Exit(err);
        }
        return template.MustParse(string(templateStr), nil);
    }();
    storage := new(vector.StringVector);

    http.Handle("/", http.HandlerFunc(func(c *http.Conn, req *http.Request) {
        params := new(tmplParams);
        part := "<ul>\n";
        for i:=storage.Len()-1; i>=0; i-- {
            msg := storage.At(i);
            part += "<li>"+html_escape(msg)+"</li>\n";
        }
        part += "</ul>";
        params.msgs = part;
        err := templ.Execute(params, c);
        if err != nil {
            log.Exit("templ.Execute:", err);
        }
    }));
    http.Handle("/post", http.HandlerFunc(func(c *http.Conn, req *http.Request) {
        req.ParseForm();
        body := req.Form["body"][0];
        storage.Push(body);
        http.Redirect(c, "/", 302);
    }));
    http.Handle("/css/", http.HandlerFunc(func(c *http.Conn, req *http.Request) {
        http.ServeFile(c, req, "." + req.URL.Path);
    }));

    fmt.Printf("http://%s/\n", *addr);
    err := http.ListenAndServe(*addr, nil);
    if err != nil {
        log.Exit("ListenAndServe:", err);
    }
}

