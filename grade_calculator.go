package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
var m = make(map[string]float64)
var letter_grade = make(map[string]string)


func main(){

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("hey there what's your name : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %s Welcome to the Grade Calculator! please enter the number of courses you want to calculate the grades for: ",name)	
	var number_of_courses int
	fmt.Scanln(&number_of_courses)

	for i := 0; i < number_of_courses; i++{
		
		fmt.Printf(" please enter the course and the grade you got in the course: ")
		var course string
		var grade float64

		_, err := fmt.Scanln(&course, &grade)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		if grade >= 85{
			letter_grade[course] = "A"
			m[course] = 4.0
		} else if grade >= 75{
			letter_grade[course] = "B+"
			m[course] = 3.5
		} else if grade >= 68{
			letter_grade[course] = "B"
			m[course] = 3.0
		} else if grade >= 55{
			letter_grade[course] = "D"
			m[course] = 1.0
		} else {
			letter_grade[course] = "F"
			m[course] = 0.0
		}
		// m[course] = grade
	
	}
	average_grade := (grade_calculator() * 5) / float64(number_of_courses*5)
	for key, value := range letter_grade{
		fmt.Println(key, " : ", value)
	}
	fmt.Println("your grade for this semister is: ", average_grade)

}	

func  grade_calculator() float64{
	var total float64
	for _, value := range m{
		total += value
	}
	if len(m) == 0{
		fmt.Println("You have not entered any grades")
		return 0
	}
	return total
}

