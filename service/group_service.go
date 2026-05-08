package service
import ("errors"
         "alati-projekat/models")

func validacijaNazivaGroup(naziv string)error{
    if(naziv==""){
        return errors.New("ne sme naziv da bude prazan")
	}else{
		return nil
	}
}

func validacijaVerzijeGroup(verzija string)error{
	if verzija==""{
		return errors.New("ne sme naziv verzije da bude prazan")
	}else{
		return nil
	}
}

func validacijaListeKonfiguracija(listaKonfiguracija []models.Config)error{
    for _,lk:=range listaKonfiguracija{
		if(lk.Naziv==""||lk.Verzija==""){
		      return errors.New("ne postoji naziv ili ne postoji naziv vrezije")

		}

	for _,p:=range lk.Parametri{
		if(p.Kljuc==""||p.Vrednost==""){
			return errors.New("ne postoji kljuc ili ne postoji vrednost")
		}
	}

	for _,l:=range lk.ListaLabela{
		if(l.Kljuc==""||l.Vrednost==""){
			return errors.New("ne postoji kljuc ili ne postoji vrednost")
		}
	}
	

	

}
return nil
}