package dirMod

import (
    "fmt"
    "os"
    "errors"
    "path/filepath"
)

var savings string = "savings"
var current string = "current"

func CreateSavingsAndCurrentDirectoriesToStoreTheDB(){
    fmt.Println(os.IsExist())
    err1 := os.MkdirAll(savings, 0755)
    err2 := os.MkdirAll(current, 0755)
    err := errors.Join(err1, err2)
    if err != nil{
        fmt.Println("While creating an savings & an current directory account we found : ", err)
        return
    }
    fmt.Println("Created an savings and an current account successfully")
}

func CreateSavingsAndCurrentFilesToStoreTheDB(){
    createFile := filepath.Join(savings, "savingsAccountDB.json")
    file1, err1 := os.Create(createFile)
    createFile = filepath.Join(current, "currentAccountDB.json")
    file2, err2 := os.Create(createFile)
    err := errors.Join(err1, err2)
    if err != nil{
        fmt.Println("while creating an savings and current file we found : ", err)
        return
    }
    defer file1.Close()
    defer file2.Close()
    fmt.Println("Created files successfully")
}
