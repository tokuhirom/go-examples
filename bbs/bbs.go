package main
import (
    "fmt";
    "template";
    "http";
    "io";
    "flag";
    "log";
    "container/vector";
);

var addr = flag.String("addr", "0.0.0.0:1718", "http service address")

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
        params := new(struct { msgs []string });
        params.msgs = storage.Data();
        err := templ.Execute(params, c);
        if err != nil {
            log.Exit("templ.Execute:", err);
        }
    }));
    http.Handle("/post", http.HandlerFunc(func(c *http.Conn, req *http.Request) {
        req.ParseForm();
        body := req.Form["body"][0];
        storage.Insert(0, body);
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

