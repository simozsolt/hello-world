package main

import (
    "fmt"
    "github.com/simozsolt/hello-world/menu"
)

func main() {
    menu.PrintMenu()

    var selected string
    for selected != "q" {  // break the loop if text == "q"
        fmt.Print("Choose: ")
        fmt.Scanln(&selected)
        if selected != "q" {
            menu.Execute(selected)

            println()
            menu.PrintMenu()
        }
    }

    fmt.Println("Good by!")

    //fmt.Printf("Hello, world\n")
    //fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
