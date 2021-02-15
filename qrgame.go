package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "io/ioutil"
    //"strings"
    "bytes"
    "github.com/google/brotli/go/cbrotli"
    //"github.com/skip2/go-qrcode"
)

func show_help() {
    fmt.Println("todo help page")
}

func load(fn string) {
    cmd := exec.Command("./tools/zbarimg/zbarimg", "-q", "--raw", "-Sbinary", fn)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    err := cmd.Run()

    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        log.Fatal(err)
        return
    }

    uncompressed, err := cbrotli.Decode(out.Bytes())
    if err != nil {
        log.Fatal(err)
        return
    }

    fp, err := os.Create("./a.out")
    if err != nil {
        log.Fatal(err)
        return
    }

    err = ioutil.WriteFile("./a.out", uncompressed, 0644)
    if err != nil {
        log.Fatal(err)
        return
    }

    defer fp.Close()
}

func main() {
    args := os.Args[1:]

    if len(args) == 2 {
        verb := args[0]
        noun := args[1]
        fmt.Println(verb, noun)

        if verb == "load" {
            load(noun)
        }
    } else {
        show_help()
    }
}
