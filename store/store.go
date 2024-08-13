package store 

import (
  "fmt"
)

type Store interface {
  Get(k string)
  Set(k string, v string)
  Delete(k string)
  Count()
  GetMap() map[string]string 
}

type MemoryStore struct { data map[string]string }

func NewMemoryStore() *MemoryStore {
  return &MemoryStore {
    data : make(map[string]string),
  }
}

func (ms *MemoryStore) Get (k string) {
  v, ok := ms.data[k] 
  if !ok {
    fmt.Println("Key unavailable")
    return
  }

  fmt.Printf("Key %s : %s\n", k, v)
}

func (ms *MemoryStore) GetMap() map[string]string { return ms.data }

func (ms *MemoryStore) Set(k, v string) { ms.data[k] = v }

func (ms *MemoryStore) Delete (k string) {
  if _, ok := ms.data[k]; !ok {
    fmt.Println("Key unavailable")
    return 
  }  

  delete(ms.data, k)
}

func (ms *MemoryStore) Count() {
  fmt.Println(len(ms.data))
}

