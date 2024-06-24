package main

import "fmt"

func main() {

	//HOW TO CREATE A MAP
	// englishTutorduDict is a map where the keys are string and the values are also strings
	//to specify that thr values are strings, we put the strings inside the []
	//to specify that thr values are strings, we put the strings after the []
	//if we want to start with a pair, we can have it inside the curly barces
	// MAPS HAVE NO SIZE OR LIMITS TO THEM

	/* 
	
	======== BEST WAY TO INTERACT WITH A USER TO FIND OUT IF SOMETHING EXIST OR NOT ========
	
	// when something doesn't exist in your library use a for loop with options if and else

	value, ok = englishTutorduDict["book"]

	if ok == true{
		fmt.Println("the key money exists with the value", value)
		delete(englishTutorduDict,"money")
	}else{
		fmt.Println("there is no pair with the key money")
	}
	
	*/

	var englishTutorduDict = map[string]string{"love": "pyar"} // DEPENDING ON WHAT VALUE YOU WILL USE

	//HOW TO ADD NEW PAIR TO OUR APP

	englishTutorduDict["friendship"] = "dosti"

	englishTutorduDict["calm"] = "sabr"

	englishTutorduDict["money"] = "paisa"

	englishTutorduDict["book"] = "kitab"

	// HOW TO REMOVE A PAIR WITH A CERTAIN KEY

	value, ok := englishTutorduDict["money"]  // Searching/checking if our dictonary a pair with "money"

	value, ok = englishTutorduDict["book"]

	// when something doesn't exist in your library use a for loop with options if and else
	if ok == true{
		fmt.Println("the key money exists with the value", value)
		delete(englishTutorduDict,"money")
	}else{
		fmt.Println("there is no pair with the key money")
	}

	//HOW TO PRINT DICTIONARY

	fmt.Println(englishTutorduDict)

	// using the RANGE formula to help us iterate the map.
	for key, value := range englishTutorduDict {
		fmt.Println(key, ":", value)
		fmt.Printf("%s : %s\n", key, value)
	}

}
