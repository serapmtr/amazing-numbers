package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Properties struct {
	even        bool
	odd         bool
	buzz        bool
	duck        bool
	palindromic bool
	gapful      bool
	spy         bool
	square      bool
	sunny       bool
	property    []string
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Welcome() {

	fmt.Println("Welcome to Amazing Numbers!")
	fmt.Println()
	fmt.Println("Supported requests:")
	fmt.Println("- enter a natural number to know its properties;")
	fmt.Println("- enter two natural numbers to obtain the properties of the list:")
	fmt.Println("	* the first parameter represents a starting number;")
	fmt.Println("	* the second parameter shows how many consecutive numbers are to be printed;")
	fmt.Println("- two natural numbers and a two properties to search for;")
	fmt.Println("- separate the parameters with one space;")
	fmt.Println("- enter 0 to exit.")
	fmt.Println()

}

func Request() (string, string, string, string) {
	var number string
	var size string
	var property string
	var secondProperty string

	fmt.Println()
	fmt.Print("Enter a request: > ")
	fmt.Scanf("%s %s %s %s", &number, &size, &property, &secondProperty)

	return number, size, property, secondProperty
}

func Final(number int, properties Properties) {
	fmt.Println()
	fmt.Println("Properties of ", number)
	fmt.Println("even: ", properties.even)
	fmt.Println("odd: ", properties.odd)
	fmt.Println("buzz: ", properties.buzz)
	fmt.Println("duck: ", properties.duck)
	fmt.Println("palindromic: ", properties.palindromic)
	fmt.Println("gapful: ", properties.gapful)
	fmt.Println("spy: ", properties.spy)
	fmt.Println("square: ", properties.square)
	fmt.Println("sunny: ", properties.sunny)
	fmt.Println()

}

// func FindFinalString(numbers []int, properties Properties) []string {
// 	final := fmt.Sprintf("%d is", numbers)

// 	if properties.buzz {
// 		final += " buzz,"
// 	}
// 	if properties.duck {
// 		final += " duck,"
// 	}
// 	if properties.even {
// 		final += " even,"
// 	}
// 	if properties.odd {
// 		final += " odd,"
// 	}
// 	if properties.gapful {
// 		final += " gapful,"
// 	}
// 	if properties.palindromic {
// 		final += " palindromic,"
// 	}
// 	if properties.spy {
// 		final += " spy,"
// 	}
// 	if properties.square {
// 		final += " square,"
// 	}
// 	if properties.sunny {
// 		final += " sunny,"
// 	}
// }

func MultipleFinal(number int, properties Properties) []string {
	fmt.Println()
	final := fmt.Sprintf("%d is", number)
	if properties.buzz {
		final += " buzz,"
	}
	if properties.duck {
		final += " duck,"
	}
	if properties.even {
		final += " even,"
	}
	if properties.odd {
		final += " odd,"
	}
	if properties.gapful {
		final += " gapful,"
	}
	if properties.palindromic {
		final += " palindromic,"
	}
	if properties.spy {
		final += " spy,"
	}
	if properties.square {
		final += " square,"
	}
	if properties.sunny {
		final += " sunny,"
	}

	fmt.Println(final)

	fmt.Println()

	finalArray := strings.Split(final, " ")
	return finalArray
}

func Check(number, size int, property, secondProperty string) (*int, *Properties, int, []string) {
	var properties Properties
	var count = 0

	if !IsNatural(number) {
		fmt.Println("The first parameter should be a natural number or zero.")
		return &number, &Properties{}, 0, properties.property
	}

	if !IsNatural(size) {
		fmt.Println("The second parameter should be a natural number.")
		return &number, &Properties{}, 0, properties.property
	}

	even := EvenOrOdd(number)

	array, isEnds := IsEndsWithSeven(number)

	removedDigit := array[len(array)-1] * 2
	remainingNumber := (number - array[len(array)-1]) / 10

	subtraction := remainingNumber - removedDigit

	isDivisible := IsDivisibleBySeven(subtraction)

	palindromic := IsPalindromic(number)

	properties.property = append(properties.property, secondProperty)
	if palindromic {
		properties.palindromic = true

	}
	if isEnds || isDivisible {
		properties.buzz = true
	} else {
		properties.buzz = false
	}

	if even {
		properties.even = true

	} else {
		properties.even = false
		properties.odd = true
	}

	if IsDuck(number) {
		properties.duck = true

	}

	if IsGapful(number) {
		properties.gapful = true

	}

	if IsSpy(number) {
		properties.spy = true
	}

	if IsSquare(number) {
		properties.square = true
	}

	if IsSunny(number) {
		properties.sunny = true

	}
	return &number, &properties, count, properties.property
}

func IsNatural(number int) bool {
	if number >= 0 {
		return true
	}

	return false
}

func EvenOrOdd(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

func IsDivisibleBySeven(number int) bool {
	if number%7 == 0 {
		return true
	}
	return false
}

func ParseNumber(number int) []int {
	slc := []int{}

	for number > 0 {
		slc = append(slc, number%10)
		number = number / 10
	}

	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}

	return slc
}

func IsEndsWithSeven(number int) ([]int, bool) {
	slc := ParseNumber(number)
	if slc[len(slc)-1] == 7 {
		return slc, true
	}
	return slc, false
}

func IsDuck(number int) bool {
	slc := ParseNumber(number)

	for i := 1; i < len(slc); i++ {
		if slc[i] == 0 {
			return true
		}
	}
	return false
}

func IsPalindromic(number int) bool {
	slc := ParseNumber(number)

	for i := 0; i < (len(slc))/2; {
		for j := len(slc) - 1; j >= i; {
			if slc[i] != slc[j] {
				return false
			}
			j--
			i++
		}
	}
	return true
}

func IsGapful(number int) bool {
	slc := ParseNumber(number)

	if len(slc) > 2 {
		concatenated := fmt.Sprintf("%d%d", slc[0], slc[len(slc)-1])
		concatNum, _ := strconv.Atoi(concatenated)

		if number%concatNum == 0 {
			return true
		}
		return false
	}
	return false
}

func IsSpy(number int) bool {
	slc := ParseNumber(number)
	var sum int
	var prod = 1

	for i := 0; i < len(slc); i++ {
		sum += slc[i]
		prod *= slc[i]
	}

	if sum == prod {
		return true
	}
	return false
}

func IsSquare(number int) bool {
	for i := 0; i <= number; i++ {
		if i*i == number {
			return true
		}
	}

	return false
}

func IsSunny(number int) bool {
	for i := 0; i < number; i++ {
		if IsSquare(number + 1) {
			return true
		}
	}

	return false
}

func Arrays(size, number int, property, secondProperty string) []int {
	var numArray []int

	propertyArray := []string{"spy", "even", "odd", "buzz", "palindromic", "gapful", "duck", "square", "sunny"}
	i := 0
	if size > 0 {
		if property != "" {
			for i < size {
				_, properties, _, second := Check(number, size, property, secondProperty)
				if !Contains(propertyArray, property) {
					fmt.Printf("The property [%s] is wrong.\n", property)
					fmt.Println("Available properties: [BUZZ, DUCK, PALINDROMIC, GAPFUL, SPY, EVEN, ODD, SQUARE, SUNNY]")
					fmt.Println()
					return numArray
				}

				if !Contains(propertyArray, secondProperty) {
					fmt.Printf("The property [%s] is wrong.\n", secondProperty)
					fmt.Println("Available properties: [BUZZ, DUCK, PALINDROMIC, GAPFUL, SPY, EVEN, ODD, SQUARE, SUNNY]")
					fmt.Println()
					return numArray
				}
				if (properties.even && property == "even" && Contains(second, "odd")) || (properties.odd && property == "odd" && Contains(second, "even")) {
					fmt.Println("The request contains mutually exclusive properties [EVEN, ODD]")
					fmt.Println("There are no numbers with these properties")
					return numArray
				}
				if (properties.duck && property == "duck" && Contains(second, "spy")) || (properties.spy && property == "spy" && Contains(second, "duck")) {
					fmt.Println("The request contains mutually exclusive properties [SPY, DUCK]")
					fmt.Println("There are no numbers with these properties")
					return numArray
				}
				if (properties.sunny && property == "sunny" && Contains(second, "square")) || (properties.square && property == "square" && Contains(second, "sunny")) {
					fmt.Println("The request contains mutually exclusive properties [SUNNY, SQUARE]")
					fmt.Println("There are no numbers with these properties")
					return numArray
				}
				if properties.even && property == "even" {
					numArray = append(numArray, number)
					number++
					i++

				}
				if properties.odd && property == "odd" {
					numArray = append(numArray, number)
					number++
					i++
				}
				if properties.buzz && property == "buzz" {
					numArray = append(numArray, number)
					number++
					i++
				}
				if properties.duck && property == "duck" {
					numArray = append(numArray, number)
					number++
					i++
				}
				if properties.palindromic && property == "palindromic" {
					numArray = append(numArray, number)
					number++
					i++
				}
				if properties.gapful && property == "gapful" {
					numArray = append(numArray, number)
					number++
					i++
				}
				if properties.spy && property == "spy" {
					numArray = append(numArray, number)
					number++
					i++
				}
				if properties.square && property == "square" {
					numArray = append(numArray, number)
					number++
					i++
				}
				if properties.sunny && property == "sunny" {
					numArray = append(numArray, number)
					number++
					i++
				}
				number = number + 1
				size++
				i++

			}
		} else {
			for i := 0; i < size; i++ {
				numArray = append(numArray, number+i)
			}
		}

	} else {

		oneNum, properties, _, _ := Check(number, size, property, secondProperty)

		Final(*oneNum, *properties)
	}
	return numArray
}

func main() {
	Welcome()
	for {

		strNumber, strSize, property, secondProperty := Request()
		if strNumber == "" {
			Welcome()
		}
		number, _ := strconv.Atoi(strNumber)
		size, _ := strconv.Atoi(strSize)
		if number == 0 {
			fmt.Println("Goodbye!")
			return
		}

		numArray := Arrays(size, number, property, secondProperty)
		for i := 0; i < len(numArray); i++ {
			num, properties, _, _ := Check(numArray[i], size, property, secondProperty)
			if num != nil {
				MultipleFinal(*num, *properties)
			}

		}

	}

}
