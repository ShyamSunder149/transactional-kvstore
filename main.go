package main

import (
  "fmt"
  "os"
)

var globalStore = make(map[string]string)
 
type Transaction struct {
  
  var store map[string]string
  var nextTransaction *Transaction

}

type TransactionStack struct {

  var currentTransaction *Transaction
  var size int 

}

type KVStore interface {
  
  Begin()
  Set()
  Get()
  Delete()

}

func Begin() *Transaction {
  return &Transaction(store : make(map[string]string), nextTransaction : nil,)  
}

func (ts *TransactionStack) addTransactionToStack(t *Transaction) {
  if ts.currentTransaction == nil { ts.currentTransaction = t } 
  else {
    t.nextTransaction = ts.currentTransaction
    ts.currentTransaction = t 
  }
  ts.size = ts.size + 1
}

func (ts *TransactionStack) removeTopOfStack() {
  ts.currentTransaction = ts.currentTransaction.nextTransaction
  ts.size = ts.size - 1
}

func (t *Transaction) Set(k, v string) {
  t.store[k] = v 
}

func (t *Transaction) Get(k string) {
  v, ok := t.store[k]
  if ok { fmt.Println("Value Unavailable") }
  else { fmt.Printf("%s : %s", k, v) }
}

func main() {

  fmt.Println("Konnichiwa, You are at K-V store v1")

  for true {
    var input string
    fmt.Printf("> ")
    fmt.Scanf("%s", &input)
    
    switch input {
      case "BEGIN" :  
      case "ROLLBACK" : 
      case "COMMIT" : 
      case "SET" : 
      case "GET" : 
      case "DELETE" : 
      case "END" :  
      case "COUNT" : fmt.Println(len(globalStore))
      case "PRINT KEYS" : 
      case "PRINT VALUES" : 
      case "PRINT KEY VALUE PAIRS" : 
      case "QUIT" : fmt.Println("Closing Session, Arigato")
                    os.Exit(0)
      default : fmt.Println("Enter a valid command")
    }
  }
}
