package pokeapi

type Pokemon struct {
	Sprites struct {
		Versions struct {
			GenerationIv struct {
				DiamondPearl struct {
					BackFemale       *string `json:"back_female"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontFemale      *string `json:"front_female"`
					FrontShinyFemale *string `json:"front_shiny_female"`
					BackDefault      string  `json:"back_default"`
					BackShiny        string  `json:"back_shiny"`
					FrontDefault     string  `json:"front_default"`
					FrontShiny       string  `json:"front_shiny"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackFemale       *string `json:"back_female"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontFemale      *string `json:"front_female"`
					FrontShinyFemale *string `json:"front_shiny_female"`
					BackDefault      string  `json:"back_default"`
					BackShiny        string  `json:"back_shiny"`
					FrontDefault     string  `json:"front_default"`
					FrontShiny       string  `json:"front_shiny"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackFemale       *string `json:"back_female"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontFemale      *string `json:"front_female"`
					FrontShinyFemale *string `json:"front_shiny_female"`
					BackDefault      string  `json:"back_default"`
					BackShiny        string  `json:"back_shiny"`
					FrontDefault     string  `json:"front_default"`
					FrontShiny       string  `json:"front_shiny"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackFemale       *string `json:"back_female"`
						BackShinyFemale  *string `json:"back_shiny_female"`
						FrontFemale      *string `json:"front_female"`
						FrontShinyFemale *string `json:"front_shiny_female"`
						BackDefault      string  `json:"back_default"`
						BackShiny        string  `json:"back_shiny"`
						FrontDefault     string  `json:"front_default"`
						FrontShiny       string  `json:"front_shiny"`
					} `json:"animated"`
					BackFemale       *string `json:"back_female"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontFemale      *string `json:"front_female"`
					FrontShinyFemale *string `json:"front_shiny_female"`
					BackDefault      string  `json:"back_default"`
					BackShiny        string  `json:"back_shiny"`
					FrontDefault     string  `json:"front_default"`
					FrontShiny       string  `json:"front_shiny"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontFemale      *string `json:"front_female"`
					FrontShinyFemale *string `json:"front_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontShiny       string  `json:"front_shiny"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontFemale      *string `json:"front_female"`
					FrontShinyFemale *string `json:"front_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontShiny       string  `json:"front_shiny"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontFemale  *string `json:"front_female"`
					FrontDefault string  `json:"front_default"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontFemale      *string `json:"front_female"`
					FrontShinyFemale *string `json:"front_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontShiny       string  `json:"front_shiny"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontFemale  *string `json:"front_female"`
					FrontDefault string  `json:"front_default"`
				} `json:"icons"`
			} `json:"generation-viii"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
		} `json:"versions"`
		BackFemale       *string `json:"back_female"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontFemale      *string `json:"front_female"`
		FrontShinyFemale *string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontFemale  *string `json:"front_female"`
				FrontDefault string  `json:"front_default"`
			} `json:"dream_world"`
			Home struct {
				FrontFemale      *string `json:"front_female"`
				FrontShinyFemale *string `json:"front_shiny_female"`
				FrontDefault     string  `json:"front_default"`
				FrontShiny       string  `json:"front_shiny"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
		} `json:"other"`
		BackDefault  string `json:"back_default"`
		BackShiny    string `json:"back_shiny"`
		FrontDefault string `json:"front_default"`
		FrontShiny   string `json:"front_shiny"`
	} `json:"sprites"`
	Species                NamedAPIResource `json:"species"`
	LocationAreaEncounters string           `json:"location_area_encounters"`
	Name                   string           `json:"name"`
	PastTypes              []struct {
		Generation NamedAPIResource `json:"generation"`
		Types      []PokemonType    `json:"types"`
	} `json:"past_types"`
	Abilities []struct {
		Ability  NamedAPIResource `json:"ability"`
		IsHidden bool             `json:"is_hidden"`
		Slot     int              `json:"slot"`
	} `json:"abilities"`
	Types []struct {
		Type NamedAPIResource `json:"type"`
		Slot int              `json:"slot"`
	} `json:"types"`
	Stats []struct {
		Stat     NamedAPIResource `json:"stat"`
		BaseStat int              `json:"base_stat"`
		Effort   int              `json:"effort"`
	} `json:"stats"`
	Forms []NamedAPIResource `json:"forms"`
	Moves []struct {
		Move                NamedAPIResource `json:"move"`
		VersionGroupDetails []struct {
			MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
			VersionGroup    NamedAPIResource `json:"version_group"`
			LevelLearnedAt  int              `json:"level_learned_at"`
		} `json:"version_group_details"`
	} `json:"moves"`
	GameIndices []struct {
		Version   NamedAPIResource `json:"version"`
		GameIndex int              `json:"game_index"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item           NamedAPIResource `json:"item"`
		VersionDetails []struct {
			Version NamedAPIResource `json:"version"`
			Rarity  int              `json:"rarity"`
		} `json:"version_details"`
	} `json:"held_items"`
	Order          int  `json:"order"`
	Height         int  `json:"height"`
	BaseExperience int  `json:"base_experience"`
	ID             int  `json:"id"`
	Weight         int  `json:"weight"`
	IsDefault      bool `json:"is_default"`
}

type PokemonType struct {
	Type NamedAPIResource `json:"type"`
	Slot int              `json:"slot"`
}
