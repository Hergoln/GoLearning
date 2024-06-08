package math

import "testing"

func TestAdd(testingContext *testing.T){
    got := Add(4, 6)
    want := 10

    if got != want {
			testingContext.Errorf("got %q, wanted %q", got, want)
    }
}