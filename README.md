# Building an interpreter with GO

This is a simple interpreter written in GO. It is based on the book [Writing an interpreter in Go](https://interpreterbook.com/).

## What is an interpreter?

An interpreter is a program that executes code written in a high-level language. It does so by reading the source code and translating it on the fly to machine code. This is different from a compiler, which translates the source code to machine code before execution.

## What I'm building here

I'm building an interpreter for a language called Monkey. Monkey is a simple programming language that supports:

- Integers
- Booleans
- Strings
- Arrays
- Hashes
- Prefix, infix and index expressions
- Conditionals
- Global and local bindings
- First-class and higher-order functions
- Closures
- A string concatenation operator
- A built-in function to print to the console
- A built-in function to read user input from the console
- A built-in function to convert integers to strings

## Why am I building this?

I want to have a go-to project to build when I want to learn a new programming language. I've built this interpreter in Ruby and JavaScript before, so I thought it would be a good idea to build it in GO as well.

## What we are building

The interpreter we’re going to build in this book will implement all these features. It will
tokenize and parse Monkey source code in a REPL, building up an internal representation of
the code called abstract syntax tree and then evaluate this tree. It will have a few major parts:
- the lexer
- the parser
- the Abstract Syntax Tree (AST)
- the internal object system
- the evaluator

## The lexer

The lexer is the first component of an interpreter or compiler. It takes the source code as input and returns a sequence of tokens that the parser then uses to build up the AST. The lexer is also called a tokenizer.

Here’s an example. This is the input one gives to a lexer:

```go 
let x = 5 + 5;
```

And what comes out of the lexer looks kinda like this:

```go 
[
    LET,
    IDENTIFIER("x"),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
] 
```

The lexer takes the source code and returns a list of tokens. Each token has a type and a literal value. The type is an integer constant, and the literal value is the actual character or string that the token represents in the source code.

A thing to note about this example: whitespace characters don’t show up as tokens. In our case
that’s okay, because whitespace length is not significant in the Monkey language. Whitespace
merely acts as a separator for other tokens. It doesn’t matter if we type this:

`let x = 5;`

Or if we type this:

`let   x   =   5;`

In other languages, like Python, the length of whitespace is significant. That means the lexer
can’t just “eat up” the whitespace and newline characters. It has to output the whitespace
characters as tokens so the parser can later on make sense of them (or output an error, of
course, if there are not enough or too many).

## The parser

The parser takes the list of tokens from the lexer and turns it into an AST. The parser is also called a syntactic analyzer. The one we are going to build is top down, recursive and Pratt parser. It is also called a recursive descent parser. It is adviced for beginners to start so because it is how we usually think about parsing code. It is also easy to implement.

Here is a fully valid program written in Monkey:

```go 
let x = 10;
let y = 15;
let add = fn(a, b) {
return a + b;
};
```

Programs in Monkey are a series of statements. In this example we can see three statements,
three variable bindings - let statements - of the following form:

`let <identifier> = <expression>;`