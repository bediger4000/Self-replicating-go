# Self-replicating Go program

Go's `fmt` formatted I/O package has an interesting verb, "%q". Suppose your Go program has composed a string value that contains newlines, double-quote characters or other escapable characters. Pass such a string as an argument to one of the `fmt` functions (like `fmt.Printf()`) and use "%q" as the verb. `fmt.Printf()` will escape all escapable characters in the resulting output.

Go also has two ways to indicate "string literal", the usual C-style double-quoted string,
and a back-tick ("grave accent") quoted string. Double-quoted strings get examined for escaped characters, like "\n" for newline, etc.
Back-quoted strings don't get examined for escaped characters, and you can embed single and double quote characters in them.

Between the "%q" verb and the back-quoted string literals, writing a self-replicating program ("quine") in Go becomes a good deal simpler than in C.

This quine is going to look like the "associate of a number" that Raymond Smullyan has fun with in the Monte Carlo lock section of [The Lady or the Tiger](http://www.amazon.com/The-Lady-Tiger-Other-Puzzles/dp/048647027X/). The combination to the Monte Carlo lock involves symbols that are effectively [Concatenative Combinators](http://tunes.org/~iepos/joy.html). The "2" symbol quotes without evaluating the rest of the lock's combination, and the "3" symbol returns the "associate" of the rest of the lock's combination. The associate of a string is just the concatenation of "string"2"string", so "323" is the self-replicating Monte Carlo lock combination. This Go quine is going look like `fmt.Printf(str, str)`, where `str` is both a valid `Printf` format string, and the argument so formatted.

You could also interpret this form as the Combinatory Logic term `M M`. There's a hidden application operation in the term, so `M M` actually ends up as `Apply(M, M)`.

## Step 1 - Go necessities

Go the language is pretty militant about some minimal formatting. We start with some boilerplate, and the "associate" of something:

```go
    package main
    
    import "fmt"
    
    func main() {
        h := `something`
        fmt.Printf(h, h)
    }
````

Note that the "something" string is a back-quoted string. We can put newlines, double quotes, whatever in it, and the Go compiler will just create an identical string literal.

If we use the "%q" format specification, we won't have to worry about "escape hell" when we write the contents of `h`. The contents of `h` are going to be a `fmt.Printf()` format string. We know we have to have the program output the "boilerplate" Go.

## Step 2 - Initial format string

```go
    package main
    
    import "fmt"
    
    func main() {
        h := `package main
        
        import "fmt"
        
        func main() {
        `
        fmt.Printf(h, h)
    }
```

The above program outputs the first 5 lines of a generic Go program, twice, separated by a goofy warning message from `fmt.Printf()`. No verb appears in the contents of `h`. If we insert a verb into the contents of `h` we get the next intermediate form.

## Step 3 - Intermediate format string

```go
    package main
    
    import "fmt"
    
    func main() {
        h := `package main
        
        import "fmt"
        
        func main() {
            h := %q
    `
        fmt.Printf(h, h)
    }
```
The output of this program is quite close to its own source code:

```go
package main

import "fmt"

func main() {
        h := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\th := %q\n"
```
We see that the "%q" format string verb has escaped every special character, making the contents of `h` difficult to read.

We indirectly see that the initial program's back-quoted string literal makes writing that string literal almost easy. The output of the intermediate-format-string program is only missing the `fmt.Prinft()` statement, and the final curly brace to end the `main()` function. That's easy to add to the string literal.

## Step 4 - Final format string

```go
    package main
    
    import "fmt"
    
    func main() {
        h := `package main
        
        import "fmt"
        
        func main() {
            h := %q
            fmt.Printf(h, h)
        }
        `
        fmt.Printf(h, h)
    }
```
The final-format-string version is the `rx.go` code in this repository.

The output of the final-format-string program constitutes the self-replicating program.

## Step 5 - Generate self-replicating program

The file named "rx.go" in this repository, when compiled and executed, writes the source code of a
self-replicating program on stdout.

    % go build rx.go
    % ./rx > ry.go
    % go build ry.go
    % ./ry > rz.go
    % diff ry.go rz.go
    %    
Diffing `rx.go` and `ry.go` shows you the difference between the generator and the self-replicator.



