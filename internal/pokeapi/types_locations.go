package pokeapi

type PokeApiLocationsResponse struct {
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
	Count    int                `json:"count"`
}

type PokeApiLocationArea struct {
	Location             NamedAPIResource      `json:"location"`
	Name                 string                `json:"name"`
	EncounterMethodRates []encounterMethodRate `json:"encounter_method_rates"`
	Names                []name                `json:"names"`
	PokemonEncounters    []pokemonEncounter    `json:"pokemon_encounters"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
}

type encounterMethodRate struct {
	EncounterMethod NamedAPIResource `json:"encounter_method"`
	VersionDetails  []struct {
		Version NamedAPIResource `json:"version"`
		Rate    int              `json:"rate"`
	} `json:"version_details"`
}

type name struct {
	Language NamedAPIResource `json:"language"`
	Name     string           `json:"name"`
}

type pokemonEncounter struct {
	Pokemon        NamedAPIResource `json:"pokemon"`
	VersionDetails []struct {
		Version          NamedAPIResource `json:"version"`
		EncounterDetails []encounter      `json:"encounter_details"`
		MaxChance        int              `json:"max_chance"`
	} `json:"version_details"`
}

type encounter struct {
	Method          NamedAPIResource   `json:"method"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	MaxLevel        int                `json:"max_level"`
	MinLevel        int                `json:"min_level"`
}
