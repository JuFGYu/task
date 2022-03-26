package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type RandomName struct {
	First_name string `json: "first_name"`
	Last_name  string `json: "last_name"`
}

type Joke struct {
	Value Value `json: "value"`
}
type Value struct {
	Joke string `json: "joke"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		name := getRandomName()
		joke := getJoke()
		newJoke := addRandomName(name, joke)

		log.Print(newJoke)

		fmt.Fprintf(w, newJoke)

	})

	http.ListenAndServe(":5000", nil)
}

func getRandomName() RandomName {
	response, err := http.Get(`https://names.mcquay.me/api/v0`)
	if err != nil {
		log.Fatal(err)
	}

	bytes, errRead := ioutil.ReadAll(response.Body)

	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(err)
		}
	}()

	if errRead != nil {
		log.Fatal(errRead)
	}

	var nameResponse RandomName

	errUnmarshal := json.Unmarshal(bytes, &nameResponse)

	if errUnmarshal != nil {
		log.Fatal(err)
	}
	log.Print(nameResponse)

	return nameResponse
}

func getJoke() Joke {
	response, err := http.Get(`http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy`)
	if err != nil {
		log.Fatal(err)
	}

	bytes, errRead := ioutil.ReadAll(response.Body)

	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(err)
		}
	}()

	if errRead != nil {
		log.Fatal(errRead)
	}

	var jokeResponse Joke

	errUnmarshal := json.Unmarshal(bytes, &jokeResponse)

	if errUnmarshal != nil {
		log.Fatal(err)
	}

	log.Print(jokeResponse)
	return jokeResponse
}

func addRandomName(randomName RandomName, joke Joke) string {
	first_attempt := strings.Replace(joke.Value.Joke, "John Doe", randomName.First_name+" "+randomName.Last_name, 1)
	return first_attempt
}
