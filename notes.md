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

## Chapter 2

Refactoring (noun): a change made to the internal structure of software to make it easier to understand and cheaper to modify without changing its observable behavior.

Example:
- Extract Function
- Inline Function
- Move Function
- Change function signature/declaration
- Extract Variable
- Replace Temp with Query
- Replace Conditional with Polymorphism

Refactoring (verb): to restructure software by applying a series of refactorings without changing its observable behavior.

Kent Beck Two Hats Theory: When you are adding a feature, you are a programmer. When you are refactoring, you are a designer.

Refactoring helps me to develop code faster

Design Stamina Hypothesis [mf-dsh]: When we apply our stamina to design, we can go faster and further

The Rule of Three: The first time you do something, you just do it. The second time you do something similar, you wince at the duplication, but you do the duplicate thing anyway. The third time you do something similar, you refactor. - Martin Fowler


