package main

import "fmt"

var cache = map[int]int{}

func fibCache(n int) int {
v, ok := cache[n]
if !ok {
v = fibRec(n)
cache[n] = v
}
return v
}

func fibRec(n int) int {
fmt.Print(n, " ")
if n <= 1 {
return n
}

return fibCache(n-1) + fibCache(n-2)
}

func main() {
/*for i := 0; i < 10;i++{
  fmt.Print(fibRec(i), ", ")
  }*/
fibRec(9)
}