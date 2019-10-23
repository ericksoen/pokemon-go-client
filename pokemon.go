package pokeclient

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

// GetRandomPokemon queries the Pokemon API and returns a random name
func (c *Client) GetRandomPokemon() (*PokemonName, error) {

	res, err := c.Do("GET", "pokemon_names.json")

	if err != nil {
		log.Printf("[ERROR] Pokemon client returned err %s", err)
		return nil, err
	}

	var pokemon PokemonNames

	if err := json.Unmarshal(res, &pokemon); err != nil {
		return nil, err
	}

	names := make([]PokemonName, 0)

	for _, v := range pokemon {
		names = append(names, v)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	elem := r.Intn(len(names))

	return &names[elem], nil
}
