package dirMod

import (
    "fmt"
    "os"
    "path/filepath"
)

var savings string = "savings"
var current string = "current"

func CreateSavingsDirectoriesToStoreTheDB(){
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
}


func CreateCurrentDirectoriesToStoreTheDB(){
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
}


func CreateSavingsFilesToStoreTheDB(){
    createFile := filepath.Join(savings, "savingsAccountTransitionsDB.json")
    file, err := os.Create(createFile)
    if err != nil{
        fmt.Println("while creating a savings file we found : ", err)
        return
    }
    defer file.Close()
    fmt.Println("Created savings account files successfully")

    createFile = filepath.Join(savings, "infoAboutYourSavingAccount.json")
    file, err = os.Create(createFile)
    if err != nil{
        fmt.Println("while creating a saving info account file we found : ", err)
        return
    }
    defer file.Close()
    fmt.Println("Created saving info account files successfully")
}


func CreateCurrentFilesToStoreTheDB(){
    createFile := filepath.Join(current, "currentAccountTransitionsDB.json")
    file, err := os.Create(createFile)
    if err != nil{
        fmt.Println("while creating an current file we found : ", err)
        return
    }
    defer file.Close()
    fmt.Println("Created current account files successfully")

    createFile = filepath.Join(current, "infoAboutYourCurrentAccount.json")
    file, err = os.Create(createFile)
    if err != nil{
        fmt.Println("while creating an current info account file we found : ", err)
        return
    }
    defer file.Close()
    fmt.Println("Created current info account account files successfully")
}


