package main


type Person struct {
    firstName string
    lastName  string
    age       int
}


func (p Person) fullName() string {
    return p.firstName + " " + p.lastName
}


    
