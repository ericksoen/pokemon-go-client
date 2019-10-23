## Overview

Use the [RapidAPI PokemonGo API](https://rapidapi.com/brianiswu/api/pokemon-go1?endpoint=apiendpoint_af8fc313-bc8c-4b86-8671-83a2212bdefc) to retrieve a random Pokemon character. The PokemonGo API requires an API key (free) to use, so you'll need to sign up for one before you can use the client.

If you've not done Go development previously, the [Learn Go With Tests](https://github.com/quii/learn-go-with-tests) repository has some clear instructions for setting up your environment.

## Usage

To use in your code, create a configuration object and then pass it as a parameter when creating a new instance of the API client:

```go
config := Config{
    APIKey: os.Getenv("POKE_API_KEY"),
    BaseUrl: os.Getenv("POKE_BASE_URL")
}

client := NewClient(config)

character := client.GetRandomPokemon()
```

## Thanks

This code adopts many of it's patterns from the [NewRelic Go client](https://github.com/paultyng/go-newrelic/blob/master/api/client.go)
