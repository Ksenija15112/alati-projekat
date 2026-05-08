package store
import "alati-projekat/models"

type Store struct{
	listaKonfiguracija []models.Config
}

func(s *Store) dodavanjeConfig(novi models.Config){

	s.listaKonfiguracija=append(s.listaKonfiguracija,novi)

}

func(s *Store) citanjeConfig()[]models.Config{
	return s.listaKonfiguracija
}