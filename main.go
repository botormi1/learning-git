package main

func main() {
  fmt.Println("Hello, Go!")
  fmt.Println("Welcome to my program.")
  var a int = 100
  fmt.Println("{}", a)
  fmt.Println("max(1, 2) = {}", max(1, 2))
}

func max(a, b int) int {
  return a > b ? a : b
}
