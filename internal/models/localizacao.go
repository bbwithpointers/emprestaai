package models

type Localizacao struct {
	Latitude  int
	Longitude int
}

func (l *Localizacao) Get() {
	// TODO : retorna um ponto no mapa
}
