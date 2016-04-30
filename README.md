# Self-replicating Go program

Go's `fmt` formatted I/O package has an interesting format specification, "%q". Suppose your Go program has composed a string value that contains newlines, double-quote characters or other escapable characters. Pass such a string as an argument to one of the `fmt` functions (like `fmt.Printf()`) and use "%q" as the output specification. `fmt.Printf()` will escape all escapable characters in the  resulting output.

Go also has two ways to indicate "string literal", the usual C-style double-quoted string,
and a back-tick ("grave accent") quoted string. Double-quoted strings get examined for escaped characters, like "\n" for newline, etc.
Back-quoted strings don't get examined for escaped characters, and you can embed single and double quote characters in them.

Between the "%q" format specification and the back-quoted string literals, writing a self-replicating program ("quine") in Go becomes a good deal simpler than in C.

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

## Step 2 - inital format string

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

The above program outputs the first 5 lines of a generic Go program, twice, separated by a goofy warning message from `fmt.Printf()`. No format specifcation appears in the contents of `h`.
