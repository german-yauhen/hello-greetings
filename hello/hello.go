package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	names := []string{"Eugene", "Viki", "Qualcuno"}
	// greetingMsg, errMsg := greetings.Greeting("Eugene")
	greetingMsgs, errMsg := greetings.Greetings(names)

	if errMsg != nil {
		// Set properties of the predefined Logger, including
		// the log entry prefix and a flag to disable printing
		// the time, source file, and line number.
		log.SetPrefix("greetings: ")
		log.SetFlags(0)
		log.Fatal(errMsg)
	}

	fmt.Println(greetingMsgs)

	greetingForEugene, isExist := greetingMsgs["Eugene"]
	fmt.Println(isExist)
	fmt.Println(greetingForEugene)
}
