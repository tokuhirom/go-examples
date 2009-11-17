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

var addr = flag.String("addr", "0.0.0.0:1718", "http service address") // Q=17, R=18
var fmap = template.FormatterMap{
    "html":     template.HTMLFormatter,
    "url+html": UrlHtmlFormatter,
}
var templ = template.MustParse(templateStr, fmap)
var storage = new(vector.StringVector);

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

    fmt.Printf("http://%s/\n", *addr);
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
    err := http.ListenAndServe(*addr, nil);
    if err != nil {
        log.Exit("ListenAndServe:", err);
    }
}

func UrlHtmlFormatter(w io.Writer, v interface{}, fmt string) {
    template.HTMLEscape(w, strings.Bytes(http.URLEscape(v.(string))));
}

const templateStr = `
<html>
<head>
    <title>bbs</title>
    <link rel="stylesheet" href="/css/blueprint/screen.css" type="text/css" media="screen, projection">
    <link rel="stylesheet" href="/css/blueprint/print.css" type="text/css" media="print">
    <!--[if lt IE 8]><link rel="stylesheet" href="/css/blueprint/ie.css" type="text/css" media="screen, projection"><![endif]-->
    <link rel="stylesheet" href="/css/bbs.css" type="text/css" media="screen, projection">
</head>
<body>
    <div class="container">
        <div class="span-24 last header">
            go bbs
        </div>
        <div class="span-24 last">
            <form method="post" action="/post">
                <input type="text" name="body" value="" />
                <input type="submit" value="post" />
            </form>
        </div>
        <div class="span-24 last">
            {msgs}
        </div>
        <hr />
        <div class="span-24 last footer">
            powered by <a href="http://golang.org/">go</a>.
        </div>
    </div>
</body>
</html>
`;

