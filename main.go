package main

import (
	"bufio"
	"errors"
	"fmt"
	"ignacio83/grettings-go/grettings"
	"log"
	"os"
	"rsc.io/quote"
	"strings"
)

func init() {
	log.SetPrefix("grettings: ")
	log.SetFlags(0)
}

func main() {
	fmt.Println(quote.Go())
	fmt.Println()

	names, err := requestAndReadNames()
	if err != nil {
		log.Fatal(err)
	}

	hellos, err := grettings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range hellos {
		fmt.Printf("Input name:%s, Greeting: %s\n", k, v)
	}
}

func requestAndReadNames() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please, type the names, you should separate then with commas (,): ")

	scanner.Scan()
	namesInput := scanner.Text()

	if len(namesInput) == 0 {
		return nil, errors.New("input is empty")
	}
	names := strings.Split(namesInput, ",")
	for i := range names {
		names[i] = strings.TrimSpace(names[i])
	}
	return names, nil
}
