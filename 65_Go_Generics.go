package main

import "fmt"

// func Sum(numbers []int) int {
//     var total int
//     for _, e := range numbers {
//         total += e
//     }
//     return total
// }

// Dengan generics
// func Sum[V int](numbers []V) V {
//     var total V
//     for _, e := range numbers {
//         total += e
//     }
//     return total
// }

// Comparable Data Type pada Fungsi Generic
func Sum[V int | float32 | float64](numbers []V) V {
    var total V
    for _, e := range numbers {
        total += e
    }
    return total
}

// Jika ingin kompatibel dengan semua tipe data maka gunakan comparable, penulisannya menjadi V comparable
func SumNumbers1(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

func SumNumbers2[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// Generic Type Constraint
type Number interface {
    int64 | float64
}

func SumNumbers3[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

// Struct Generic
type UserModel[T int | float64] struct {
    Name string
    Scores []T
}

func (m *UserModel[int]) SetScoresA(scores []int) {
    m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64) {
    m.Scores = scores
}


func main() {
    total1 := Sum([]int{1, 2, 3, 4, 5})
    fmt.Println("total:", total1)

	fmt.Println("\nComparable Data Type pada Fungsi Generic\n")
	total2 := Sum([]float32{2.5, 7.2})
	fmt.Println("total:", total2)
	
	total3 := Sum([]float64{1.23, 6.33, 12.6})
	fmt.Println("total:", total3)	

	fmt.Println("\nDengan keyword comparable\n")
	ints := map[string]int64{ "first": 34, "second": 12 }
    floats := map[string]float64{ "first": 35.98, "second": 26.99 }

    fmt.Printf("Generic Sums 2 with Constraint: %v and %v\n",
        SumNumbers2(ints),
        SumNumbers2(floats))
		
	fmt.Printf("Generic Sums 3 with Constraint: %v and %v\n",
        SumNumbers3(ints),
        SumNumbers3(floats))

	fmt.Println("\nStruct Generic\n")
    var m1 UserModel[int]
    m1.Name = "Noval"
    m1.Scores = []int{1, 2, 3}
    fmt.Println("scores:", m1.Scores)

    var m2 UserModel[float64] 
    m2.Name = "Noval"
    m2.SetScoresB([]float64{10, 11})
    fmt.Println("scores:", m2.Scores)	
}