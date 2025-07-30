package main

// Task 1: Basic string function
func ToUpper(s string) string {
	// TODO: Convert string to uppercase: return strings.ToUpper(s)
	return ""
}

// Task 2: String processing functions
func CountVowels(s string) int {
	// TODO: Count vowels in string:
	// vowels := "aeiouAEIOU"
	// count := 0
	// for _, char := range s {
	//     if strings.ContainsRune(vowels, char) {
	//         count++
	//     }
	// }
	// return count
	return 0
}

func ReverseString(s string) string {
	// TODO: Reverse string:
	// runes := []rune(s)
	// for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
	//     runes[i], runes[j] = runes[j], runes[i]
	// }
	// return string(runes)
	return ""
}

// Task 3: Numeric functions
func IsEven(n int) bool {
	// TODO: Check if number is even: return n%2 == 0
	return false
}

func Abs(n int) int {
	// TODO: Return absolute value:
	// if n < 0 {
	//     return -n
	// }
	// return n
	return 0
}

func Sum(numbers []int) int {
	// TODO: Sum all numbers in slice:
	// total := 0
	// for _, num := range numbers {
	//     total += num
	// }
	// return total
	return 0
}

// Task 4: Functions with error handling
func ParseInt(s string) (int, error) {
	// TODO: Parse string to integer: return strconv.Atoi(s)
	return 0, nil
}

func Divide(a, b int) (float64, error) {
	// TODO: Divide with error handling:
	// if b == 0 {
	//     return 0, fmt.Errorf("division by zero")
	// }
	// return float64(a) / float64(b), nil
	return 0, nil
}

// Task 5: Edge case functions
func FindMax(numbers []int) (int, error) {
	// TODO: Find maximum with error handling:
	// if len(numbers) == 0 {
	//     return 0, fmt.Errorf("empty slice")
	// }
	// max := numbers[0]
	// for _, num := range numbers[1:] {
	//     if num > max {
	//         max = num
	//     }
	// }
	// return max, nil
	return 0, nil
}

func IsPalindrome(s string) bool {
	// TODO: Check if string is palindrome:
	// if len(s) <= 1 {
	//     return true
	// }
	// return s[0] == s[len(s)-1] && IsPalindrome(s[1:len(s)-1])
	return false
}

// Task 6: Complex functions
type Person struct {
	Name string
	Age  int
}

func FilterAdults(people []Person) []Person {
	// TODO: Filter adults (age >= 18):
	// var adults []Person
	// for _, person := range people {
	//     if person.Age >= 18 {
	//         adults = append(adults, person)
	//     }
	// }
	// return adults
	return nil
}

func GroupByAge(people []Person) map[int][]Person {
	// TODO: Group people by age:
	// groups := make(map[int][]Person)
	// for _, person := range people {
	//     groups[person.Age] = append(groups[person.Age], person)
	// }
	// return groups
	return nil
}
