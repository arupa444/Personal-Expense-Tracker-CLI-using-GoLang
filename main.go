package main

import (
    "fmt"
    "os"
)

func main(){
    files, err := os.ReadDir("Utils")
    if err != nil{
        fmt.Println("Error while creating a dict : ",err)
        return
    }
    fmt.Println("Created myDir successfully and the files are", files)
}