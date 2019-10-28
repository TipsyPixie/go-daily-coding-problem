package problem5

import "testing"

func TestCar(t *testing.T) {
    pair := Cons(3, 4)
    if Car(pair) != 3 {
        t.Fail()
    }
    if Cdr(pair) != 4 {
        t.Fail()
    }
}
