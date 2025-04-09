package main

import "fmt"

func main(){
	var studentGardes = make(map[string]float32)
	var studentName string
	var numberOfSubjects int
	fmt.Println("Welcome to displaying grades of student")
	fmt.Println("Enter the name of the student:")
	fmt.Scanln(&studentName)
	fmt.Println("Enter the number of subjects:")
	fmt.Scanln(&numberOfSubjects)
	for i:=0; i<numberOfSubjects;i++{
		var subject string
		var grade float32
		fmt.Println("Enter the name of subject ",i+1,":")
		fmt.Scanln(&subject)
		fmt.Println("Enter your Grade:")
		fmt.Scanln(&grade)
		for grade<0 || grade>100 {
			fmt.Println("Entered grade must be between 0 and 100. , Please insert again.")
			fmt.Scanln(&grade)
		}
		studentGardes[subject] = grade
	}
	displayGrades(studentGardes)
	calculateGrade(studentGardes)

	


}

func displayGrades(studentGardes map[string]float32){
	fmt.Println("This are your Grades")
	for key,value := range(studentGardes){
		fmt.Println(key,":",value)
	}
	

}

func calculateGrade(studentGardes map[string]float32){
	fmt.Println("Calculating your avarage score")
	var avarage float32
	var summ float32
	for _,value := range(studentGardes){
		summ += value
	}
	avarage = summ / float32(len(studentGardes))
	fmt.Println("Your avarage score is:",avarage)

}