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
	count       int
}

func Welcome() {

	fmt.Println("Welcome to Amazing Numbers!")
	fmt.Println()
	fmt.Println("Supported requests:")
	fmt.Println("- enter a natural number to know its properties;")
	fmt.Println("- enter two natural numbers to obtain the properties of the list:")
	fmt.Println("	* the first parameter represents a starting number;")
	fmt.Println("	* the second parameter shows how many consecutive numbers are to be printed;")
	fmt.Println("- two natural numbers and a property to search for;")
	fmt.Println("- separate the parameters with one space;")
	fmt.Println("- enter 0 to exit.")
	fmt.Println()

}

func Request() (string, string, string) {
	var number string
	var size string
	var property string

	fmt.Println()
	fmt.Print("Enter a request: > ")
	fmt.Scanf("%s %s %s", &number, &size, &property)

	return number, size, property
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
	fmt.Println()

}

func MultipleFinal(number int, properties Properties) {
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

	fmt.Println(final)
	fmt.Println()

}

func Check(number, size int, property string) (*int, *Properties, int) {
	var properties Properties
	var count = 0

	if !IsNatural(number) {
		fmt.Println("The first parameter should be a natural number or zero.")
		return &number, &Properties{}, 0
	}

	if !IsNatural(size) {
		fmt.Println("The second parameter should be a natural number.")
		return &number, &Properties{}, 0
	}

	even := EvenOrOdd(number)

	array, isEnds := IsEndsWithSeven(number)

	removedDigit := array[len(array)-1] * 2
	remainingNumber := (number - array[len(array)-1]) / 10

	subtraction := remainingNumber - removedDigit

	isDivisible := IsDivisibleBySeven(subtraction)

	palindromic := IsPalindromic(number)
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
		properties.odd = false

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

	return &number, &properties, count
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

func Arrays(size, number int, property string) []int {
	var numArray []int

	//propertyArray := [7]string{"even", "odd", "buzz", "duck", "palindromic", "gapful", "spy"}
	i := 0
	if size > 0 {
		if property != "" {
			// if  {
			// 	fmt.Printf("The property [%s] is wrong.\n", property)
			// 	fmt.Println("Available properties: [BUZZ, DUCK, PALINDROMIC, GAPFUL, SPY, EVEN, ODD]")
			// 	return numArray
			// }
			for i < size {
				_, properties, _ := Check(number, size, property)
				if properties.even && property == "even" {
					numArray = append(numArray, number)
					number++
					i++

				} else if properties.odd && property == "odd" {
					numArray = append(numArray, number)
					number++
					i++
				} else if properties.buzz && property == "buzz" {
					numArray = append(numArray, number)
					number++
					i++
				} else if properties.duck && property == "duck" {
					numArray = append(numArray, number)
					number++
					i++
				} else if properties.palindromic && property == "palindromic" {
					numArray = append(numArray, number)
					number++
					i++
				} else if properties.gapful && property == "gapful" {
					numArray = append(numArray, number)
					number++
					i++
				} else if properties.spy && property == "spy" {
					numArray = append(numArray, number)
					number++
					i++
				} else if !strings.Contains(property, "even") || !strings.Contains(property, "odd") || !strings.Contains(property, "buzz") || !strings.Contains(property, "duck") || !strings.Contains(property, "palindromic") || !strings.Contains(property, "gapful") || !strings.Contains(property, "spy") {
					fmt.Printf("The property [%s] is wrong.\n", property)
					fmt.Println("Available properties: [BUZZ, DUCK, PALINDROMIC, GAPFUL, SPY, EVEN, ODD]")
					return numArray
				} else {
					number = number + 1
					size++
					i++
				}

			}
		} else {
			for i := 0; i < size; i++ {
				numArray = append(numArray, number+i)
			}
		}

	} else {

		oneNum, properties, _ := Check(number, size, property)

		Final(*oneNum, *properties)
	}
	return numArray
}

func ReadWriteRequests() {

}

func main() {
	for {
		Welcome()
		strNumber, strSize, property := Request()
		if strNumber == "" {
			Welcome()
		}
		number, _ := strconv.Atoi(strNumber)
		size, _ := strconv.Atoi(strSize)
		if number == 0 {
			fmt.Println("Goodbye!")
			return
		}

		numArray := Arrays(size, number, property)
		for i := 0; i < len(numArray); i++ {

			num, properties, _ := Check(numArray[i], size, property)
			if num != nil {
				MultipleFinal(*num, *properties)
			}

		}

	}

}
