# Self-replicating Go program
## And bonus narcissistic program!

Go's `fmt` formatted I/O package has an interesting verb, "%q". Suppose your Go program has composed a string value that contains newlines, double-quote characters or other escapable characters. Pass such a string as an argument to one of the `fmt` functions (like `fmt.Printf()`) and use "%q" as the verb. `fmt.Printf()` will escape all escapable characters in the resulting output.

Go also has two ways to indicate "string literal", the usual C-style double-quoted string,
and a back-tick ("grave accent") quoted string. Double-quoted strings get examined for escaped characters, like "\n" for newline, etc.
Back-quoted strings don't get examined for escaped characters, and you can embed single and double quote characters in them.

Between the "%q" verb and the back-quoted string literals, writing a self-replicating program ("quine") in Go becomes a good deal simpler than in C.

This quine is going to look like the "associate of a number" that Raymond Smullyan has fun with in the 
[Monte Carlo lock](https://bediger4000.github.io/mcm.html)
section of [The Lady or the Tiger](http://www.amazon.com/The-Lady-Tiger-Other-Puzzles/dp/048647027X/).
The combination to the Monte Carlo lock involves symbols that are effectively [Concatenative Combinators](http://tunes.org/~iepos/joy.html). The "2" symbol quotes without evaluating the rest of the lock's combination, and the "3" symbol returns the "associate" of the rest of the lock's combination. The associate of a string is just the concatenation of "string"2"string", so "323" is the self-replicating Monte Carlo lock combination. This Go quine is going look like `fmt.Printf(str, str)`, where `str` is both a valid `Printf` format string, and the argument so formatted.

You could also interpret this form as the Combinatory Logic term `M M`.
There's a hidden application operation in the term, so `M M` actually ends up as `Apply(M, M)`.

## Step 1 - Go necessities

Go the language is pretty militant about some minimal formatting.
We start with some boilerplate,
and the "associate" of something:

```go
    package main
    
    import "fmt"
    
    func main() {
        h := `something`
        fmt.Printf(h, h)
    }
````

Note that the "something" string is a back-quoted string.
We can put newlines,
double quotes, whatever in it,
and the Go compiler will just create an identical string literal.

If we use the "%q" format specification, we won't have to worry about "escaping hell"
when we write the contents of `h`.
The contents of `h` are going to be a `fmt.Printf()` format string.
We know we have to have the program output the "boilerplate" Go.

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

Informed by the output of the program from Step 3, we can fill in the back-quoted string literal.
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
The final-format-string version should reside in the `rx.go` file in this repository.

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

The actual self-replicating program (`ry.go` or `rz.go` in the example above)
is almost identical to
[mhilton's Go quine](https://github.com/mhilton/quine/blob/master/quine.go),
differing only in the name of the sole variable.
It appears that is the form of the minimum,
officially formatted, self-replicating Go program.

## Bonus Narcissist Program

[Narcissist programs](https://rosettacode.org/wiki/Narcissist) read an input, and output a
"1" if the input matches the source code of the Narcissist, and output a "0" if it doesn't
match the source code.

My Narcissist program is closely related to the self-replication program.
Using golang's backquoted literals,
I modified `rx.go` to generate a program.
The generated program re-creates its own source in the manner of a
self-replicating program,
except keeping that source in a `string` type variable.
Then it compares bytes on stdin to the copy of source code,
quitting on byte value mismatches,
input too long after a match,
or input errors.
If the Narcissist hits end-of-file on stdin
after matching all input bytes to source bytes,
it outputs a "1" character.
Otherwise, it outputs a "0" character.

To create and try a Narcissist program:

    $ make
    ...
	$ ./narcissist < narcissist.go
    1
    $

### Almost Narcissist  Program

I created an "Almost Narcissistic" program as well. Instead of checking for
file identity by comparing bytes and file lengths, an Almost-Narcissist reads
bytes and calculates a checksum, or hash or CRC value. Using that value, it
makes a decent guess at whether its input is the same as its source code.

I chose CRC32 because I had a vague recollection that you could easily
create a CRC32 collision,
where other hashes were much harder to create a collision.
My Almost-Narcissist creates its own source code in-memory,
just as the Narcissist program does.
It calculates a CRC32 value for those in-memory bytes.
It then calculates a CRC32 for whatever it reads on stdin,
until end-of-file.
If the CRC32 values match, the input bytes are _probably_ the source code of the program.
Since one can [generate CRC32 collisions](https://github.com/bediger4000/crc32-file-collision-generator)
with ease,
one can readily create a file that fools the Almost-Narcissist.
Using a better hash,
the Almost-Narcissist would get better at recognizing its own source.

To create and try my Almost-Narcissist program:

    $ make almost_narcissist
    ...
	$ ./almost_narcissist < almost_narcissist.go
    1
    $

### Self-Encrypting Program

I saw [this web page](https://evervault.com/blog/this-code-encrypts-itself)
on self-encrypting programs.
These are clearly a variant of plain old self-replication,
and the narcissist program.

I wrote a Go version that outputs [Base 64-encoded](https://en.wikipedia.org/wiki/Base64)
version of itself.
Here, Base 64-encoding stands in for whatever encryption you might wish to apply.
I'm lazy, I didn't want to generate some key pair and go to the trouble
of learning how to decrypt some complicated format.

```
$ make self_encrypting
...
$ ./self_encrypting > x.b64
$ base64 -d x.b64 > x.go
$ go build x.go
$ ./x > y.b64
$ cksum x.b64 y.b64
...
$ diff self_encrypting.go x.go
```

The program `self_encrypting` outputs a Base 64
encoded version of its own source code.
You have to decode the Base-64-encoded source code to compile it,
and then compile that decoded Base-64-encoded source code.
