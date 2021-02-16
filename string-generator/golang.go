package main

import (
    "fmt"
    "time"
    "math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")

func generateString(length int) string {
    psd := make([]rune, length)
    for i := range psd {
        psd[i] = letters[rand.Intn(len(letters))]
    }
    return string(psd)
}

func generateInt(length int) string {
    psd := make([]rune, length)
    for i := range psd {
        psd[i] = numbers[rand.Intn(len(numbers))]
    }
    return string(psd)
} 

func main() {
    rand.Seed(time.Now().UnixNano())

    fmt.Println(generateString(4) + generateInt(2))
}
