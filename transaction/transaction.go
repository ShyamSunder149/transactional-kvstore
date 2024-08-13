package transaction

import (
  "transactional-kvstore/store"
  "fmt"
)

type Transaction struct {
  store     store.Store 
  next      *Transaction
}

type TransactionManager struct {
  globalStore  store.Store 
  top          *Transaction
  size         int
}

func NewTransactionManager() *TransactionManager {
  return &TransactionManager { 
    globalStore : store.NewMemoryStore(),
    top : nil,
    size : 0,
  }
}

func (tm *TransactionManager) Begin() {
  temp := &Transaction{ store : store.NewMemoryStore() }
  temp.next = tm.top
  tm.top = temp
  tm.size++
}

func (tm *TransactionManager) Commit() {
  if tm.size == 0 {
    fmt.Println("No Active Transaction")
    return 
  }

  if tm.top.next == nil {
    commitChanges(tm.top.store, tm.globalStore)
    return 
  }

  commitChanges(tm.top.store, tm.top.next.store) 
  tm.top = tm.top.next
  tm.size--
}

func commitChanges(from, to store.Store) {
  for k, v := range from.GetMap() { to.Set(k, v) }
}

// as of now rollback is done in a way to reverse all the entries in the map of the transaction 
func (tm *TransactionManager) Rollback() {
  if tm.size == 0 {
    fmt.Println("No Active Transaction")
    return  
  }

  clear(tm.top.store.GetMap())
}

func (tm *TransactionManager) End() {
  if tm.size == 0  {
    fmt.Println("No Active Transaction")
    return 
  }

  tm.top = tm.top.next 
  tm.size-- 
}

func (tm *TransactionManager) CurrentStore() store.Store {
  if tm.size == 0 { return tm.globalStore }
  return tm.top.store 
}

func (tm *TransactionManager) GetCurrentTop() *Transaction {
  if tm.top == nil { return nil }
  return tm.top
}
