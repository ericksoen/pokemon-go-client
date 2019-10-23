package pokeclient

// PokemonName models the response object from the API
type PokemonName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PokemonNames resolves the dynamic property structure returned from the API, e.g.,
// "1": {"id": 1, "name": "bulbasaur"}, "2"...
type PokemonNames map[string]PokemonName
