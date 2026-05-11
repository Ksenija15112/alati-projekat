package store

import (
	"alati-projekat/models"
	"errors"

	"github.com/google/uuid"
)

type Store struct {
	listaKonfiguracija []models.Config
	listaGrupa         []models.KonfiguracionaGrupa
}

func (s *Store) DodavanjeConfig(novi models.Config) error {
	for _, v := range s.listaKonfiguracija {
		if v.Naziv == novi.Naziv && v.Verzija == novi.Verzija {
			return errors.New("konfiguracija sa tim nazivom i verzijom već postoji")
		}
	}

	novi.ID = uuid.New().String()
	s.listaKonfiguracija = append(s.listaKonfiguracija, novi)
	return nil
}

func (s *Store) CitanjeConfig() []models.Config {
	return s.listaKonfiguracija
}

func (s *Store) FindById(ID string) (models.Config, error) {
	for _, v := range s.listaKonfiguracija {
		if ID == v.ID {
			return v, nil
		}
	}
	return models.Config{}, errors.New("nije pronadjen id")
}

func (s *Store) Delete(ID string) error {
	for i, v := range s.listaKonfiguracija {
		if ID == v.ID {
			s.listaKonfiguracija = append(s.listaKonfiguracija[:i], s.listaKonfiguracija[i+1:]...)
			return nil
		}
	}

	return errors.New("nije pronadjen ID")
}

func (s *Store) Update(ID string, nova models.Config) error {
	for _, v := range s.listaKonfiguracija {
		if ID == v.ID {
			return s.DodavanjeConfig(nova)
		}
	}

	return errors.New("nije pronadjen ID")
}

func (s *Store) FindByNazivIVerzija(Naziv, Verzija string) (models.Config, error) {
	for _, v := range s.listaKonfiguracija {
		if v.Naziv == Naziv && v.Verzija == Verzija {
			return v, nil
		}
	}
	return models.Config{}, errors.New("nije pronadjena konfiguracija sa tim nazivom i verzijom")

}
func (s *Store) DodavanjeGrupe(nova models.KonfiguracionaGrupa) error {
	for _, v := range s.listaGrupa {
		if v.Naziv == nova.Naziv && v.Verzija == nova.Verzija {
			return errors.New("grupa sa tim nazivom i verzijom već postoji")
		}
	}
	nova.ID = uuid.New().String()
	s.listaGrupa = append(s.listaGrupa, nova)
	return nil
}

func (s *Store) FindGrupaByNazivIVerzija(naziv, verzija string) (models.KonfiguracionaGrupa, error) {
	for _, v := range s.listaGrupa {
		if v.Naziv == naziv && v.Verzija == verzija {
			return v, nil
		}
	}
	return models.KonfiguracionaGrupa{}, errors.New("grupa nije pronađena")
}

func (s *Store) DeleteGrupa(ID string) error {
	for i, v := range s.listaGrupa {
		if v.ID == ID {
			s.listaGrupa = append(s.listaGrupa[:i], s.listaGrupa[i+1:]...)
			return nil
		}
	}
	return errors.New("grupa nije pronađena")
}

func (s *Store) DodavanjeConfigUGrupu(naziv, verzija string, config models.Config) error {
	for i, v := range s.listaGrupa {
		if v.Naziv == naziv && v.Verzija == verzija {
			s.listaGrupa[i].ListaKonfiguracija = append(s.listaGrupa[i].ListaKonfiguracija, config)
			return nil
		}
	}
	return errors.New("grupa nije pronađena")
}

func (s *Store) UklanjranjeConfigIzGrupe(naziv, verzija, configNaziv string) error {
	for i, v := range s.listaGrupa {
		if v.Naziv == naziv && v.Verzija == verzija {
			for j, c := range s.listaGrupa[i].ListaKonfiguracija {
				if c.Naziv == configNaziv {
					s.listaGrupa[i].ListaKonfiguracija = append(
						s.listaGrupa[i].ListaKonfiguracija[:j],
						s.listaGrupa[i].ListaKonfiguracija[j+1:]...,
					)
					return nil
				}
			}
			return errors.New("konfiguracija nije pronađena u grupi")
		}
	}
	return errors.New("grupa nije pronađena")
}
