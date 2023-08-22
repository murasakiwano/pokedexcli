package pokeapi

type PokeApiLocationResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	Count int `json:"count"`
}

type PokeApiExploreLocationData struct {
	Location             namePlusURL           `json:"location"`
	Name                 string                `json:"name"`
	EncounterMethodRates []encounterMethodRate `json:"encounter_method_rates"`
	Names                []name                `json:"names"`
	PokemonEncounters    []pokemonEncounter    `json:"pokemon_encounters"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
}

type encounterMethodRate struct {
	EncounterMethod namePlusURL `json:"encounter_method"`
	VersionDetails  []struct {
		Version namePlusURL `json:"version"`
		Rate    int         `json:"rate"`
	} `json:"version_details"`
}

type name struct {
	Language namePlusURL `json:"language"`
	Name     string      `json:"name"`
}

type pokemonEncounter struct {
	Pokemon        namePlusURL `json:"pokemon"`
	VersionDetails []struct {
		Version          namePlusURL        `json:"version"`
		EncounterDetails []encounterDetails `json:"encounter_details"`
		MaxChance        int                `json:"max_chance"`
	} `json:"version_details"`
}

type encounterDetails struct {
	Method          namePlusURL `json:"method"`
	ConditionValues []any       `json:"condition_values"`
	Chance          int         `json:"chance"`
	MaxLevel        int         `json:"max_level"`
	MinLevel        int         `json:"min_level"`
}

type namePlusURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
