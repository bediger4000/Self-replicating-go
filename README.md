# Self-replicating-go
Self-replicating Go program

Go's fmt formatted I/O package has an interesting format specification, "%q".
Go also has two ways to indicate "string literal", the usual C-style double-quoted string,
and a back-tick ("grave accent") quoted string. Double-quoted strings get examined for escaped characters, like "\n" for newline, etc.
Back-quoted strings don't get examined for escaped characters, and you can embed single and double quote characters in them.

Between the "%q" format specification and the back-quoted string literals, writing a self-replicating program in Go becomes a good deal simpler than in C.
