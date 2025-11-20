package main

import (
    "github.com/arupa444/Personal-Expense-Tracker-CLI-using-GoLang/utils/createAndDeleteDir"
)

func main(){

//     To create the saving account dir
    dirMod.CreateSavingsDirectoriesToStoreTheDB()
//     To create the saving account file in the dir
    dirMod.CreateSavingsFilesToStoreTheDB()
//     To create the current account dir
    dirMod.CreateCurrentDirectoriesToStoreTheDB()
//     To create the current account file in the dir
    dirMod.CreateCurrentFilesToStoreTheDB()

}