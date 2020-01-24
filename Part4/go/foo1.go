// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
for j:=0; j < 100000; j++ {
i +=1
}
}

func decrementing() {
for j:=0; j < 100000; j++ {
i -=1
}
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
go incrementing()
go decrementing()
    time.Sleep(100*time.Millisecond)
    Println("The magic number is:", i)
}
