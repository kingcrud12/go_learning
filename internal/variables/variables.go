// Package variables explore tout ce qui concerne les variables en Go :
// déclaration, inférence de type, valeurs zéro, conversion, portée,
// variables multiples et pointeurs. Les constantes, iota et le shadowing
// sont volontairement exclus (voir les packages constantes et shadowing).
package variables

import "fmt"

// declarationExplicite montre la syntaxe complète : var nom type = valeur.
func declarationExplicite() {
	var age int = 30
	var prenom string = "Yann"
	var actif bool = true

	fmt.Println(age, prenom, actif)
}

// declarationInferee montre que Go peut déduire le type à partir de la valeur.
func declarationInferee() {
	var age = 30        // int
	var prenom = "Yann" // string
	var taille = 1.78   // float64

	fmt.Println(age, prenom, taille)
}

// declarationCourte montre l'opérateur := utilisable uniquement
// à l'intérieur d'une fonction.
func declarationCourte() {
	age := 30
	prenom := "Yann"
	actif := true

	fmt.Println(age, prenom, actif)
}

// declarationMultiple regroupe plusieurs variables sur une seule ligne
// ou dans un bloc var(...).
func declarationMultiple() {
	var a, b, c = 1, 2, 3
	x, y := "hello", "world"

	var (
		nom   string = "Dupont"
		age   int    = 25
		actif bool   = true
	)

	fmt.Println(a, b, c, x, y, nom, age, actif)
}

// valeursZero montre la valeur par défaut attribuée à une variable
// déclarée sans valeur initiale.
func valeursZero() {
	var i int     // 0
	var f float64 // 0
	var s string  // ""
	var b bool    // false

	fmt.Printf("int=%v float=%v string=%q bool=%v\n", i, f, s, b)
}

// conversionDeType montre qu'il n'y a pas de conversion implicite en Go :
// il faut convertir explicitement entre types numériques compatibles.
func conversionDeType() {
	var entier int = 42
	var flottant float64 = float64(entier)
	var entierDepuisFlottant int = int(flottant * 2.5)

	fmt.Println(entier, flottant, entierDepuisFlottant)
}

// portee montre qu'une variable n'existe que dans le bloc où elle est déclarée.
func portee() {
	x := "portée de fonction"

	if true {
		y := "portée du bloc if"
		fmt.Println(x, y)
	}
	// y n'est pas accessible ici

	fmt.Println(x)
}

// echangerValeurs illustre l'affectation multiple pour permuter deux variables
// sans variable temporaire.
func echangerValeurs() {
	a, b := 1, 2
	a, b = b, a

	fmt.Println(a, b)
}

// pointeurs montre comment une variable peut être référencée par son adresse
// et modifiée indirectement via un pointeur.
func pointeurs() {
	age := 30
	pointeurAge := &age // pointeurAge contient l'adresse de age

	*pointeurAge = 31 // modifie la valeur de age via le pointeur

	fmt.Println(age)          // 31
	fmt.Println(*pointeurAge) // 31
	fmt.Println(pointeurAge)  // adresse mémoire
}

// Run exécute toutes les démonstrations du package.
func Run() {
	declarationExplicite()
	declarationInferee()
	declarationCourte()
	declarationMultiple()
	valeursZero()
	conversionDeType()
	portee()
	echangerValeurs()
	pointeurs()
}
