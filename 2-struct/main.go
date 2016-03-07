package main

import (
  "fmt"
  "encoding/json"
  "time"
)

type user struct {
  Name string `json:"name"`
  DOB string `json:"date_of_birth"`
  City string `json:"city"`
}

func (u *user) sayHello() {
  fmt.Printf("Hello %s\n", u.Name)
}

func (u *user) saySentence() {
  dob, err2 := time.Parse("January 2, 2006", u.DOB)
  if err2 != nil {
    fmt.Printf("Date is not in the correct format %v", err2)
  }

  elapsedTime := time.Since(dob)
  years := elapsedTime / time.Hour / 24 / 365
  fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, years)
  }
func main() {
  u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
  _, err := json.Marshal(u)
    if err != nil {
      fmt.Printf("Cannot marshal %v", err)
    }
    u.sayHello()
    u.saySentence()
}
