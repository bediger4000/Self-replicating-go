# Self-replicating Go program

Go's `fmt` formatted I/O package has an interesting format specification, "%q". Suppose your Go program has composed a string value that contains newlines, double-quote characters or other escapable characters. Pass such a string as an argument to one of the `fmt` functions (like `fmt.Printf()`) and use "%q" as the output specification. `fmt.Printf()` will escape all escapable characters in the  resulting output.

Go also has two ways to indicate "string literal", the usual C-style double-quoted string,
and a back-tick ("grave accent") quoted string. Double-quoted strings get examined for escaped characters, like "\n" for newline, etc.
Back-quoted strings don't get examined for escaped characters, and you can embed single and double quote characters in them.

Between the "%q" format specification and the back-quoted string literals, writing a self-replicating program in Go becomes a good deal simpler than in C.

