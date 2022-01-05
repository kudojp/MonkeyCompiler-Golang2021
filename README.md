<details>
<summary>Please also refer to my blog posts related to this project.</summary>

- Part 1. [Design of the Monkey compiler](https://www.wantedly.com/users/67312544/post_articles/363007)
- Part 2. [How a pratt parser in Moneky compiler works](https://www.wantedly.com/users/67312544/post_articles/364335)
- Part 3. [How to compile global bindings](https://www.wantedly.com/users/67312544/post_articles/365686)
- Part 4. [How to compile functions](https://www.wantedly.com/users/67312544/post_articles/367694)
- Part 5, [What Monkey does not have, but Java has](https://www.wantedly.com/users/67312544/post_articles/366601)

</details>

# Monkey Compiler In Go

## About this project

This is a Monkey language 

This is a [Monkey](https://monkeylang.org/) is a programming language which has minimum features
compiler written in Go. The language is called [Monkey](https://monkeylang.org/). I referenced these two books.

- [Writing An Interpreter In Go (Thorsten Ball)](https://interpreterbook.com/)
- [Writing A Compiler In Go (Thorsten Ball)](https://compilerbook.com/)

<img src="https://user-images.githubusercontent.com/44487754/138540981-d84fe021-86fd-41d3-8587-7070b101d769.png" height="300"><img src="https://user-images.githubusercontent.com/44487754/138540951-41167952-9f0d-49ff-8889-57daa7fba2d6.png" height="300"><img src="https://user-images.githubusercontent.com/44487754/138540965-52b709f7-d4d1-4c96-81f0-ad3de144d041.png" height="300">

For "Writing An Interpreter In Go", I read the Japanese translated version, which is "[Go 言語でつくるインタプリタ](https://www.oreilly.co.jp/books/9784873118222/)".

<p align="right">(<a href="#top">back to top</a>)</p>


## Getting Started


### How to build and run

```sh
$ go build .
$ go run .
```
### How to test

```sh
go test ./...
```

<p align="right">(<a href="#top">back to top</a>)</p>


## Project History


### Monkey interpreter

Writing An Interpreter In Go

I created a Monkey interpreter (REPL) with this book.
PR#1 ~ PR#15 are in the scope of this book, and the version [2a928ad](https://github.com/kudojp/MonkeyInterpreter-Golang2021/commit/2a928adc2255b07605bea252dfc929a79115f171) is the completed one.

Note that package `monkey/evaluator` is modified later in part2.

### Monkey compiler

Part2. Writing A Compiler In Go

I created a Monkey compiler + VM with this book.

- Byte code = opcode + operands
  - Constants are held as an object in a constant pool (slice) and the byte code has an index of that as an operand (instead of embedding the values directly in the bytecode).

Compiler:

- creates bytecode from ast
- creates a constant pool as an array of objects

VM:

- creates objects from bytecode
- uses the constant pool built in the compiler

### (Appendix) Benchmark of the interpreter vs. the compiler

Performance of interpreter vs compiler

Run the snippet below with our interpreter/ our compiler to compare their performances.

```js
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      return 1;
    } else {
      fibonacci(x-1) + fibonacci(x-2)
    }
  }
}
fibonacci(35)
```

The result reveals that our compiler runs **3 times faster** than our interpreter.

```sh
$ go run benchmark/main.go -engine=vm
engine=vm, result=9227465, duration=11.787990726s

$ go run benchmark/main.go -engine=eval
engine=eval, result=9227465, duration=37.105657782s
```

<p align="right">(<a href="#top">back to top</a>)</p>


## Contributing

If you have a suggestion that would make this project better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>
