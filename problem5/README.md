## Daily Coding Problem: Problem #5 [Medium]

Good morning! Here's your coding interview problem for today.

This problem was asked by Jane Street.

`cons(a, b)` constructs a pair, and `car(pair)` and `cdr(pair)` returns the first and last element of that pair. For example, `car(cons(3, 4))` returns 3, and `cdr(cons(3, 4))` returns 4.

Given this implementation of cons:

```python
def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair
```

Implement car and cdr.

P.S. Go implementation of cons:

```go
type Func func(interface{}, interface{}) interface{}

type Pair func(Func) interface{}

func Cons(a interface{}, b interface{}) Pair {
    return func(f Func) interface{} { return f(a, b) }
}
```
