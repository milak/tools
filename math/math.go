package math

import (

)

func Even(number int) bool {
    return number%2 == 0
}
func Odd(number int) bool {
    return !Even(number)
}
