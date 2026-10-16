package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.3" )

func eqCu6FZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ofC4l4AzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOVb8iayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tA0kvGcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSRu0QqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKa0cJI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAoely9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqrN3TJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g3fNqfQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x59UHr4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZtIDQmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3v9JdM0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1fg0zQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tB47izyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mrjBnxH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRKGYMCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6qr4V9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BD6xTevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOOMtJb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAwoX0PSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQdSD67LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxmL1BzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3MnjDnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8lnOepu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7PWdVlHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6A0jYG4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOAOjBflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwbW3MfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5SJ5ucrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fge5xDy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RP7cL036Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TwIt3QTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPeQn0EGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eneZSP36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmWP3BJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWsLOxh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GN4NbbxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0sWTmqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJKPie6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XssJ5nssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJP7V5RPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ru3XqWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrTPF6PuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkgMKfLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EA075meUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2sZRctzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEhj7a22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjPGXZJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5NmuR5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HccaGUb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3SB3kmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkY7huB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnmpUgszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5346S6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cc6t7oZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Izz0OZwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwA9MNrxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rljn0L5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xon11LzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0isut6hpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKCqoezfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NA2mvlBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Ng2A9hYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVnpwckTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwnzTynWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxHM7H4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0RmXxABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wWmF1hcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjLCYzoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBrGYJSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k54PfoiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqoD7f1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJNxOdtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duWdGVb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PQKwpGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNqG7UfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Blo5bq49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcSOOZ4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqoxFRrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnp3wmO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmy3kHmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCjPKAXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2YjWlYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujdsgiPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3h6n7SRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvKvhMvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rx5tTfEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pwm8PD8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21R6kDGsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBnXBJcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vodPar65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZ2EcJqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNlMpg44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJxLsSuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mrijg7qAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMBMmIFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SeHX9M5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8jtILQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMIFH61zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZzhNVzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jD2phfjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLRG4Re9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wa8zJ6TPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcgE16YrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIgLQQqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJp1vxhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiV3gOIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S927YU9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHudBOMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WKJA99aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxGuRj3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R48NY4RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yIbAcMOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixuyRMsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmvMsfnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fT4ZW97hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHtpq7INWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELVHxwYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WazaNZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShuxtWDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivE7Sx3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2gmhqUgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXvSxgEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kBEWvuOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRgtUWwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNHJnIesWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rstHpV4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZj63NlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMXEj1NHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7H2skWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2uITHL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nwg8aIcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDv3sX4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSBNtf5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UomXRE9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcsRFk1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooRYtPwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOoX8vtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EWsGWs4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func St7uaS3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqiDbYHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyVVwLI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeIggPTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XC4Wc0R2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SovnXZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jnfc8m5WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x91AzNNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MajJkZROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3RxcAfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJNazJbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lVl8KhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nY6jcE2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufVESxlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGxmwhFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8gAqWxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3UVncRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1SfozGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YNNzkR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFthU5wqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPZDgdWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRHuP5fyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7gK6RLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ws6V0s4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndeAtSU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0mrDIl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLOG6FBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPi7IXARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkmJ0EF0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1z0uS16HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DuHyBsLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PRIUCcHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erTqnJofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lm1W0lJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcwiWF8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcKrDwwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ffr5CVR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADMXGYfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCcmA6D3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uy81SO65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJk6K1L7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZGdJLQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdMrYliKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NI4qm6L1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeVx6Fi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func is9sBZOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXyPIWgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbnVWr2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taFXQBpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzLayheLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O24ye6LPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KILXsZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uWgANeNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rR1e6SXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hujFLiikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sOvuqTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBEVCa0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JQzRH08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJRQISmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNQaFTYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxMQww0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WtgJm0LcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJtDfwpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPx8beWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgYROgvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwqbPccSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zrfktyTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrSbnQKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SC7DnqWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Amme5MmjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvylSIo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSRjvZK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a3yge2yOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEQHRoUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func He00vuxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4oH6bo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Z4s7GF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGHo1HO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69AV3lIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwiu8jcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CNuWz2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFM8twreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btSyprVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WqO0ZMOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAW1eoSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9NYjou1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIhZCyEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6EU7aPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qz6fp0cxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcK4uZqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1ArbknUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOnfe7PoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pziTbQ1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLpCBOdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Clq3vvIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyNefQSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4I1MDN94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMHD5g0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcCSDHfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWMSBJgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8ipc3V5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZQdfCusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2sjGJvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBb4YCB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYArveH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtST21FwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pPAOdXRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxmTDW8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 45FsE92EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSv0xMDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IeBdi7WnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HH4Kn8TVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oso8w2YJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqpfkDzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNJHftBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvdZB8NbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Am21ubN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6QNUeFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQXXZivaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lffVVwtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTq6oMYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgEBSRJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yu4cmMGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCiVX7KiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vg4Ep3beWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMwF47g5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoNGjt8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MoItnm90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdefwRBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjaiU2EpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkQKheRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jdat6GmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yO8SUvPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o6jrKRe5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QvNFsReVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyKxQGXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRZGOWiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0Sc2GFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2zBvvFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2M17wN2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lir59aX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chWIjzLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yRrfOslTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMG7WM1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sWAotXNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJ9JmcURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxQ1k46iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjkA5FMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aB5BXtX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTAjPab3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VT4uxFw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mB4b2dEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpVbfVbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iau6VS9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTydUG39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcRAk5VOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0SoF9rvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6OO7uqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWJY9pwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfjmvOMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pawFFnDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCVVcp2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JCUtyrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0963uWcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRHrjYSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPDfgJTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COtEjXhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5fxxYGaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiNm72RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mi44TlTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEtudbXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byKS9RuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PdfN9KwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yabfsQqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3cdNYQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHKjUHnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLsjnZyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dE5g2bhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Fca5mHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIif5C9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlig0M8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFLg5HuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbOjfy2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMRpIG17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8miahaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVmO4uVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K34R9lDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJeP5M5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L60kHX5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M0COEI58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUOf3A2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tq69LMx8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bs4BMpLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwJQKPCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BenHYiNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSFTWmU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wnz8wEGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YnudoBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZ0tA5KdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5asxFtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGWERf6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8gzjP9GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKd8RJKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AXxwiD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sddpaxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQ3Muc6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yoqrqlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaNRTQewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q08neOmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3r2QKOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pe44V51iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U2MUx35sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96ILt02JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qw6FCamrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kctWXZyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86DKXw9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9D98KTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yd8snewUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xb7bVMEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KaSs2iHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcmKOdwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TEQzhjKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6HPMHReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7G2Psuy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fl3NY3K7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0bDJvzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C98hadIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGQDBbl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hSLodEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5mwfiAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8Xv33ZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJzAxLx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMlr1WbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EhsjBEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnZtcxvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwrQ6Wf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mf1O4gLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7XUJJ18OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nJx39nYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ArPYzZWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXUDJcv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvdRMlcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WOBN5KSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 490VgvxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgVAXw3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJ6wghIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqxaZVowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYxFMhl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhKHo2S5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8scrYd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hioOgAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKsueGuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rI2ox8CaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uXgFKDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5ZaGjMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZQnssarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uAQHKO5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkWkFsGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axVjV4WeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9JXW8xIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQ1VVt57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mbfJQw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yn8N6qX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AON9a2loWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJtdZLECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJ1x3jr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQJynOfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JjMyHg54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJ9J7qx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGqKdI4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyggLA1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YZqU5qBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ok5UGFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Gqh10dUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MaawooGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9Xp6JciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4IBlQuOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1r4Zl7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUp4wlS0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CIuvHUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeU9dTyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCKNlqqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5yFgAU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xh5RMXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inDw1jleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AH5WbuLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xVAeIW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdF3LDQ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkdY7ossWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ONR5ZCCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2kfJ8wbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9uMrL3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRByfKXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ooQONT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsXaKhafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xnuLgny0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRF49l0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0sJJs0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZEJcsgbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBHvYLZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiqdMCiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVWFNTsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Czz9pU8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wbf3xUFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAcyVDFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0pWjqLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bEbRnsRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZRsjVduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUMZukWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tliSo6QRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rzmx7buZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9fmLylCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCkGOz8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyzEHQQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDqUsHXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbFfJbZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Enfdl8lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2DXESIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNGKEcToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukIUVtyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ci87NUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DK75MwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNqhFaaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AhxcZLnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yE9Tmfm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PO4b8c7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JOmVJKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHWW1u5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ez8wwd6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrkLNeCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZiOL6mZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRNBSCjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7NHc2yaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jo1JEDifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0opgalBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kG9oxt3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ApEN5BCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLQoqd4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JabEomLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVIlQ0N7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkXIaFE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIMhnu6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjsecMT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4wEdQQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6Y9kqBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfZ7uEoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bxx0y2SjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPwTHXYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vvw2mJjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfc3buCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqLOXLrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZp6WjMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPbYJ6XNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoHHylCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fNxTvAtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1cSnwrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTliwPnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVzcbFVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6AchjpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4j2kyEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6DbXvnjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaRmsCaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3iNU0xtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9dzyRI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9yKLtLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lF5hY5ygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axKvAbm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlLdjAicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6fh5RtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ggj2GE8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxlZTPhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 817A5F5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WurhzaWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrqCKoFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmwxcPtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLiFcz0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6gBk3ikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBY81ONwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQqNyLxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuFwZJh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTeNpwWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCJCERHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znIWCjK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKcTiKKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Op9P4nv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTiKzNrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKtkR4BAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxatdFtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOX4sxZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQSySylZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLrijWwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9r8oEw1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxdoC2BgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTtuq81mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljGdAJQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cphg6SL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7blkR3uqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wq9JtIG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xTOfyliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wn8KkbOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rOODZysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFpaI6tlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7eVZyHjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2A6lUmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNIgUL0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9RhPWZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fds4jtYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhLfRmo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdU9A2STWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNsvREciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RaocEWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 011042qaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwS694RHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIzUgDZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVpW1bd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EKexFMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouGzzYenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqO7y213Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkPz3LhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCoGx1wfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fui2lKChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUcsPbGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9bozqiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwbZhaf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVndUuB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2oMEEP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MkimSnVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5KJNKdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OgDVSBFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFgEJukEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDcNGOrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJukWRthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mkKBNUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tm5B1cSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AjCFrTInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpsPytLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkJ1h53tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dnolfNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKK2n8iMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xb8KdlssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1J2XAc41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8ZTnIHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgbectQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GaS08cm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzKNubgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOrf6GWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0Krd0I1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EimX39vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NH6C9MLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Iq5e7oGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSValDjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d462EMUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AIEIJNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htufNU6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTi48UkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVQfvVlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQxdCZwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCrh5zDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhzoNy5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHGVdsrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1s7JDbgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7HddRkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLKyu1IAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvK5GveNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awIcjWXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PFiHmSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2k0GWHUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fK3yNWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tC0nOKrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjhgovlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ewn19N1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spX9D6spWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b92wZfsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9uvFff1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69fORRTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E39zMz7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5FXXueUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsB5Iad2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAAdP2AsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wy4p0E0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDpz2F89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Qj9JtqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6lohfyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpD8EGd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdZxEP80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsDnVpRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbIqxNFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WHJoCHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftp3cAJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9xxQDiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSQkSZFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0E3k32qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxSiaRFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYZ5VdauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHQ3Rt1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uOVjO6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3wYqvL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRb7ZHCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0q6N2hpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnjzm5LcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkvaaiL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZ8C7JFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVaNgLAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGYxxFqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1mcw1181Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ncVuYtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1XIDVtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyG1NG8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvGENmKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftjYKwroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 504DaWOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lS8IlwkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5lIfKUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zickoINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdwSIdBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWYt4o3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dn8xeEX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaRmJ1n6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jL7rjPGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPasiuEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwwpVdztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3E6iTHgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hR5RiGM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOCaEp90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtVm8WCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LpybRQVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3YOgSmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjEFy3YaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stf89DP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4lMqPpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5TlqBU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipIX3XfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwaAUrVtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zQaupjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UgKbp9fSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8oryD2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYidi44MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kO9yxP6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsX8BzE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0qKLYFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSBHCIRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PevNAV21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlQyFqQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwDflAtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqAIVqwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgkknLQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhS8fDh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eW6FVQCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQ0jsGcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8I5xrPOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQ8iACSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1jNWTEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8AEu6vYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPymFOASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdmO0kbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rp3wx4XYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUGBX8cXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWHMSiNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9IC3N67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAmcFeIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8VhMLMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iydpEm6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBfpR3YZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8y0LYI12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 939E0aXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dl7VPjQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0MHiFlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxVfH9xMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Y81N579Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q60VeMFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nHAniJFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 473oeltGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vzdymCf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuP2IUAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IxXLYwnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O85EOIreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DiIhVsaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDkFOWVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyrvfbcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URiVhWHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRbDG1fHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxruHAEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3m8wy5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qurUdzjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdzNF2lBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Gpkp826Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjLD9zFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wppgeJDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fi6LM69YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdWzjFlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mPOBTENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzIy3UxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6dr03B4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbZVxtWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYj6Arh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FmWvOPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfxA7FjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZTXhkDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnrwZuTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 25MFtL8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Z4JQcrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bh0QOhb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ff2eaVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TD2CJNcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIyIyIgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xe5X9EsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func si0OeqfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RWaWqrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rPJp6tHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQGNNl7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxuhyLvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqAMrQTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nP8ICMHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7K96P7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSonVKluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxDrRUlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0F50ZLE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cGBWUPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTlnDqWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfnSEpOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C80RDKhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3AyfFeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlrwrvMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJbkHBkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHJ6MNUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hAoHTRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWdORq8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xt2Qilx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsWVf3fgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVcBLb0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtiEtVYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NZOK6RGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmsOlDjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yuMZCkWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kX7VrarvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uqrIsxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tNvCc4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzjBjlk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1hZqo10sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zllaQ2Z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Spm0FO8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iWT5jutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fu8aBUVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMY1NtQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWdwhiEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xB1Tsz85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrKcXe1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qtUTMxJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GoC0b6WoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AqG7fcNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sz5ECyXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deCn6ZjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrYGsiDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsy4dzxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWqrMxjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fu7H4ZsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tzg0urVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGQTerssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bb8ZxpPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfyDvNNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKO6frhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5nsKpV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bypyA72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISUGWxLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvuU1UJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RvSseL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPHTFRCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6IM6D0ZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0vvJOprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4gFafQHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oldbvZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBIwPRRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DM0jUWB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TPXyqDWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4XsQFbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rc8dhKKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBWxhyCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZE0E0fqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N06YVJGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNkW9aMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3ccL6h4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVM6MhU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0laDMSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hz6nUhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Dq1bCiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Lv95JMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L53JMinpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 642xE2JZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkT8zWJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vo0ivNmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func noa6plmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEJdMfZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hW070nfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHgu1fX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZ2bJap1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUws8VL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMTNPOtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUqHwoCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3kJaJKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjg6oT73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QPcILAZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UvGYvjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9rCpec8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXp1Lfm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDFKv4TkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mutQ2qpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3upaq6PGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijQjyqYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lECu0pl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pQafgIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjR0VLuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4P3iaNjFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbHIV5IyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dP967MukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpHWSfZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7Z5M7hpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EBHkU1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64doCgy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bHsdpIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VY2DO1HqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbzTpSPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s3FLEG64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CcZ4bosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diWawXrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFWlurypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ME1PRuMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pug4azkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YPX5FBU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4B6N9TCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZ1ddVuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjme5DVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRSoTwljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1Ywc4Z0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMoW8UjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUrB8bSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WBdtD3hEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXADdXAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsMGdx39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWGcARfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhjZG8NxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0mk6nFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUN1J2d9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Skv93tnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVYY8RtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZi4qjiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5mnqraxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJ3Ajy67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhmHMrR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dI8dXkDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SCN1gETCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7h3lY8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJHk0WvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUFwtpCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J56XEa4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgFKkVSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enie7HU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0a7pdEkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8jmuwIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ziABEy7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSKrZJviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zaREwYA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ycKM9S0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpyH2rxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NM8qB5olWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OcRdPWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2IW6HN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9kslflpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNpd95m9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLFd4VLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0trRmnXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgUCmrBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Y2W4qBsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmX0l9GJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqKDbVXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGyIOqKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5amVKQtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aIOAW9dMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSwhZL0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcIxuLVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysjxklnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7FN1I1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAvIXcYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJl89oXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQPO8kBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MD7cgVkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXIyZhZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSeiCLNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ALQFnhXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QuWpZuHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ah5Cs0s0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBfVZTBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jA3N5OGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkYzuMmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gd7JvNMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxmQvBeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNVNdNHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jotlnUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFPSZF8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHYmO2v2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFu874qRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPTmxV5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3bTSz7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8JnRFgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDS9UAv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6ayvgNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKwOptZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybkvAOGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWEPF5ghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTwNdJuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7RMvFMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJAx4k5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8AqtDuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func StKKZ9f2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKKUKqddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FMhiNAJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qSA95v0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4V6LqpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uuyt9N6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obT5J6wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qAvOAaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RB7obW8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiptGgCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhULKVBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nay9idCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owxLfS4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhFZvWHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCs1tVXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9njqkma3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XNhjLiLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxJJYzKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsCUUx9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Te5eJMyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3XuQEAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwffkYDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YK3WqBB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iB8mx6tEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wvQzfT8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgQVBgXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aRPpBcWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func METNpuVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxwgi9atWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zFj6XxUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYxNR4eSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PERGfdhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhPB1YCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GX6d7YfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JctjGfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhCnHMwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OB9rZIdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULcWhzDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qy3h6WSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbuFEO7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIrzdimoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCXutgAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1vDROEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVO8GgnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66yX6dUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qR9iVxOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXwrg58NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0h89lbzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upx8AcL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZA3BDaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mOcy8C2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fzk57hBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func StKLXK45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baLwXwgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9WEicgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMLwPADgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXhA6KMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQE3Pk3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S62vZgKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WsrK9nRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJ9AOsgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89JjthvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wAqSoDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxo9DdoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlzsoikPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KaBQlvj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A82G9NBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ydc63J7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHVUjbr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekQIPwNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwup3NiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZjLvjR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2ExoX7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BGZbltjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXe2bV3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pH9kQdxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jln2IJ0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHxvBEF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHHQNjP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsiDQaFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJ9n1SUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIMl8rpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOd476X5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYTCKam4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dab2qeANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWDCd0POWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3aQewOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OQoI9ZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sz2HQD4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVO8s1NVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9vRZ2vOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gO6MMhtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxsqvF1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FCLn9dBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fduhLgNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VpHHGPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lv0R59UGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtLWNt3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BI5QAbY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCRaqk8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzLFle7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmpDMl0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiCSBRdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeqehnL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6j6rZhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2DDjJIgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BQiGd65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSY3TY7KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trzVhF7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdyJVls3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSTinrL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c61NP3SsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sY1CWY9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBaWtQdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmI5jghcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouIQaYHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7fol78nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0D2jrfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Okimi961Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lyNqfw75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdkfZD7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVwxJPoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfjdiDk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzWSbAIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nriWbKO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjPi82yvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEQs80kmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVAumSiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8d7hDztyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEr14ouzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkBJ5G2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eR1gH0BEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sX9WWruIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func la8gyC3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqnn1xZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTgZgV0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sU8Vi2mpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AsiQZzyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zU6xL5umWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HZ7R3IAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4voQTZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JvsETDPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CaVh2ZkAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5feccpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoSBRHoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nqUo9EOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYvMc2HHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiWP2qMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBInWj8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ByzDnb83Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cWKz1s51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdGOsrtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8pWFtpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQJkBuLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flDRZXCtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5bUkCB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func de39pc1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXRXhXG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JA5zYr7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJuic7SDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4O6qGivtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SaO2z5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZthH1kf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tftWz1y6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2Z74GdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSM9XcoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCmGPqXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLwviyJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMn2WjJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgqAfssQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHvAyUxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzujObadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpc77dD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBGgIA6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LNJP3ggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OCYwcetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPFP9K5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykuVfGJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqxwwkEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2ipwWRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIKEmSNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjO4KaKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9inoPzAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvr3ARAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1z68G9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwW9mQG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8jLgD5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjOqzY5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bs7nur4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UL09IXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q63ph0t8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wnmIPMruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYcc9ym3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyPIrR7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9wdOE1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdZ3TQWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXUpUWngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XgFDHPlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WX77tMw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tAE02Yx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xGy4R0LQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ep869kfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aoS7z4DDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PP9Hx3tXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxLJoYVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lulisa5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoXUQMAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1uhUenBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X915fFiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSv60y0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzXkuuzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opv5MRRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u2nubLVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6FLfZbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8h19QmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MF5UsyefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2zjk3gOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ej6ZtuZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLnIzu2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bmbmuFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAUMN3YlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0kAemdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BQeHvbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxhxU3ULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzconoaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxTiecVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

