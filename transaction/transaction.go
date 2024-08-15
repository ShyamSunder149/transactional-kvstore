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

func (tm *TransactionManager) checkActiveTransaction() bool {
  if(tm.size == 0) {
    fmt.Println("You are not inside a transaction")
    return false 
  }
  return true
}

func (tm *TransactionManager) CurrentStore() store.Store {
  if tm.size == 0 { return tm.globalStore }
  return tm.top.store 
}

func (tm *TransactionManager) GetCurrentTop() *Transaction {
  if tm.top == nil { return nil }
  return tm.top
}

func (tm *TransactionManager) Begin() {
  temp := &Transaction{ store : store.NewMemoryStore() }
  temp.next = tm.top
  tm.top = temp
  tm.size++
}

func (tm *TransactionManager) Commit() {
  if(!tm.checkActiveTransaction()) { return }

  if(tm.GetCurrentTop().next == nil) { 
    commitChanges(tm.GetCurrentTop().store, tm.globalStore)
  } else {
    commitChanges(tm.GetCurrentTop().store, tm.GetCurrentTop().next.store)
  }

  tm.top = tm.top.next
  tm.size--
}

func commitChanges(from, to store.Store) {
  for k, v := range from.GetMap() { to.Set(k, v) }
}

// as of now rollback is done in a way to reverse all the entries in the map of the transaction 
func (tm *TransactionManager) Rollback() {
  if(!tm.checkActiveTransaction()) { return }
  clear(tm.GetCurrentTop().store.GetMap())
}

func (tm *TransactionManager) End() {
  if(!tm.checkActiveTransaction()) { return }
  tm.top = tm.top.next 
  tm.size-- 
}

