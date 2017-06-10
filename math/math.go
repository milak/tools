// Usefull math functions
package math
// Returns true if the number is Even
func Even(number int) bool {
    return number%2 == 0
}
// Returns true if the number is Odd
func Odd(number int) bool {
    return !Even(number)
}
