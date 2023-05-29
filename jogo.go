package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    number := rand.Intn(100) + 1
    guess := 0
    fmt.Println("Adivinhe o número entre 1 e 100.")
    for guess != number {
        fmt.Scan(&guess)
        if guess < number {
            fmt.Println("Muito baixo!")
        } else if guess > number {
            fmt.Println("Muito alto!")
        } else {
            fmt.Println("Parabéns! Você acertou!")
        }
    }
}
