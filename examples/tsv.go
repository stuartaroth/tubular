package main

import (
	"fmt"
	"github.com/stuartaroth/tubular"
)

func main() {
	pages := []tubular.Page{
		tubular.Page{
			Name: "Horror",
			Rows: [][]interface{}{
				[]interface{}{"Title", "Director", "Year Released"},
				[]interface{}{"A Nightmare on Elm Street", "Wes Craven", 1984},
				[]interface{}{"Black Christmas", "Bob Clark", 1974},
				[]interface{}{"Friday the 13th", "Sean S. Cunningham", 1980},
				[]interface{}{"Halloween", "John Carpenter", 1978},
				[]interface{}{"The Texas Chain Saw Massacre", "Tobe Hooper", 1974},
			},
		},
		tubular.Page{
			Name: "Sci-Fi",
			Rows: [][]interface{}{
				[]interface{}{"Title", "Director", "Year Released"},
				[]interface{}{"2001: A Spacey Odyssey", "Stanley Kubrick", 1968},
				[]interface{}{"Aliens", "James Cameron", 1986},
				[]interface{}{"Back to the Future", "Robert Zemeckis", 1985},
				[]interface{}{"Blade Runner", "Ridley Scott", 1982},
				[]interface{}{"Star Wars", "George Lucas", 1977},
			},
		},
	}

	filename, err := tubular.WriteSeparatedValuesFile("Movies", "tsv", "\n", "\t", pages)
	if err != nil {
		fmt.Println("Error writing separated values file:", err)
		return
	}

	fmt.Println("Successfully wrote separated values file:", filename)
}
