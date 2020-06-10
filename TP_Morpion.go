package main

import "github.com/tadvi/winc"

var mainWindow = winc.NewForm(nil)

var btnCase1 = winc.NewPushButton(mainWindow)
var btnCase2 = winc.NewPushButton(mainWindow)
var btnCase3 = winc.NewPushButton(mainWindow)
var btnCase4 = winc.NewPushButton(mainWindow)
var btnCase5 = winc.NewPushButton(mainWindow)
var btnCase6 = winc.NewPushButton(mainWindow)
var btnCase7 = winc.NewPushButton(mainWindow)
var btnCase8 = winc.NewPushButton(mainWindow)
var btnCase9 = winc.NewPushButton(mainWindow)



const taille int  = 9
const j1 string = "X"
const j2 string = "O"
var joueur string = j1

var tableauJeu = [taille]string{"","","", "","","" ,"","",""}

func main(){

    mainWindow.SetSize(600, 600)
	mainWindow.SetText("TP Morpion")
	
    btnFermer := winc.NewPushButton(mainWindow)
    btnFermer.SetText("Fermer")
    btnFermer.SetPos(250, 500)
    btnFermer.SetSize(100, 40)
	btnFermer.OnClick().Bind(wndOnClose)

	
    btnCase1.SetText("ø")
    btnCase1.SetPos(20, 20)
	btnCase1.SetSize(100, 100)
	btnCase1.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[0] != "X" && tableauJeu[0] != "O"{
			Changement(1, mainWindow)
		}
	})
	
    btnCase2.SetText("ø")
    btnCase2.SetPos(120, 20)
	btnCase2.SetSize(100, 100)
	btnCase2.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[1] != "X" && tableauJeu[1] != "O"{
			Changement(2, mainWindow)
		}
	})
	
    btnCase3.SetText("ø")
    btnCase3.SetPos(220, 20)
	btnCase3.SetSize(100, 100)
	btnCase3.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[2] != "X" && tableauJeu[2] != "O"{
			Changement(3, mainWindow)
		}
	})
	
    btnCase4.SetText("ø")
    btnCase4.SetPos(20, 120)
	btnCase4.SetSize(100, 100)
	btnCase4.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[3] != "X" && tableauJeu[3] != "O"{
			Changement(4, mainWindow)
		}
	})
	
    btnCase5.SetText("ø")
    btnCase5.SetPos(120, 120)
	btnCase5.SetSize(100, 100)
	btnCase5.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[4] != "X" && tableauJeu[4] != "O"{
			Changement(5, mainWindow)
		}
	})
	
    btnCase6.SetText("ø")
    btnCase6.SetPos(220, 120)
	btnCase6.SetSize(100, 100)
	btnCase6.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[5] != "X" && tableauJeu[5] != "O"{
			Changement(6, mainWindow)
		}
	})

    btnCase7.SetText("ø")
    btnCase7.SetPos(20, 220)
	btnCase7.SetSize(100, 100)
	btnCase7.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[6] != "X" && tableauJeu[6] != "O"{
			Changement(7, mainWindow)
		}
	})

    btnCase8.SetText("ø")
    btnCase8.SetPos(120, 220)
	btnCase8.SetSize(100, 100)
	btnCase8.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[7] != "X" && tableauJeu[7] != "O"{
			Changement(8, mainWindow)
		}
	})

    btnCase9.SetText("ø")
    btnCase9.SetPos(220, 220)
	btnCase9.SetSize(100, 100)
	btnCase9.OnClick().Bind(func(e *winc.Event){
		if tableauJeu[8] != "X" && tableauJeu[8] != "O"{
			Changement(9, mainWindow)
		}
	})

    mainWindow.Center()
    mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
    winc.RunMainLoop()
}


func wndOnClose(arg *winc.Event) {
    winc.Exit()
}

func Changement(caseNum int, mainWindow *winc.Form){

	switch caseNum{

		case 1 : 
			btnCase1.SetText(joueur)
			tableauJeu[0] = joueur
			break
		case 2 : 
			btnCase2.SetText(joueur)
			tableauJeu[1] = joueur
			break
		case 3 : 
			btnCase3.SetText(joueur)
			tableauJeu[2] = joueur
			break
		case 4 : 
			btnCase4.SetText(joueur)
			tableauJeu[3] = joueur
			break
		case 5 : 
			btnCase5.SetText(joueur)
			tableauJeu[4] = joueur
			break
		case 6 : 
			btnCase6.SetText(joueur)
			tableauJeu[5] = joueur
			break
		case 7 : 
			btnCase7.SetText(joueur)
			tableauJeu[6] = joueur
			break
		case 8 : 
			btnCase8.SetText(joueur)
			tableauJeu[7] = joueur
			break
		case 9 : 
			btnCase9.SetText(joueur)
			tableauJeu[8] = joueur
			break
	}

	var gagne = gagner()
	if gagne == true {
		gagneWindow := winc.NewForm(nil)
		gagneWindow.SetSize(200, 120)
		gagneWindow.SetText("Félicitation")

		btn := winc.NewPushButton(gagneWindow)
		btn.SetText("Bravo !! Vous avez gagné")
		btn.SetPos(10, 10)
		btn.SetSize(150, 60)
		btn.OnClick().Bind(wndOnClose)

		gagneWindow.Center()
		gagneWindow.Show()
		gagneWindow.OnClose().Bind(wndOnClose)
		winc.RunMainLoop()

	} else {
		if joueur == j1 {
			joueur = j2
		} else {
			joueur = j1
		}
	}

	
}

func gagner() bool{

	var gagne bool = false

	//Test pour voir si j1 gagne
	if tableauJeu[0] == "X" && tableauJeu[1] == "X" && tableauJeu[2] == "X"{
		gagne = true 
	} else if tableauJeu[3] == "X" && tableauJeu[4] == "X" && tableauJeu[5] == "X"{
		gagne = true
	} else if tableauJeu[6] == "X" && tableauJeu[7] == "X" && tableauJeu[8] == "X"{
		gagne = true
	} else if tableauJeu[0] == "X" && tableauJeu[3] == "X" && tableauJeu[6] == "X"{
		gagne = true
	} else if tableauJeu[1] == "X" && tableauJeu[4] == "X" && tableauJeu[7] == "X"{
		gagne = true
	} else if tableauJeu[2] == "X" && tableauJeu[5] == "X" && tableauJeu[8] == "X"{
		gagne = true
	} else if tableauJeu[0] == "X" && tableauJeu[4] == "X" && tableauJeu[8] == "X"{
		gagne = true
	} else if tableauJeu[2] == "X" && tableauJeu[4] == "X" && tableauJeu[6] == "X"{
		gagne = true
	} 
	
	//Test pour voir si j2 gagne
	if tableauJeu[0] == "O" && tableauJeu[1] == "O" && tableauJeu[2] == "O"{
        gagne = true
    } else if tableauJeu[3] == "O" && tableauJeu[4] == "O" && tableauJeu[5] == "O"{
        gagne = true
    } else if tableauJeu[6] == "O" && tableauJeu[7] == "O" && tableauJeu[8] == "O"{
        gagne = true
    } else if tableauJeu[0] == "O" && tableauJeu[3] == "O" && tableauJeu[6] == "O"{
        gagne = true
    } else if tableauJeu[1] == "O" && tableauJeu[4] == "O" && tableauJeu[7] == "X"{
        gagne = true
    } else if tableauJeu[2] == "O" && tableauJeu[5] == "O" && tableauJeu[8] == "O"{
        gagne = true
    } else if tableauJeu[0] == "O" && tableauJeu[4] == "O" && tableauJeu[8] == "O"{
        gagne = true
    } else if tableauJeu[2] == "O" && tableauJeu[4] == "O" && tableauJeu[6] == "O"{
        gagne = true
    }


	return gagne
}

func WinJ1() {
	WinWindow := winc.NewForm(nil) // VOUS N ETES PAS PRET
	WinWindow.SetSize(400, 300)
	WinWindow.SetText("Joueur 1 WIN")
	WinWindow.Show()
	WinWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()

	btn := winc.NewPushButton(WinWindow) // BOUTON FERMER
	btn.SetText("Fermer")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)
}