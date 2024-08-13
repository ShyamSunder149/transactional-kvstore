package main 

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "transactional-kvstore/commands"
  "transactional-kvstore/transaction"
)

func main() {
  
  reader := bufio.NewReader(os.Stdin)
  transactionManager := transaction.NewTransactionManager()

  for {
    fmt.Printf("> ")
    text, _ := reader.ReadString('\n')
    args := strings.Fields(text)

    if len(args) == 0 { continue }

    cmd := commands.Get(args[0])
    if cmd == nil { continue }

    cmd(args[1:], transactionManager)
  }
}


