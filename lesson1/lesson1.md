## What is golang


- Statically linked
- Statically typed
- Compiles to native machine code
- Built in concurrency primitives 
- No classes OO works with type embedding

## Advantages

- Fast compilation.
- Great scalability. 
- Simplicity.
- Multi platform


## Installation
### mac
    brew install go
    
### ubuntu
    sudo apt-get install golang-go


### Running a program

    package main
    
    import "fmt"
    
    func main() {
        fmt.Println("Hello world")
    }

<!-- -->

    go run hello.go
    
    or 
To create a runnable binary

    go build hello.go
    
    
## Basic data structures.

### Built in types.

    bool
    string
    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr
    byte // Same as uint8
    rune - A readable unicode character.
    float32 float64
    complex64 complex128

<!-- -->
    var foo int
    var foo int = 2
    foo := 4 // Automatically type inference.
    const constant = "Hello"
    
More info on variable declaration can be found here:
    https://blog.golang.org/gos-declaration-syntax
    
    
    
### Arrays.

Arrays have fixed length.

    var a [10]int // Array with length 10
    a[1] = 1 
    b := a[1]    
    var a = [2]int{1, 2} // Type inference.


### Slices

In slices, the length is not specificed. Arrays serve as the underlying value of this type.

    var a []int                              
    var a = []int {1, 2, 3}           
    a := []int{1, 2, 3} // Type inference.

    // Looping over arrays and slices.
    for i, v := range a {
        // i is in the index, v is the value
    }
    
    a =  append(a,2,3)

Another way to create a slice is by using the "make" keyword, this creates an array and returns a slice that references to that array

A slice can be copied with the built in copy function:
    
    func copy(dst, src []T) int

Appending to a slice is done with the append function. if a greater slice capacity is needed, it will reallocate a new underlying array.
    
    func append( s []T, x ...T) []T
    

you can read more about this mechanism here:

https://blog.golang.org/go-slices-usage-and-internals

    a = make([]int, len(int), capacity(int))


### Maps
    
    m = make(map[string]int)
    var m map[string]int
    m["a"] = 1
    delete(m, "a")
    
    elem, _ := m["key"] Retrieve the "key" element. if _ is replaced with a val, it will return a boolean is the key exists.
    
    // Literal.
    aaa := map[string]int{
        "hello":1,
        "bye":2,
    }
    

### Control Flow
    
    if name == "ohad {
        return "hello"
    } else {
        return "bye"
    }
    
More info can be found here:
https://blog.golang.org/go-maps-in-action


### Loops
Go has no while loop and no do/while loop. we use "for" loop only.

        // Classic for loop
        for i := 1; i < 5; i++ {}
        
        //While loop.
        for i < 5  {}
        
        // While true.
        for {}
        
### Functions

In golang the type goes after the parameter name and the return value type is in the end of the function.
    
    func func1(param1 string, param2 int) {}
    // Same Type + Return value type.
    func func2(param1, param2 int) (int) {}
    
    // Multiple return values.
    func func3() (int, string) {
        return 1, "ohad"
    }

Functions can have variadic arguments

    func printlist(args ...string) {
         for _, v:= range args {
             fmt.Println(v)
        }
    }
    
    printlist("Hello", "bye")


### Casting

    var ohad int = 3
    var fdfd float64 = float64(i)
    
### Packges

Packages are imported at the top of the file.
like this

    import "fmt"

or multiline

    import (
        "fmt"
        "net/http" // this will be imported as "http"
    )
    
Lower case functions will be private
Upper case functions will be public 

