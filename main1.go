package main 

import (
  "bufio"
  "os"
  "fmt"
  "strings"
)

func main() {
  
  reader := bufio.NewScanner(os.Stdin)
  transactionManager := transactionManager.NewManager()

  for {
    fmt.Println("> ")
    text, _ := reader.ReadString("\n")
    args := strings.Fields(text)

    if len(args) == 0 { continue }

    cmd, ok := commands.Get(args[0])
    if !ok { 
      fmt.Println("Command Not found")
      continue
    }
  }

}


