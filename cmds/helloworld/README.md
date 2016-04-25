
# Hello World

Hello World is a traditional program to typing in when first learning a new programming language.
Here's what Hello World looks like in Go.

```go
  package main

  import "fmt"

  func main() {
  	fmt.Println("Hello World!")
  }
```

It has three essentials blocks you'll see in most Go programs.  

1. A package statement
2. An import statement (listing the packages used in this package)
3. A "main" function that provides a hook for the operating system to run our program.

You'll not that the fist line `package main` sets the name of our package.  The name
is special in the sense that this lets the Go compiler know we are building a executable
program.  If you writing a `package main` you will need a *main* function for your
program to do anything useful.  Actually you'll probable need more then that.

The second block is the `import "fmt"` statement. In this block we list one or more
packages that we need for this package to work.  In our case we're going to use the
*fmt* package to display our message in the terminal. You can include multiple
packages in one statement. If I need both the *fmt* package and say the *time*
package the import block might look something like

```go
  import (
    "fmt"
    "time"
  )
```

Notice that we group the multi-line package block win parenthesis with one package
name on each line in quotes.  Typically the standard packages that come with go are
one or two works (two or more words are separated with a "/") enclosed in quotes.
This notation lets the Go compile have a name space of sorts. Go uses the quoted string
name to find the source code for the package. The path described by the string needs to be
unique based on the way you've installed Go (more about this later).

The final part of the *main* function. It's pretty simple. I am using the format package
to print a line of text to the console.

Here's where you can find out more:

* [A Tour of Go](https://tour.golang.org/welcome/1)
* [Learn Go](https://github.com/golang/go/wiki/Learn) - a list of learning resources for Go
* [Introduction to Programming in Go](http://www.golang-book.com/) - aka "the golang book"
  * [packages](http://www.golang-book.com/books/intro/11)
* [Godocs on fmt](https://godoc.org/fmt)
