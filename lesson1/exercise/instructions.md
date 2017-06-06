Create a program that acts like a phonebook

gets Data from and phone number from the user, adds to a list and retrieves it

command syntax can be done as following.

save ohad 123456

get ohad -> expected result "123456"


Tip:

to read a string from the standard input you need to use the bufio which contains a scanner struct.

    package main
    
    import (
        "bufio"
        "fmt"
        "os"
    )
    
    func main() {
        reader := bufio.NewScanner(os.Stdin)
        _ = reader.Scan()
        text := reader.Text()
        fmt.Println(text)
    }
