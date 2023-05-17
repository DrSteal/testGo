package main

import (
	"fmt"
)

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a, b, c int
// 	fmt.Println("Enter a: ")
// 	fmt.Scan(&a)

// 	fmt.Println("Enter b: ")
// 	fmt.Scan(&b)

// 	fmt.Println("Enter c: ")
// 	fmt.Scan(&c)

// 	D := (b*b) - 4*(a*c)

// 	if D > 0 {
// 		var x1 float64
// 		var x2 float64

// 		x1 = (-b + math.Sqrt(D)) / (2*a)
// 		x2 = (-b - math.Sqrt(D)) / (2*a)

// 		fmt.Println("Your the equation has 2 roots")

// 	} else if D = 0 {
// 		x = -b / (2*a)
// 	}

// 	}

// package main

// import "fmt"

// type User struct {
// 	name     string
// 	age      int64
// 	password string
// 	score    []int
// }

// func (u User) getHighScore() int {
// 	high := 0

// 	for _, sc := range u.score {
// 		if high < sc {
// 			high = sc
// 		}
// 	}

// 	return high

// }

// func main() {
// 	user := User{"Henk", 20, "pass", []int{33, 91, 942, 1, 5456}}
// 	fmt.Println(user.getHighScore())

// }

// package main

// import (
// 	"fmt"
// 	time2 "time"
// )

// type myStruct struct {
// 	count  int
// 	result string
// }

// func main() {

// 	str := getRunesString("12345")
// 	time := time2.Now().Format("2006-01-02")
// 	fmt.Println(">", str.count, " ", str.result, " ", time)
// }

//	func getRunesString(s string) myStruct {
//		// получить слайс из строки
//		counter := 0
//		strRunes := []rune(s)
//		for i, r := range strRunes {
//			fmt.Println(i, " - ", string(r), " ", i%2)
//			if i%2 == 0 && i+1 < len(strRunes) {
//				//надо поменять значения
//				tmp := strRunes[i]
//				strRunes[i] = strRunes[i+1]
//				strRunes[i+1] = tmp
//				counter++
//			}
//		}
//		return myStruct{
//			count:  counter,
//			result: string(strRunes),
//		}
//	}

// func statusByName(name string) string {
// 	// функция проверяет, что строка name начинается с подстроки "Mr."
// 	if strings.HasPrefix(name, "Mr.") {
// 		return "married man"
// 	} else if strings.HasPrefix(name, "Mrs.") {
// 		return "married woman"
// 	} else {
// 		return "single person"
// 	}
// }

// func main() {
// 	n := "Mr. Doe"
// 	fmt.Println(n + " is a " + statusByName(n)) // Mr. Doe is a married man

// 	n = "Mrs. Berry"
// 	fmt.Println(n + " is a " + statusByName(n)) // Mrs. Berry is a married woman

// 	n = "Karl"
// 	fmt.Println(n + " is a " + statusByName(n)) // Karl is a single person
// }

// type UserCreateRequest struct {
// 	FirstName string // не может быть пустым; не может содержать пробелы
// 	Age       int    // не может быть 0 или отрицательным; не может быть больше 150
// }

// func main() {
// 	u := UserCreateRequest{"fjs", 3}
// 	fmt.Println(u)
// }

// func Validate(req UserCreateRequest) string {
// 	if strings.Contains(req.FirstName, " ") || req.Age <= 0 {
// 		fmt.Println("invalid request")
// 	}
// 	return fmt.Sprintf("%s%d", req.FirstName, req.Age)
// }

// func Validate(req UserCreateRequest) string {
// 	if strings.Contains(req.FirstName, " ") || req.Age <= 0 {
// 		fmt.Println("invalid request")
// 	}
// 	return fmt.Println("fsdf")
// }

// func main() {
// 	var arr [100]int
// 	for _, x := range arr {
// 		rand.Seed(time.Now().UnixNano())
// 		x = rand.Intn(100 + 1)
// 		fmt.Println(x)

// 	}

// }

const (
	year     = 365
	leapYear = int32(366)
)

func main() {
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(int32(hours) * leapYear)
}

// func Solution(word string) string {
// 				runes := []rune(word)
// 				for i, j :=
// }
