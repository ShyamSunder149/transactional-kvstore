 package commands 

 import (
    "fmt"
    "os"
    "transactional-kvstore/transaction"
 )

type Command func([] string, *transaction.TransactionManager) 

var commands = map[string] Command {
  "BEGIN" : beginTransaction,
  "ROLLBACK" : rollbackTransaction,
  "COMMIT" : commitTransaction,
  "END" : endTransaction,
  "SET" : setValue,
  "GET" : getValue,
  "DELETE" : deleteValue,
  "COUNT" : countValues,
  "QUIT" : quitProgram,
}

func Get(name string) Command {
  cmd, ok := commands[name]
  if !ok {
    fmt.Println("Command Not found")
    return nil
  }
  return cmd 
}

func beginTransaction (_ []string, tm *transaction.TransactionManager) { tm.Begin() }

func rollbackTransaction (_ []string, tm *transaction.TransactionManager) { tm.Rollback() }

func commitTransaction (_ []string, tm *transaction.TransactionManager) { tm.Commit() }

func endTransaction (_ []string, tm *transaction.TransactionManager) { tm.End() }

func setValue(args []string, tm *transaction.TransactionManager) {
  if(len(args) != 2) {
    fmt.Println("No of arguments is incorrect")
    return 
  }
  
  if tm.GetCurrentTop() == nil {
    fmt.Println("You are not inside a transaction")
    return 
  }

  tm.CurrentStore().Set(args[0], args[1])
}

func getValue(args []string, tm *transaction.TransactionManager) {
  if len(args) != 1 {
    fmt.Println("No of Arguments is incorrect")
    return 
  }

  tm.CurrentStore().Get(args[0])
}

func deleteValue(args []string, tm *transaction.TransactionManager) {
  if len(args) != 1 {
    fmt.Println("No of Arguments is incorrect")
    return
  } 

  tm.CurrentStore().Delete(args[0])
}

func countValues(_ []string, tm *transaction.TransactionManager) { tm.CurrentStore().Count() }

func quitProgram(_ []string, tm *transaction.TransactionManager) { os.Exit(0) }
