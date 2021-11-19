This readme would be tidied up after I finish reading both of the books. Sorry for the mess for now!

# Monkey Compiler In Go

A compiler written in Go. The language built is called [Monkey](https://monkeylang.org/).

I followed these two sequential books.

- [Writing An Interpreter In Go (Thorsten Ball)](https://interpreterbook.com/)
- [Writing A Compiler In Go (Thorsten Ball)](https://compilerbook.com/)

<img src="https://user-images.githubusercontent.com/44487754/138540965-52b709f7-d4d1-4c96-81f0-ad3de144d041.png" height="300"><img src="https://user-images.githubusercontent.com/44487754/138540951-41167952-9f0d-49ff-8889-57daa7fba2d6.png" height="300"><img src="https://user-images.githubusercontent.com/44487754/138540981-d84fe021-86fd-41d3-8587-7070b101d769.png" height="300">

For "Writing An Interpreter In Go", I read the Japanese translated version, which is "[Go 言語でつくるインタプリタ](https://www.oreilly.co.jp/books/9784873118222/)".

## Part1. Writing An Interpreter In Go

I created a Monkey interpreter (REPL) with this book.
PR#1 ~ PR#15 are in the scope of this book, and the version [2a928ad](https://github.com/kudojp/MonkeyInterpreter-Golang2021/commit/2a928adc2255b07605bea252dfc929a79115f171) is the completed one.

Note that package `monkey/evaluator` is modified later in part2.

## Part2. Writing A Compiler In Go (WIP)

I created a Monkey compiler + VM with this book.

- Byte code = opcode + operands
  - Constants are held as an object in a constant pool (slice) and the byte code has an index of that as an operand (instead of embedding the values directly in the bytecode).

Compiler:

- creates bytecode from ast
- creates a constant pool as an array of objects

VM:

- creates objects from bytecode
- uses the constant pool built in the compiler
