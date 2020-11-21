package cities

type City struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Est        int    `json:"est"`
}

var Cities = []City{
	{
		Name:       "Moscow",
		Country:    "Russia",
		Population: 14_000_000,
		Est:        1147,
	},
	{
		Name:       "Kiev",
		Country:    "Ukraine",
		Population: 4_000_000,
		Est:        882,
	},
	{
		Name:       "Minsk",
		Country:    "Belarus",
		Population: 2_000_000,
		Est:        1067,
	},
	{
		Name:       "Vilnius",
		Country:    "Lietuva",
		Population: 500_000,
		Est:        1323,
	},
}

func AllCities() []City {
	result := make([]City, 0)
	for _, c := range Cities {
		result = append(result, c)
	}
	return result
}
