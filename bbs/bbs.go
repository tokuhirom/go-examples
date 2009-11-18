package main
import (
    "fmt";
    "template";
    "http";
    "io";
    "flag";
    "log";
    "container/vector";
    "sqlite3";
);

var addr = flag.String("addr", "0.0.0.0:1718", "http service address")

func main() {
    sqlite3.Initialize();
    defer func () {
        log.Stdout("closing sqlite3");
        sqlite3.Shutdown();
    }();

    dbh := new(sqlite3.Handle);
    dbh.Open("bbs.db");
    defer dbh.Close();

    st,err := dbh.Prepare("CREATE TABLE IF NOT EXISTS entry (id INTEGER PRIMARY KEY, body VARCHAR(255));");
    if (err != "") {
        log.Exit(err);
    }
    st.Step();
    st.Finalize();

    flag.Parse();
    templ := func () *template.Template {
        templateStr, err := io.ReadFile("tmpl/top.tmpl");
        if err != nil {
            log.Exit(err);
        }
        return template.MustParse(string(templateStr), nil);
    }();

    http.Handle("/", http.HandlerFunc(func(c *http.Conn, req *http.Request) {
        params := new(struct { msgs []string });
        st,err = dbh.Prepare("SELECT * from entry ORDER BY id DESC limit 30;");
        if err != "" {
            log.Exit(err);
        }
        storage := new(vector.StringVector);
        func () {
            for {
                rv := st.Step();
                switch {
                case rv==sqlite3.SQLITE_DONE:
                    return;
                case rv==sqlite3.SQLITE_ROW:
                    body := st.ColumnText(1);
                    storage.Push(body);
                default:
                    println(rv);
                    log.Exit(dbh.ErrMsg());
                }
            };
        }();
        st.Finalize();
        params.msgs = storage.Data();
        err := templ.Execute(params, c);
        if err != nil {
            log.Exit("templ.Execute:", err);
        }
    }));
    http.Handle("/post", http.HandlerFunc(func(c *http.Conn, req *http.Request) {
        req.ParseForm();

        body := req.Form["body"][0];
        st,err = dbh.Prepare("INSERT INTO entry (body) VALUES (?)");
        if err != "" {
            log.Exit(err);
        }
        if st.BindText(1, body) != sqlite3.SQLITE_OK {
            log.Exit("cannot bind: ", dbh.ErrMsg());
        }
        st.Step();
        st.Finalize();

        http.Redirect(c, "/", 302);
    }));
    http.Handle("/css/", http.HandlerFunc(func(c *http.Conn, req *http.Request) {
        http.ServeFile(c, req, "." + req.URL.Path);
    }));

    // run httpd
    func() {
        fmt.Printf("http://%s/\n", *addr);
        err := http.ListenAndServe(*addr, nil);
        if err != nil {
            log.Exit("ListenAndServe:", err);
        }
    }();
}

