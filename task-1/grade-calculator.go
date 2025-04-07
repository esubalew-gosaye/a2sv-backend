package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func readFromInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	name, _ := reader.ReadString('\n')
	return strings.TrimSpace(name)
}

func validateGrade(grade float64) bool{
	if grade <= 0 || grade > 100 {
		fmt.Println("Grade must be between 1 and 100")
		return false
	}
	return true
}

func askName() string {
	return readFromInput("Enter your name: ")
}

func askGrade() map[string]float64 {
	courses := map[string]float64{}

	num, er := strconv.Atoi(readFromInput("How many course do you take: "))
	if er != nil {
		fmt.Println("Invalid input, please enter integer")
		return courses
	}
	for i := 0; i < num; i += 1 {
		courseName := readFromInput(fmt.Sprintf("Enter course %v name: ", i + 1))
		courseGrade, er := strconv.ParseFloat(readFromInput(fmt.Sprintf("Enter course %v grade: ", i + 1)), 64)
		if er != nil {
			fmt.Println("Invalid input, please enter integer")
			return courses
		}
		if validateGrade(courseGrade){
			courses[courseName] = courseGrade
		}
	}
	return courses
}

func calculateAvg(grades map[string] float64) float64 {
	var sum = 0.0;
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}

func displayTranscript(name string, grades map[string] float64, avg float64) {
	fmt.Println()
	fmt.Println("-------------------------------------")
	fmt.Printf("Student Name: %v \n", name)
	fmt.Println("-------------Course List-------------")
	for subject, grade := range grades {
		fmt.Printf("%-20v %v \n", subject, grade)
	}
	fmt.Println("-------------------------------------")
	fmt.Printf("Total AVG: %0.2f \n", avg)
	fmt.Println("-------------------------------------")
}

func main() {
	
	name := askName()
	grades := askGrade()
	avg := calculateAvg(grades)
	displayTranscript(name, grades, avg)
	
}