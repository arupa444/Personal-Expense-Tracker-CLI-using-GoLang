package main

import (
    "fmt"
    "os"
	"path/filepath"
	"CLIPacks/Utils"
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

	_, err = file.WriteString("package trio\n\nimport(\n\t\"fmt\"\n)\n\nfunc Loon(){\nfmt.Println(\"Hey bro i coded it...\")\n}")
	if err != nil{
	    fmt.Println("Error while writing", err, file)
	    return
	}
    fmt.Println("Created and wrote some syntax's on the create file")
    trio.Loon()
}