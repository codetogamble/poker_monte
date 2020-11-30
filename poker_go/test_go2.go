package main

import(
  "fmt"
  "path"
  "path/filepath"
)

func main(){
  fmt.Println(path.Dir("/home/shubham/poker_auto"))
  fmt.Println(filepath.Dir("/home/shubham/poker_auto"))
}
