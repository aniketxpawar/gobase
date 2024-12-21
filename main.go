package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aniketxpawar/gobase/db"
)

func main() {
	fmt.Println("Welcome to Gobase")
    database, err := db.NewDatabase("data.json")
    if err != nil {
        fmt.Println("Error Initializing Database: ",err)
        return
    }

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("> ")
        query,_ := reader.ReadString('\n')
        response, err := database.ExecuteQuery(query)
        if err != nil {
            fmt.Println("Error: ", err)
        } else{
            fmt.Println(response)
        }
    }

}