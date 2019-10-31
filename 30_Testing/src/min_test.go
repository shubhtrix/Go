package min_test

import (
    "fmt"
    "testing"
)

func IntMin(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

// A test is created by writing a function with a name beginning with Test.
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
    // t.Error* will report test failures but continue executing the test. 
    // t.Fail* will report test failures and stop the test immediately.
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

// Writing tests can be repetitive, so it’s idiomatic to use a table-driven 
// style, where test inputs and expected outputs are listed in a table and 
// a single loop walks over them and performs the test logic.
func TestIntMinTableDriven(t *testing.T) {
    var tests = []struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    // t.Run enables running “subtests”, one for each table entry. 
    // These are shown separately when executing go test -v.
    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}

/*
func main() {

	a := IntMin(5, 7)
	fmt.Println("A :", a)
}*/
