package store
import "alati-projekat/models"
import  "github.com/google/uuid"
import "errors"

type Store struct{
	listaKonfiguracija []models.Config
}

func(s *Store) dodavanjeConfig(novi models.Config){
    novi.ID = uuid.New().String()
	s.listaKonfiguracija=append(s.listaKonfiguracija,novi)

}

func(s *Store) citanjeConfig()[]models.Config{
	return s.listaKonfiguracija
}

func(s *Store) findById(ID string)(models.Config,error){
    for _,v:=range s.listaKonfiguracija{
		if(ID==v.ID){
           return v,nil
		}
	}
	return models.Config{},errors.New("nije pronadjen id")
}

func(s *Store)brisanje(ID string)error{
	for i,v:=range s.listaKonfiguracija{
         if(ID==v.ID){
			s.listaKonfiguracija=append(s.listaKonfiguracija[:i],s.listaKonfiguracija[i+1:]...)
			return nil
		 }
	}
  
  return errors.New("nije pronadjen ID")
}