package main

import (
    "fmt"
    "os"
	"path/filepath"
)

func main(){
	utils := "Utils"
    err := os.MkdirAll(utils, 0755)
    if err != nil{
        fmt.Println("Error while creating a dict : ",err)
        return
    }

    filePath := filepath.Join(utils, "try.go")
    file, err := os.Create(filePath)
    if err != nil{
        fmt.Println("Error while creating a file : ",err)
        return
    }
	defer file.Close()
	_, err := file.WriteString("""package try
	import(
	    "os"
	    "fmt"
	)""")
	if err != nil{
	    fmt.Println("Error while writing", err)
	    return
	}

}