// Deteksi Race Condition Menggunakan Go Race Detector
// Go menyediakan fitur untuk deteksi race condition. Cara penggunaannya adalah dengan menambahkan flag -race pada saat eksekusi aplikasi.
// misal : go run -race 60_Mutex.go 

package main

import (
    "fmt"
    "runtime"
    "sync"
)
// Penerapan sync.Mutex

// Before sync.Mutex
// type counter struct {
//     val int
// }

// func (c *counter) Add(int) {
//     c.val++
// }

// func (c *counter) Value() (int) {
//     return c.val
// }

// // After sync.Mutex
// type counter struct {
//     sync.Mutex
//     val int
// }

// func (c *counter) Add(int) {
//     c.Lock()
//     c.val++
//     c.Unlock()
// }

// func (c *counter) Value() (int) {
//     return c.val
// }

// func main() {
//     runtime.GOMAXPROCS(2)

//     var wg sync.WaitGroup
//     var meter counter

//     for i := 0; i < 1000; i++ {
//         wg.Add(1)

//         go func() {
//             for j := 0; j < 1000; j++ {
//                 meter.Add(1)
//             }

//             wg.Done()
//         }()
//     }

//     wg.Wait()
//     fmt.Println(meter.Value())
// }


// Penggunaan tanpa di-embed/di-tempel
type counter struct {
    val int
}

func (c *counter) Add(int) {
    c.val++
}

func (c *counter) Value() (int) {
    return c.val
}

func main() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
    var mtx sync.Mutex
    var meter counter

    for i := 0; i < 1000; i++ {
        wg.Add(1)

        go func() {
            for j := 0; j < 1000; j++ {
                mtx.Lock()
                meter.Add(1)
                mtx.Unlock()
            }

            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(meter.Value())
}
