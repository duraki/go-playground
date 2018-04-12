package main

import "fmt"

type Saiyan struct {
	Name string
	Power int
}

func main()  {
	// Address of operator Saiyan
	goku := &Saiyan{
		Name: "Goku",
		Power: 9000,
	}

	//Super(goku)
	goku.Super()
	fmt.Println(goku)
}

// Allocation example (both results are same)
func allocate() {
	/**
	goku := new(Saiyan) // allocates memory required for type Saiyan
	gohan := &Saiyan{} // allocates memory required for type Saiyan
	**/
}

// Enforce super on Saiyan. Expecting pointer to value of type Saiyan.
func Super(s *Saiyan) {
	//s.Power += 10000
	s = &Saiyan{"Gohan", 1000}
}


// Functions on Structures
func (s *Saiyan) Super() {
	s.Power += 10000
}

// Structures constructor is created this way
//
// Note that we may use *Saiyan + &Saiyan to return pointer reference, or
// Saiyan + Saiyan to return object of Saiyan type
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name: name,
		Power: power,
	}
}