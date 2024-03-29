package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Ashmit-05/go-interpreter/repl"
)

func main()  {
  user, err := user.Current()

  if err != nil {
    panic(err)
  }

  fmt.Printf("Congratulations %s, you have now become a space monkey!\n",user.Username)
  repl.Start(os.Stdin, os.Stdout)
  
}
