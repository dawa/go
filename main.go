package main

import (
  "fmt"
  "strconv"
)

func sayHi() (x, y string) {

     x = "Habari"

     y = "dunia!"

     return

}

func (m *Movie) summary() string {

  r := strconv.FormatFloat(m.Rating, 'f', 1, 64)

  return m.Name + ", " + r
	
}

type Movie struct {

  Name   string

  Rating float64

}

func main() {
/*
	defer fmt.Println("I am run after the function completes")
	defer fmt.Println("I am the first defer statement")

      defer fmt.Println("I am the second defer statement")

      defer fmt.Println("I am the third defer statement")
	fmt.Println(sayHi())

	numbers := []int{1, 2, 3, 4}

      for i, n := range numbers {

          fmt.Println("The index of the loop is", i)

          fmt.Println("The value from the array is", n)

      }
	
	var cheeses = make([]string, 4)
	cheeses[0] = "Marioles"
	cheeses[1] = "Ricotta"
        cheeses[2] = "Mozarella"
	cheeses[3] = "Swiss"
	cheeses = append(cheeses, "Camembert")

	fmt.Println(len(cheeses))
	fmt.Println(cheeses)

	cheeses = append(cheeses[2:]) //, cheeses[:2]...)

        fmt.Println(len(cheeses))
        fmt.Println(cheeses)

	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses[1:])
	//copy(smellyCheeses, cheeses)
	fmt.Println(smellyCheeses)

	var players = make(map[string]int)
	players["cook"] = 32

      players["bairstow"] = 27
      players["stokes"] = 26
      fmt.Println(players["cook"])
      fmt.Println(players["bairstow"])
      fmt.Println(players)	
*/



	var m Movie
	fmt.Printf("%+v\n", m)

	m.Name = "Metropolis"

	m.Rating = 95

	fmt.Printf("%+v\n", m)

	c := Movie{Name: "Citizen Kane", Rating: 10}
	fmt.Printf("%+v\n", c)


	fmt.Println(m.summary())

}

