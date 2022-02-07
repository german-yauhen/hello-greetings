package greetings

import (
	"regexp"
	"testing"
)

/*
Calls greetings.Greeting with a name, checking for a valid return value.
*/
func TestGreetingName(t *testing.T) {
	name := "Eugene"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greeting(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatal("a name expected, however an error occured")
	}
}

/*
Calls greetings.Greeting with an empty name, checking for an error.
*/
func TestGreetingEmpty(t *testing.T) {
	msg, err := Greeting("")
	if len(msg) > 0 || err == nil {
		t.Fatalf("an error expected")
	}
}

/*
Calls greetings.Greeting with an name that contains numbers. Intentionally failed
*/
func TestGreetingNameContainsNumbers(t *testing.T) {
	msg, err := Greeting("Eug3n3")
	if len(msg) == 0 || err != nil {
		t.Fatalf("a name expected, however an error occured: %s", err)
	}
}