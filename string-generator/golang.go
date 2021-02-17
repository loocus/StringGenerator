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
    rand.Seed(time.Now().UnixNano()) //idk why but if i do that, the password generation is more faster

    fmt.Println(generateString(4) + generateInt(2))
}
//Today : 17/02/2021, I will greatly optimize the code
