package dirMod

import (
    "fmt"
    "os"
    "errors"
    "path/filepath"
)

var savings string = "savings"
var current string = "current"

func CreateSavingsAndCurrentDirectoriesToStoreTheDB()bool{
    if _, err := os.Stat(savings); os.IsNotExist(err){
        err1 := os.MkdirAll(savings, 0755)
        if err1 != nil{
            fmt.Println("While creating an savings & an current directory account we found : ", err1)
            return
        }
        fmt.Println("Created an savings account successfully")
    }else{
        fmt.Println("savings account already exists")
    }


    if _, err := os.Stat(current); os.IsNotExist(err){
        err2 := os.MkdirAll(current, 0755)
        if err2 != nil{
            fmt.Println("While creating an savings & an current directory account we found : ", err2)
            return
        }
        fmt.Println("Created an current account successfully")
    }else{
        fmt.Println("current account already exists")
    }

    return
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
