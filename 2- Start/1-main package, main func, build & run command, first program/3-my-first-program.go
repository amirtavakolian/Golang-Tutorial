package main // your program must have one main package

import "fmt"

// your program must have one main func
func main() {
	fmt.Print("Hello World")
}

/*

# Remember:
you have to put your projects in "src" folder inside your workspace

run 'go env' to find the path of workspace inside the GOPATH

------------------------------------

use "go build" or "go run" to see the result of your program

======================================================

change func main to func hello then "go build" and "go run"

Error ==> runtime.main_main·f: function main is undeclared in the main package

it says, the main function is not declared in main package
so, you have to declare main function in main package.

main function is the entry point of your program.
the point, when you run your program, it will run from the main function.

-----------------------------------------

now, change "package main" to "package hello" then "go build -o ....."

the program will build (exe file will be created) but, that exe file will not run.

if you run ls -al, that exe file will not have the letter x at the end of -rwxr-xr--
which it shows that exe file is not execute-able.

so, if no main package in your program, the exe file will be created but will not execute
( it will not an executable file) ==> so you MUST have main package to get an executable file as output

*/
