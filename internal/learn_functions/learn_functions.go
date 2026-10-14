// Package learn_functions explore les fonctions en Go : paramètres et retour
// simple, retours multiples, retour nommé, variadique, fonction en tant que
// valeur, higher-order function, closure, fonction anonyme, récursivité et
// defer.
package learn_functions

import "fmt"

// addition est une fonction simple : elle prend deux paramètres et renvoie
// une seule valeur.
func addition(a int, b int) int {
	return a + b
}

// diviser renvoie plusieurs valeurs : le quotient, le reste et une erreur
// éventuelle si le diviseur vaut 0.
func diviser(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("division par zéro impossible")
	}

	return a / b, a % b, nil
}

// rectangle utilise des retours nommés : perimetre et aire sont déclarés
// dans la signature et renvoyés automatiquement avec un simple "return".
func rectangle(largeur float64, hauteur float64) (perimetre float64, aire float64) {
	perimetre = 2 * (largeur + hauteur)
	aire = largeur * hauteur

	return // renvoie perimetre et aire sans les répéter
}

// somme accepte un nombre variable d'arguments grâce au paramètre variadique
// "...int". À l'intérieur de la fonction, nombres se comporte comme un []int.
func somme(nombres ...int) int {
	total := 0
	for _, n := range nombres {
		total += n
	}

	return total
}

// multiplier est une fonction classique, utilisée plus bas pour montrer
// qu'une fonction peut être assignée à une variable comme n'importe quelle
// valeur.
func multiplier(a int, b int) int {
	return a * b
}

// appliquerATous est une higher-order function : elle prend une fonction en
// paramètre et l'applique à chaque élément du slice.
func appliquerATous(nombres []int, operation func(int) int) []int {
	resultats := make([]int, len(nombres))
	for i, n := range nombres {
		resultats[i] = operation(n)
	}

	return resultats
}

// creerCompteur est une closure : elle renvoie une fonction qui capture et
// modifie la variable "compte" entre chaque appel.
func creerCompteur() func() int {
	compte := 0

	return func() int {
		compte++
		return compte
	}
}

// factorielle est une fonction récursive : elle s'appelle elle-même jusqu'à
// atteindre le cas de base (n <= 1).
func factorielle(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorielle(n-1)
}

// fonctionSimple montre l'appel d'une fonction basique avec paramètres et
// une seule valeur de retour.
func fonctionSimple() {
	resultat := addition(3, 4)

	fmt.Println(resultat)
}

// retoursMultiples montre comment récupérer plusieurs valeurs renvoyées
// par une même fonction.
func retoursMultiples() {
	quotient, reste, err := diviser(17, 5)
	if err != nil {
		fmt.Println("erreur :", err)
		return
	}

	fmt.Println(quotient, reste)

	_, _, err = diviser(10, 0)
	fmt.Println("division par zéro :", err)
}

// retourNomme montre l'appel d'une fonction utilisant des retours nommés.
func retourNomme() {
	perimetre, aire := rectangle(4, 5)

	fmt.Println(perimetre, aire)
}

// parametreVariadique montre qu'une fonction variadique peut être appelée
// avec zéro, un ou plusieurs arguments, ou avec un slice déplié via "...".
func parametreVariadique() {
	fmt.Println(somme())
	fmt.Println(somme(1, 2, 3))

	valeurs := []int{10, 20, 30, 40}
	fmt.Println(somme(valeurs...))
}

// fonctionCommeValeur montre qu'une fonction peut être assignée à une
// variable puis appelée via cette variable.
func fonctionCommeValeur() {
	operation := multiplier

	fmt.Println(operation(6, 7))
}

// fonctionEnParametre montre l'utilisation de appliquerATous avec une
// fonction nommée puis une fonction anonyme passées en argument.
func fonctionEnParametre() {
	nombres := []int{1, 2, 3, 4}

	doubles := appliquerATous(nombres, func(n int) int {
		return n * 2
	})

	fmt.Println(doubles)
}

// closure montre qu'un compteur créé par creerCompteur conserve son état
// entre les appels, car chaque closure capture sa propre variable.
func closure() {
	compteur1 := creerCompteur()
	compteur2 := creerCompteur()

	fmt.Println(compteur1()) // 1
	fmt.Println(compteur1()) // 2
	fmt.Println(compteur1()) // 3
	fmt.Println(compteur2()) // 1, indépendant de compteur1
}

// fonctionAnonyme montre une fonction sans nom, ici immédiatement invoquée
// (IIFE), puis une autre assignée à une variable avant d'être appelée.
func fonctionAnonyme() {
	func(message string) {
		fmt.Println(message)
	}("fonction anonyme immédiatement invoquée")

	carre := func(n int) int {
		return n * n
	}

	fmt.Println(carre(5))
}

// recursivite montre l'appel d'une fonction qui s'appelle elle-même pour
// calculer une factorielle.
func recursivite() {
	fmt.Println(factorielle(5)) // 5*4*3*2*1 = 120
}

// deferOrdreLifo montre que plusieurs defer s'exécutent dans l'ordre
// inverse de leur déclaration (Last In, First Out), juste avant que la
// fonction ne rende la main.
func deferOrdreLifo() {
	fmt.Println("début")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("fin")
	// ordre d'affichage : début, fin, defer 3, defer 2, defer 1
}

// ressource simule une ressource à fermer (fichier, connexion réseau...).
type ressource struct {
	nom string
}

// fermer libère la ressource. C'est le genre d'opération que l'on reporte
// classiquement avec defer juste après l'ouverture.
func (r ressource) fermer() {
	fmt.Println("fermeture de la ressource :", r.nom)
}

// ouvrirRessource simule l'ouverture d'une ressource.
func ouvrirRessource(nom string) ressource {
	fmt.Println("ouverture de la ressource :", nom)
	return ressource{nom: nom}
}

// deferCasUsage montre le cas d'usage classique du defer : garantir la
// fermeture d'une ressource quel que soit le chemin de sortie de la fonction.
func deferCasUsage() {
	r := ouvrirRessource("fichier.txt")
	defer r.fermer() // s'exécutera à la toute fin de la fonction

	fmt.Println("traitement de la ressource en cours...")
}

// Run exécute toutes les démonstrations du package.
func Run() {
	fonctionSimple()
	retoursMultiples()
	retourNomme()
	parametreVariadique()
	fonctionCommeValeur()
	fonctionEnParametre()
	closure()
	fonctionAnonyme()
	recursivite()
	deferOrdreLifo()
	deferCasUsage()
}
