# Notes

> notes taken during the course

## Setup

```bash
go mod init github.com/filipe1309/mf-refactoring-go-examples
```

```bash
go run main.go
```

```bash
go build
./mf-refactoring-go-examples
```

```bash
go test
```

## Chapter 1

Refactoring change the internal structure of the code without changing its external behavior

```bash
go mod init github.com/filipe1309/mf-refactoring-go-examples/chapter_1
go mod tidy
```

```bash
go run before.go
```

The essence of the process of refactoring: small steps and tests after each step to ensure that the code is still working

Refactoring changes the program in small steps in a way that if you make a mistake, you can easily find it

Any fool can write code that a computer can understand. Good programmers write code that humans can understand


https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go


Therefore, my general advice about performance is: first make it right, then make it fast. - Brian Kernighan

When programming, follow the camp rule: Always leave the code cleaner than you found it
