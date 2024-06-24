package main

import (
	"encoding/csv" // Specific file to use csv
	"fmt"
	"os" //Without os we cannot open a file
	"strconv"
)

type student struct { // create a struct for student
	loginName  string
	firstName  string
	familyName string
	level      int
	class      string
	bestresult int
}

func main() {

	file, err := os.Open("Students.csv") // How to open csv file

	Students := []student{}

	if err != nil { // if there is an error you must create a panic to show an error
		fmt.Println("Error opening the file", err)
		return
	}
	reader := csv.NewReader(file) // How to create a reader to csv file

	reader.Read()

	for { // infinite loop which doesnt stop only ends if you do a break

		line, err := reader.Read() //How to read a line in csv file, make sure you create variable

		if err != nil { // active err variable
			break
		}

		var newstudent student
		

		newstudent.loginName = line[0] // reading field by field from each line
		newstudent.firstName = line[1]
		newstudent.familyName = line[2]
		newstudent.level, err = strconv.Atoi(line[3]) // if its something that has "INT" in it you use strconv
		if err != nil {
			fmt.Println("There was a problem in conversion", err)
			newstudent.level = 0
		}
		newstudent.class = line[4]
		newstudent.bestresult, err = strconv.Atoi(line[5])
		if err != nil {
			fmt.Println("There was a problem in conversion", err)
			newstudent.level = 0
		}
		if newstudent.firstName[0] == 'A' && newstudent.level == 3 {
			Students = append(Students, newstudent)
		}
	}

	newFile, err := os.Create("newStudent.csv")
	if err != nil { // if there is an error you must create a panic to show an error
		fmt.Println("Error opening the file", err)
		return
	}
defer newFile.Close()

writer := csv.NewWriter(newFile) //writer is a variable that will allow us to write to the file
writer.Flush()

writer.Write([]string{"Login Name", "First Name", "Family Name","Level", "Class" ,"Best Result"})

for _, student := range Students {
	row := []string {
		student.loginName,
		student.firstName,
		student.familyName,
		strconv.Itoa(student.level),
		student.class,
		strconv.Itoa(student.bestresult),
	}
	writer.Write(row)
}

fmt.Println("Filtered students saved to newStudents.csv successfully.")

for _, student := range Students {
	fmt.Printf("The student %s was imported from the csv file to the program\n", student.loginName)

}
}


/*

	for _, student := range Students {
		fmt.Printf("The student %s was imported from the csv file to the program\n", student.loginName)

		This challenge is for you to "take" all the students that start with letter A in their first name and
		are in level 3 and save them on a new file (you will need to create it) called newStudents.csv

		TIP: REMEMBER THAT THE FIRST LINE IN A CSV FILE HAS TO BE THE NAMES OF COLUMNS


*/
