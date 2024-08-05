package main

import (
	"fmt"
)

func askNameAndSubject() (string, int) {
	var name string
	var noSubjects int

	println("Enter your name, please: ")
	fmt.Scan(&name)
	println("Enter the number of subjects you are taking: ")
	fmt.Scan(&noSubjects)

	return name, noSubjects

}

func calculateAverageSubjects(noSubjects int) map[string]int {
	var subjectName string
	var grade int
	var subjectGradeMap = make(map[string]int)

	for i := 0; i < noSubjects; i++ {

		println("Enter the name of the subject: ")
		fmt.Scan(&subjectName)
		println("Enter the grade for the subject: ")
		fmt.Scan(&grade)

		for grade < 0 || grade > 100 {
			println("please enter a valid grade: ")
			fmt.Scan(&grade)

		}
		subjectGradeMap[subjectName] = grade

	}

	return subjectGradeMap

}

func main() {
	name, noSubjects := askNameAndSubject()
	subjectGradeMap := calculateAverageSubjects(noSubjects)

	var totalGrade int
	for _, grade := range subjectGradeMap {
		totalGrade += grade
	}

	average := totalGrade / noSubjects
	fmt.Printf("The Name of the student is %s: \n", name)
	println("===========================================")
	println("Here are the grades for the subjects: ")
	for key, val := range subjectGradeMap {
		fmt.Printf("The grade for %s is %d\n", key, val)

	}
	println("===========================================")
	fmt.Printf("The average grade for %s is %d ", name, average)
}
