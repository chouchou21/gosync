package main

import "fmt"

type FuncType func(int, int) int

func Calc(a, b int, fTest FuncType) (result int) {
  fmt.Println("Calc")
  result = fTest(a, b) //这个函数还没有实现，但程序编译可以通过，--多态
  return
}

func main() {
  fmt.Println("hello, world")
}