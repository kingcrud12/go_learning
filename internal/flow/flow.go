// Package flow explore les structures de contrôle en Go : if/else,
// switch (classique et type switch), les différentes formes de for
// (classique, "while", infini), break, continue, range et les labels.
package flow

import "fmt"

// siSimple montre la syntaxe de base d'un if/else if/else.
func siSimple() {
	note := 14

	if note >= 16 {
		fmt.Println("mention très bien")
	} else if note >= 14 {
		fmt.Println("mention bien")
	} else if note >= 12 {
		fmt.Println("mention assez bien")
	} else {
		fmt.Println("pas de mention")
	}
}

// siAvecInitialisation montre qu'un if peut exécuter une instruction
// d'initialisation avant la condition, la variable créée n'étant visible
// que dans le if/else qui suit.
func siAvecInitialisation() {
	diviseur := 0

	if quotient := 10; diviseur != 0 {
		fmt.Println(quotient / diviseur)
	} else {
		fmt.Println("division par zéro évitée")
	}
	// quotient n'est pas accessible ici
}

// switchClassique montre un switch sur une valeur, avec plusieurs valeurs
// regroupées dans un même case et un case par défaut.
func switchClassique() {
	jour := "samedi"

	switch jour {
	case "samedi", "dimanche":
		fmt.Println("week-end")
	case "lundi", "mardi", "mercredi", "jeudi", "vendredi":
		fmt.Println("jour ouvré")
	default:
		fmt.Println("jour inconnu")
	}
}

// switchSansExpression montre la forme "switch true", pratique pour
// remplacer une longue chaîne de if/else if par des cases lisibles.
func switchSansExpression() {
	age := 17

	switch {
	case age < 13:
		fmt.Println("enfant")
	case age < 18:
		fmt.Println("adolescent")
	default:
		fmt.Println("adulte")
	}
}

// switchAvecType montre le type switch, qui permet d'exécuter un case
// différent selon le type dynamique contenu dans une interface{}.
func switchAvecType() {
	valeurs := []interface{}{42, "texte", 3.14, true}

	for _, v := range valeurs {
		switch valeurTypee := v.(type) {
		case int:
			fmt.Printf("entier : %d\n", valeurTypee)
		case string:
			fmt.Printf("chaîne : %q\n", valeurTypee)
		case float64:
			fmt.Printf("flottant : %g\n", valeurTypee)
		default:
			fmt.Printf("type non géré : %v\n", valeurTypee)
		}
	}
}

// forClassique montre la forme la plus connue du for : init; condition; post.
func forClassique() {
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}
}

// forCommeWhile montre que Go n'a pas de mot-clé while : un for avec
// seulement une condition joue ce rôle.
func forCommeWhile() {
	compteur := 0

	for compteur < 3 {
		fmt.Println("compteur =", compteur)
		compteur++
	}
}

// forInfiniAvecBreak montre un for sans aucune condition, qui boucle
// indéfiniment jusqu'à ce qu'un break l'interrompe.
func forInfiniAvecBreak() {
	i := 0

	for {
		if i >= 3 {
			break
		}
		fmt.Println("itération infinie, i =", i)
		i++
	}
}

// continueDansBoucle montre comment continue passe directement à
// l'itération suivante sans exécuter le reste du corps de la boucle.
func continueDansBoucle() {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue // on saute les nombres impairs
		}
		fmt.Println("pair :", i)
	}
}

// rangeSurSlice montre comment parcourir un slice avec range, en récupérant
// à la fois l'index et la valeur.
func rangeSurSlice() {
	fruits := []string{"pomme", "banane", "cerise"}

	for index, fruit := range fruits {
		fmt.Printf("fruit[%d] = %s\n", index, fruit)
	}
}

// rangeSurMap montre comment parcourir une map avec range, en récupérant
// la clé et la valeur (l'ordre de parcours n'est pas garanti en Go).
func rangeSurMap() {
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	for nom, age := range ages {
		fmt.Printf("%s a %d ans\n", nom, age)
	}
}

// labelsAvecBreak montre comment un label permet de sortir directement
// de deux boucles imbriquées avec un seul break.
func labelsAvecBreak() {
Externe:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break Externe // sort des deux boucles d'un coup
			}
			fmt.Printf("i=%d j=%d\n", i, j)
		}
	}
}

// labelsAvecContinue montre comment un label permet de passer à l'itération
// suivante de la boucle externe depuis l'intérieur de la boucle interne.
func labelsAvecContinue() {
Lignes:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue Lignes // passe directement à i+1
			}
			fmt.Printf("i=%d j=%d\n", i, j)
		}
	}
}

// Run exécute toutes les démonstrations du package.
func Run() {
	siSimple()
	siAvecInitialisation()
	switchClassique()
	switchSansExpression()
	switchAvecType()
	forClassique()
	forCommeWhile()
	forInfiniAvecBreak()
	continueDansBoucle()
	rangeSurSlice()
	rangeSurMap()
	labelsAvecBreak()
	labelsAvecContinue()
}
