package store 

import (
  "fmt"
)

type Store interface {
  Get(k string)
  Set(k, string, v string)
  Delete(k, string)
  Count()
}

type MemoryStore struct { store map[string]string }

func NewMemoryStore() *MemoryStore {
  return &MemoryStore {
    store : map[string]string
  }
}

func (ms *MemoryStore) Get (k string) {
  if v, ok := ms.Store[k]; if !ok {
    fmt.Println("Key unavailable")
    return
  }

  fmt.Printf("Key : %s", v)
}

func (ms *MemoryStore) Set (k string, v string) {
  ms.Store[k] = v
}

func (ms *MemoryStore) Delete (k string) {
  if _, ok := ms.Store[k]; !ok {
    fmt.Println("Key uunavailable")
    return 
  }  

  delete(ms.Store, k)
}

func (ms *MemoryStore) Count() {
  return len(ms.Store)
}

