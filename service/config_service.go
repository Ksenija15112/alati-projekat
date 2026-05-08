package service

import ("errors"
         "alati-projekat/models")

func validacijaNaziva(naziv string) error {
	if naziv == "" {
		return errors.New("ne sem naziv da bude prazan")
	}else{
		return nil
	}
}

func validacijaVerzije(verzija string)error{
	if verzija==""{
		return errors.New("ne sme naziv verzije da bude prazan")
	}else{
		return nil
	}
}

func validacijaParametara(parametri []models.Parametar)error{
	for _,v:=range parametri{
		if(v.Kljuc==""|| v.Vrednost==""){
			return errors.New("mora biti uneto")
		}
	}
	return nil

}

func validacijaLabela(listaLabela []models.Labela)error{
	for _,v:=range listaLabela{
		if(v.Kljuc==""|| v.Vrednost==""){
			return errors.New("mora biti uneto")
		}
	}
	return nil
}