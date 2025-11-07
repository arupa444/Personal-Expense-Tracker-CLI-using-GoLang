package main

import (
    "fmt"
    "os"
)

func main(){
    err := os.Mkdir("myDir", 0755)
    if err != nil{
        fmt.Println("Error while creating a dict : ",err)
        return
    }
    fmt.Println("Created myDir successfully")
}