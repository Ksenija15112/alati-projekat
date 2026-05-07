package models

type Config struct{
	Naziv string
	Verzija string
	Parametri []Parametar
}

type Parametar struct{
	Kljuc string
	Vrednost string
}

  