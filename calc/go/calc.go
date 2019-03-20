package main

import (
    "fmt"
    "log"
    "os"
    "strings"
)

// A golang version port of calc on https://pegjs.org/online

func main()  {
    log.SetFlags(log.Lshortfile | log.LstdFlags)
    log.SetPrefix(fmt.Sprintf("pid = %v ", os.Getpid()))


    r,err := ParseReader("1.md", strings.NewReader(os.Args[1]))

    log.Printf("return=%v err=%v", r, err)

    log.Printf("main exit")
}
