package problem5

import "testing"

var pair = Cons(3, 4)

func TestCar(t *testing.T) {
    if Car(pair) != 3 {
        t.Fail()
    }
}

func TestCdr(t *testing.T) {
    if Cdr(pair) != 4 {
        t.Fail()
    }
}
