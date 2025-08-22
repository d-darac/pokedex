package pokeapi

type PokeAPIBaseResource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PokeAPINamedList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeAPILocationArea struct {
	PokeAPIBaseResource
	GameIndex            int `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type PokeAPIPokemon struct {
	PokeAPIBaseResource
	BaseExperience int  `json:"base_experience"`
	Height         int  `json:"height"`
	IsDefault      bool `json:"is_default"`
	Order          int  `json:"order"`
	Weight         int  `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order int `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string  `json:"back_default"`
		BackFemale       *string `json:"back_female"`
		BackShiny        string  `json:"back_shiny"`
		BackShinyFemale  *string `json:"back_shiny_female"`
		FrontDefault     string  `json:"front_default"`
		FrontFemale      *string `json:"front_female"`
		FrontShiny       string  `json:"front_shiny"`
		FrontShinyFemale *string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string  `json:"front_default"`
				FrontFemale  *string `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string  `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       string  `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string  `json:"back_default"`
				BackFemale       *string `json:"back_female"`
				BackShiny        string  `json:"back_shiny"`
				BackShinyFemale  *string `json:"back_shiny_female"`
				FrontDefault     string  `json:"front_default"`
				FrontFemale      *string `json:"front_female"`
				FrontShiny       string  `json:"front_shiny"`
				FrontShinyFemale *string `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault  string `json:"back_default"`
					BackGray     string `json:"back_gray"`
					FrontDefault string `json:"front_default"`
					FrontGray    string `json:"front_gray"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"crystal"`
				Gold struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"gold"`
				Silver struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"silver"`
			} `json:"generation-ii"`
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
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string  `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        string  `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string  `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        string  `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string  `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        string  `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string  `json:"back_default"`
						BackFemale       *string `json:"back_female"`
						BackShiny        string  `json:"back_shiny"`
						BackShinyFemale  *string `json:"back_shiny_female"`
						FrontDefault     string  `json:"front_default"`
						FrontFemale      *string `json:"front_female"`
						FrontShiny       string  `json:"front_shiny"`
						FrontShinyFemale *string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string  `json:"back_default"`
					BackFemale       *string `json:"back_female"`
					BackShiny        string  `json:"back_shiny"`
					BackShinyFemale  *string `json:"back_shiny_female"`
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string  `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string  `json:"front_default"`
					FrontFemale      *string `json:"front_female"`
					FrontShiny       string  `json:"front_shiny"`
					FrontShinyFemale *string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string  `json:"front_default"`
					FrontFemale  *string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Types []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		} `json:"types"`
	} `json:"past_types"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Abilities []struct {
			Ability struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		} `json:"abilities"`
	} `json:"past_abilities"`
}

type PokeAPIPokemonSpecies struct {
	PokeAPIBaseResource
	Order                int  `json:"order"`
	GenderRate           int  `json:"gender_rate"`
	CaptureRate          int  `json:"capture_rate"`
	BaseHappiness        int  `json:"base_happiness"`
	IsBaby               bool `json:"is_baby"`
	IsLegendary          bool `json:"is_legendary"`
	IsMythical           bool `json:"is_mythical"`
	HatchCounter         int  `json:"hatch_counter"`
	HasGenderDifferences bool `json:"has_gender_differences"`
	FormsSwitchable      bool `json:"forms_switchable"`
	GrowthRate           struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"growth_rate"`
	PokedexNumbers []struct {
		EntryNumber int `json:"entry_number"`
		Pokedex     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokedex"`
	} `json:"pokedex_numbers"`
	EggGroups []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"egg_groups"`
	Color struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"color"`
	Shape struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"shape"`
	EvolvesFromSpecies struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"evolves_from_species"`
	EvolutionChain struct {
		URL string `json:"url"`
	} `json:"evolution_chain"`
	Habitat struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		AwesomeNames []struct {
			AwesomeName string `json:"awesome_name"`
			Language    struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
		} `json:"awesome_names"`
		Names []struct {
			Name     string `json:"name"`
			Language struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"language"`
		} `json:"names"`
		PokemonSpecies []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon_species"`
	} `json:"habitat"`
	Generation struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"generation"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Version struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"flavor_text_entries"`
	FormDescriptions []struct {
		Description string `json:"description"`
		Language    struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"form_descriptions"`
	Genera []struct {
		Genus    string `json:"genus"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"genera"`
	Varieties []struct {
		IsDefault bool `json:"is_default"`
		Pokemon   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"varieties"`
}
