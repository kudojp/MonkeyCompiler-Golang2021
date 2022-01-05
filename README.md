<details>
<summary>Please also refer to my blog posts related to this project.</summary>

<br/>

  **Beautiful algorithms used in a compiler and a stack based virtual machine**

- Part 1. [Design of the Monkey compiler](https://www.wantedly.com/users/67312544/post_articles/363007)
- Part 2. [How a pratt parser in Monkey compiler works](https://www.wantedly.com/users/67312544/post_articles/364335)
- Part 3. [How to compile global bindings](https://www.wantedly.com/users/67312544/post_articles/365686)
- Part 4. [How to compile functions](https://www.wantedly.com/users/67312544/post_articles/367694)
- Part 5, [What Monkey does not have, but Java has](https://www.wantedly.com/users/67312544/post_articles/366601)

</details>

# Monkey Compiler In Go

## About this project

This is a compiler and a virtual machine of the [Monkey programming language](https://monkeylang.org/).
The design overview of this project is described [here](https://www.wantedly.com/users/67312544/post_articles/363007).

As a guide of this project, I referenced these two books written by [Thorsten Ball](https://thorstenball.com/).

- [Writing An Interpreter In Go](https://interpreterbook.com/)
- [Writing A Compiler In Go](https://compilerbook.com/)

<img src="https://user-images.githubusercontent.com/44487754/138540951-41167952-9f0d-49ff-8889-57daa7fba2d6.png" height="240"><img src="https://user-images.githubusercontent.com/44487754/138540965-52b709f7-d4d1-4c96-81f0-ad3de144d041.png" height="240">

<p align="right">(<a href="#top">back to top</a>)</p>

## Getting Started


### How to build and run

```
$ go build .
$ go run .
```
### How to test

```
$ go test ./...
```

<p align="right">(<a href="#top">back to top</a>)</p>


## Project History

In this project, Monkey was first implemented as an interpreter. Then, it has been updated to a compiler. Both of them are REPL style.

### [v1.0](https://github.com/kudojp/MonkeyCompiler-Golang2021/releases/tag/v1.0). Monkey Interpreter

The first version of Monkey was a tree walking interpreter.

PR#1 ~ PR#15 are the processes of implementation, where the lexer, the parser, and the evaluator are implemented. The algorithm used in the parser in documented [here](https://www.wantedly.com/users/67312544/post_articles/364335).

### [v2.0](https://github.com/kudojp/MonkeyCompiler-Golang2021/releases/tag/v2.0). Monkey Compiler

The second and the latest version of Monkey is a bytecode compiler and virtual machine. The previous version has been refactored and converted into this version in PR#16 and later.

The algorithms used for compiling and executing local variables are documented [here](https://www.wantedly.com/users/67312544/post_articles/365686), and those for functions are [here](https://www.wantedly.com/users/67312544/post_articles/367694).

### (Appendix) Benchmarking of the interpreter vs. the compiler

Here is the benchmark testing of the interpreter version and the compiler version.
The snippet below is executed by each of them.

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

The result reveals that the compiler version is **more than 3 times** as performant as the interpreter version.

```sh
# the interpreter version
$ go run benchmark/main.go -engine=eval
engine=eval, result=9227465, duration=37.105657782s

# the compiler version
$ go run benchmark/main.go -engine=vm
engine=vm, result=9227465, duration=11.787990726s
```

<p align="right">(<a href="#top">back to top</a>)</p>


## Contributing

If you have a suggestion that would make this project better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". Some possible improvements are listed [here](https://www.wantedly.com/users/67312544/post_articles/366601). Don't forget to give the project a star! Thanks again.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>
