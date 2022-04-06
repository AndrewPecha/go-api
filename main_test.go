package main

import (
	"testing"
)

func TestGetPeopleData(t *testing.T) {
	var people []Person
	people = getPeopleData()
	sut := people[1].FirstName
	if sut != "Andrew" {
		t.Error("Expected Andrew, got", sut)
	}
}
