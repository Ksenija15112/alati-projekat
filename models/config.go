package models

type Config struct{
	Naziv string
	Verzija string
	Parametri []Parametar
	ListaLabela []Labela
	ID string
}

type Parametar struct{
	Kljuc string
	Vrednost string
}
type Labela struct{
	Kljuc string
	Vrednost string
}

  