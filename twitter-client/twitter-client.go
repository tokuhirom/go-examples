// ref. http://mattn.kaoriya.net/software/lang/go/20091112002814.htm

package main

import (
    "fmt";
    "json";
    "io";
    "http";
)

func main() {
    r, _, err := http.Get("http://twitter.com/statuses/public_timeline.json");
    if err == nil {
        b, _ := io.ReadAll(r.Body);
        j, _, _ := json.StringToJson(string(b));
        for i := 0; i < j.Len(); i++ {
        data := j.Elem(i);
        fmt.Printf("%s: %s\n",
            data.Get("user").Get("screen_name"),
            data.Get("text"));
        }
    }
}

