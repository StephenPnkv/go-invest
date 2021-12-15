package main

import (
    "log"
    "os"
    "strings"

    "github.com/andybalholm/cascadia"
    "golang.org/x/net/html"
)

func main() {
    r := strings.NewReader(`<!DOCTYPE html>
<html>
    <head>
        <title>
            Title of the document
        </title>
    </head>
    <body>
        body content
        <p>more content</p>
    </body>
</html>`)
    doc, err := html.Parse(r)
    if err != nil {
        log.Fatal(err)
    }

    body := cascadia.MustCompile("body").MatchFirst(doc)
    html.Render(os.Stdout, body)
}
