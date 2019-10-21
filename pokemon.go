package pokeclient

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

// GetRandomPokemon queries the Pokemon API and returns a random name
func (c *Client) GetRandomPokemon() (string, error) {

	res, err := c.Do("GET", "pokemon_names.json")

	if err != nil {
		log.Printf("[ERROR] Pokemon client returned err %s", err)
		return "", err
	}

	var pokemon PokemonNames

	if err := json.Unmarshal(res, &pokemon); err != nil {
		return "", err
	}

	names := make([]string, 0)

	for _, v := range pokemon {
		names = append(names, v.Name)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	elem := r.Intn(len(names))

	return names[elem], nil
}
