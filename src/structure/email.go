package st

import "fmt"

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func (s *EmailData) ToString() string {
	return fmt.Sprintf("%s;%s;%d\n", s.Country, s.Provider, s.DeliveryTime)
}
