package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	config := Config{
		APIKey: os.Getenv("POKE_API_KEY"),
	}

	client := New(config)

	log.Printf(config.APIKey)
	log.Printf("[INFO] Pokemon client configured")

	res, err := client.Do("GET", "pokemon_names.json")

	if err != nil {
		log.Printf("[ERROR] Pokemon client returned err %s", err)
	}

	var pokemon PokemonNames

	if err := json.Unmarshal(res, &pokemon); err != nil {
		panic(err)
	}

	names := make([]string, 0)

	for _, v := range pokemon {
		names = append(names, v.Name)
		// fmt.Println(k)
		// fmt.Println(v)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	elem := r.Intn(len(names))

	fmt.Println(names[elem])

}
