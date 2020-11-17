package cities

type City struct {
	Name       string
	Country    string
	Population int
	Est        int
}

type Service struct {
	Cities []*City
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) AddCity(name string, country string, population int, est int) {
	var c City
	c.Name = name
	c.Country = country
	c.Population = population
	c.Est = est
	s.Cities = append(s.Cities, &c)
	return
}
