package dirMod

import (
    "fmt"
    "os"
    "errors"
)

var savings string = "savings"
var current string = "current"

func CreateSavingsAndCurrentDirectoriesToStoreTheDB(){
    err1 := os.MkdirAll(savings, 0755)
    err2 := os.MkdirAll(current, 0755)
    err := errors.Join(err1, err2)
    if err != nil{
        fmt.Println("While creating an savings & an current account we found : ", err)
        return
    }
    fmt.Println("Created an savings and an current account")
}
