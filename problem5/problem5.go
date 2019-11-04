package problem5

type Func func(interface{}, interface{}) interface{}

type Pair func(Func) interface{}

func Cons(a interface{}, b interface{}) Pair {
	return func(f Func) interface{} { return f(a, b) }
}

func Car(pair Pair) interface{} {
	return pair(func(a interface{}, b interface{}) interface{} { return a })
}

func Cdr(pair Pair) interface{} {
	return pair(func(a interface{}, b interface{}) interface{} { return b })
}
