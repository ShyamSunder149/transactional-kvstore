package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

var GlobalStore = make(map[string]string)

type Transaction struct {
  store   map[string]string 
  next    *Transaction
}

type TransactionStack struct {
  top     *Transaction
  size    int
}

func (ts *TransactionStack) PushTransaction() {
  temp := Transaction{ store : make(map[string]string) }
  temp.next = ts.top 
  ts.top = &temp
  ts.size++
}

func (ts *TransactionStack) PopTransaction() {
  if ts.size == 0 {
    fmt.Println("No Active Transactions")
    return 
  }

  node := &Transaction{}
  ts.top = ts.top.next 
  node.next = nil 
  ts.size--
}

func (ts *TransactionStack) Peek() *Transaction {
  return ts.top 
}

func (ts *TransactionStack) Commit() {

  ActiveTransaction := ts.Peek()
  if ActiveTransaction == nil {
    fmt.Println("Nothing to commit")
    return 
  }

  store := ActiveTransaction.next.store  
  if store == nil {
    for k, v := range ActiveTransaction.store { GlobalStore[k] = v }
    return 
  }
  
  for k, v := range ActiveTransaction.store { store[k] = v }
}

func (ts *TransactionStack) RollBackTransaction() {

  if ts == nil {
    fmt.Println("You are not inside a transaction")
    return 
  }

  store := ts.Peek().store 
  if len(store) == 0 {
    fmt.Println("Nothing to rollback")
    return 
  }

  for k := range store { delete(ts.Peek().store, k) }
}

func Get(key string, ts *TransactionStack) {
  
  store := GlobalStore

  if ts != nil && ts.Peek() != nil { store = ts.Peek().store }
  if _, ok := store[key]; !ok { 
    fmt.Printf("Key %s unavailable\n", key)
    return 
  }

  fmt.Println(store[key])
} 

func Set(key string, val string, ts *TransactionStack) {
  
  if ts == nil || ts.Peek() == nil {
    fmt.Println("You are not inside a transaction")
    return 
  }

  store := ts.Peek().store 
  store[key] = val 
}

func Count(ts *TransactionStack) {
  
  store := GlobalStore

  if ts != nil && ts.top != nil { store = ts.Peek().store }

  fmt.Println(len(store))
}

func Delete(key string, ts *TransactionStack) {
  
  store := GlobalStore

  if(ts != nil && ts.top != nil) { store = ts.Peek().store }
  if _, ok := store[key]; !ok { 
    fmt.Printf("Key %s unavailable\n", key)
    return
  }
  
  delete(store, key)
}

func checkTransactionAndExecute(ts *TransactionStack, method func()) {
  if ts.size == 0 {
    fmt.Println("You are not inside a Transaction")
    return 
  }

  method()
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  items := &TransactionStack{}
  for {
    fmt.Printf("> ")
    text, _ := reader.ReadString('\n')
    operation := strings.Fields(text)
    switch operation[0] {
      case "BEGIN": items.PushTransaction()
      case "ROLLBACK" : checkTransactionAndExecute(items, items.RollBackTransaction)
      case "COMMIT" : checkTransactionAndExecute(items, items.Commit)
      case "END" : checkTransactionAndExecute(items, items.PopTransaction)
      case "SET" : Set(operation[1], operation[2], items)
      case "GET" : Get(operation[1], items)
      case "DELETE" : Delete(operation[1], items)
      case "COUNT" : Count(items)
      case "QUIT" : os.Exit(0)
      default : fmt.Println("Command Not Found")
    }
  }
}
