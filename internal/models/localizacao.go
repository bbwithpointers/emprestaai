package models

type Localizacao struct {
	Latitude  int `json:"lat"`
	Longitude int `json:"long"`
}

func (l *Localizacao) Get() string {
	// TODO : retorna um ponto no mapa
	return "localicazao"
}
