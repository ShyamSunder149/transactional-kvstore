package main

import (
  "fmt"
  "os"
  "strings"
)

var globalStore = make(map[string]string)
 
type Transaction struct {
  
  Store map[string]string
  nextTransaction *Transaction

}

type TransactionStack struct {

  currentTransaction *Transaction
  size int 

}

type KVStore interface {
  
  Begin()
  Set()
  Get()
  Delete()

}

func Begin() *Transaction { 
  return &Transaction {
    Store : make(map[string]string), 
    nextTransaction : nil,
  }
}

func (ts *TransactionStack) addTransactionToStack(t *Transaction) {
  if ts.currentTransaction == nil { 
    ts.currentTransaction = t 
  } else {
    t.nextTransaction = ts.currentTransaction
    ts.currentTransaction = t 
  }
  ts.size = ts.size + 1
}

func (ts *TransactionStack) removeTopOfStack() {
  ts.currentTransaction = ts.currentTransaction.nextTransaction
  ts.size = ts.size - 1
}

func (t *Transaction) Set(k, v string) { t.Store[k] = v }

func (t *Transaction) Get(k string) {
  if v, ok := t.Store[k]; ok {
    fmt.Printf("%d : %d", k, v)
  } else if v, ok := globalStore[k]; ok {
    fmt.Printf("%d : %d", k, v)
  } else {
    fmt.Println("Value Unavailable")
  }
}

func (t * Transaction) TransferElementsToPrevStore() {
  if t.nextTransaction == nil {
    for k, v := range t.Store { globalStore[k] = v }
  } else {
    temp := t.nextTransaction
    for k, v := range t.Store { temp.Store[k] = v } 
  }
}

// printing inside a Transaction will only get the Transaction store keys.
func (t *Transaction) PrintKeyValues() {
  for k, v := range t.Store { fmt.Printf("%s : %s\n", k, v) }
}

func (t *Transaction) PrintKeys() {
  for k, _ := range t.Store { fmt.Printf("%s\n", k) }
}

func PrintKeyValuesGlobal() {
  for k, v := range globalStore { fmt.Printf("%s : %s\n", k, v) }
}

func PrintKeysGlobal() {
  for k, _ := range globalStore { fmt.Printf("%s\n", k) }
}

func main() {

  fmt.Println("Konnichiwa, You are at K-V store v1")
  var txn Transaction
  var ts TransactionStack

  for true {
    var input string
    fmt.Printf("> ")
    fmt.Scanf("%s", &input)
    inputSplit := strings.Split(input, " ")
    
    var k, v string 

    token := inputSplit[0]
    if len(inputSplit) > 2 {
      k = inputSplit[1]
      v = inputSplit[2]
    } else if len(inputSplit) > 1 {
      k = inputSplit[1]
    } 
    
    switch token {
      case "BEGIN" : txn := Begin() 
      case "ROLLBACK" :  
      case "COMMIT" : 
      case "SET" : txn.Set(k, v)
      case "GET" : txn.Get(k)
      case "DELETE" : 
      case "END" :  txn.TransferElementsToPrevStore()
                    ts.removeTopOfStack() 
      case "COUNT" : fmt.Println(len(globalStore))
      case "PRINT KEYS" : if ts.size == 0 {
                            PrintKeysGlobal()
                          } else {
                            fmt.Println("Note : Within the transaction the keys of the transaction are alone printed")
                            txn.PrintKeys()  
                          }
      case "PRINT KEY VALUE PAIRS" : if ts.size == 0 {
                                        PrintKeyValuesGlobal()
                                     } else {
                                        fmt.Println("Note : Within the transaction the key values of the transaction are alone printed")
                                        txn.PrintKeyValues()
                                     }
      case "QUIT" : fmt.Println("Closing Session, Arigato")
                    os.Exit(0)
      default : fmt.Println("Enter a valid command")
    }
  }
}
