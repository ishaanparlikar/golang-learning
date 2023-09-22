package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
  welcome := "Welcome to go, Please Fill your input: "
  fmt.Printf(welcome)
  reader := bufio.NewReader(os.Stdin)

  //comma ok syntax
  input,_ := reader.ReadString('\n')

  fmt.Println("User Entered", input)

  //conversion of variables and handling errors

  rating,err := strconv.ParseFloat(strings.TrimSpace(input),64)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Your Rating is ", rating)
  }
}
