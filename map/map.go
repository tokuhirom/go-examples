package main
import ( "fmt"; );

func main() {
    m := make(map[string] string);
    m["Foo"] = "Bar";
    m["Hoge"] = "Fuga";
    fmt.Printf("%s\n", m["Foo"]);
}
