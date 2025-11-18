package dirMod

import (
    "fmt"
    "os"
    "errors"
)

func CreateSavingsAndCurrentDirectoriesToStoreTheDB(){
    err1 := os.Mkdir("saving", 0755)
    err2 := os.Mkdir("current", 0755)
    err := errors.Join(err1, err2)
    if err != nil{
        fmt.Println("While creating an savings & an current account we found : ", err)
        return
    }
    fmt.Println("Created an savings and an current account")
}
