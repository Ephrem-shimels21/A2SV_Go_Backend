// package main

// import (
// 	"fmt"
// )

// func askNameAndSubject() (string, int) {
// 	var name string
// 	var noSubjects int

// 	println("Enter your name, please: ")
// 	fmt.Scan(&name)
// 	println("Enter the number of subjects you are taking: ")
// 	fmt.Scan(&noSubjects)

// 	return name, noSubjects

// }

// func calculateAverageSubjects(noSubjects int) map[string]int {
// 	var subjectName string
// 	var grade int
// 	var subjectGradeMap = make(map[string]int)

// 	for i := 0; i < noSubjects; i++ {

// 		println("Enter the name of the subject: ")
// 		fmt.Scan(&subjectName)
// 		println("Enter the grade for the subject: ")
// 		fmt.Scan(&grade)

// 		for grade < 0 || grade > 100 {
// 			println("please enter a valid grade: ")
// 			fmt.Scan(&grade)

// 		}
// 		subjectGradeMap[subjectName] = grade

// 	}

// 	return subjectGradeMap

// }

// func main() {
// 	name, noSubjects := askNameAndSubject()
// 	subjectGradeMap := calculateAverageSubjects(noSubjects)

// 	var totalGrade int
// 	for _, grade := range subjectGradeMap {
// 		totalGrade += grade
// 	}

// 	average := totalGrade / noSubjects
// 	fmt.Printf("The Name of the student is %s: \n", name)
// 	println("===========================================")
// 	println("Here are the grades for the subjects: ")
// 	for key, val := range subjectGradeMap {
// 		fmt.Printf("The grade for %s is %d\n", key, val)

// 	}
// 	println("===========================================")
// 	fmt.Printf("The average grade for %s is %d ", name, average)
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func askNameAndSubject(reader *bufio.Reader) (string, int, error) {
	var name string
	var noSubjects int

	fmt.Println("Enter your name, please: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter the number of subjects you are taking: ")
	noSubjectsStr, _ := reader.ReadString('\n')
	noSubjectsStr = strings.TrimSpace(noSubjectsStr)
	noSubjects, err := strconv.Atoi(noSubjectsStr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid number of subjects")
	}

	return name, noSubjects, nil
}

func calculateAverageSubjects(reader *bufio.Reader, noSubjects int) (map[string]int, error) {
	subjectGradeMap := make(map[string]int)

	for i := 0; i < noSubjects; i++ {
		fmt.Println("Enter the name of the subject: ")
		subjectName, _ := reader.ReadString('\n')
		subjectName = strings.TrimSpace(subjectName)

		fmt.Println("Enter the grade for the subject: ")
		gradeStr, _ := reader.ReadString('\n')
		gradeStr = strings.TrimSpace(gradeStr)
		grade, err := strconv.Atoi(gradeStr)
		if err != nil || grade < 0 || grade > 100 {
			return nil, fmt.Errorf("invalid grade")
		}

		subjectGradeMap[subjectName] = grade
	}

	return subjectGradeMap, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	name, noSubjects, err := askNameAndSubject(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	subjectGradeMap, err := calculateAverageSubjects(reader, noSubjects)
	if err != nil {
		fmt.Println(err)
		return
	}

	var totalGrade int
	for _, grade := range subjectGradeMap {
		totalGrade += grade
	}

	average := totalGrade / noSubjects
	fmt.Printf("The Name of the student is %s:\n", name)
	fmt.Println("===========================================")
	fmt.Println("Here are the grades for the subjects: ")
	for key, val := range subjectGradeMap {
		fmt.Printf("The grade for %s is %d\n", key, val)
	}
	fmt.Println("===========================================")
	fmt.Printf("The average grade for %s is %d\n", name, average)
}
