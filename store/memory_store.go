package store
import "alati-projekat/models"
import  "github.com/google/uuid"
import "errors"

type Store struct{
	listaKonfiguracija []models.Config
	listaGrupa []models.KonfiguracionaGrupa
}

func(s *Store) DodavanjeConfig(novi models.Config)error{
    for _, v := range s.listaKonfiguracija {
        if v.Naziv == novi.Naziv && v.Verzija == novi.Verzija {
            return errors.New("konfiguracija sa tim nazivom i verzijom već postoji")
        }
    }
	
	novi.ID = uuid.New().String()
	s.listaKonfiguracija=append(s.listaKonfiguracija,novi)
    return nil
}

func(s *Store) CitanjeConfig()[]models.Config{
	return s.listaKonfiguracija
}

func(s *Store) FindById(ID string)(models.Config,error){
    for _,v:=range s.listaKonfiguracija{
		if(ID==v.ID){
           return v,nil
		}
	}
	return models.Config{},errors.New("nije pronadjen id")
}

func(s *Store)Delete(ID string)error{
	for i,v:=range s.listaKonfiguracija{
         if(ID==v.ID){
			s.listaKonfiguracija=append(s.listaKonfiguracija[:i],s.listaKonfiguracija[i+1:]...)
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

func(s *Store) FindByNazivIVerzija(Naziv,Verzija string)(models.Config,error){
    for _,v:=range s.listaKonfiguracija{
		if(v.Naziv==Naziv && v.Verzija==Verzija){
           return v,nil
		}
	}
	return models.Config{},errors.New("nije pronadjena konfiguracija sa tim nazivom i verzijom")
}