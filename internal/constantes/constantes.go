// Package constantes explore les constantes en Go et l'identifiant spécial iota.
package constantes

import "fmt"

// constanteSimple montre la déclaration basique d'une constante typée et
// d'une constante dont le type est inféré.
func constanteSimple() {
	const pi float64 = 3.14159
	const bienvenue = "Bonjour"

	fmt.Println(pi, bienvenue)
}

// constantesMultiples montre un bloc const regroupant plusieurs constantes.
func constantesMultiples() {
	const (
		min = 0
		max = 100
	)

	fmt.Println(min, max)
}

// Jour est un type nommé basé sur int, utilisé pour simuler une énumération.
type Jour int

// Les jours de la semaine, générés automatiquement grâce à iota.
// iota démarre à 0 dans chaque bloc const et s'incrémente de 1 à chaque ligne.
const (
	Lundi    Jour = iota // 0
	Mardi                // 1
	Mercredi             // 2
	Jeudi                // 3
	Vendredi             // 4
	Samedi               // 5
	Dimanche             // 6
)

// iotaBasique affiche la valeur numérique attribuée par iota à chaque constante.
func iotaBasique() {
	fmt.Println(Lundi, Mardi, Mercredi, Jeudi, Vendredi, Samedi, Dimanche)
}

// Octet, Kilooctet, Megaoctet... montrent comment iota permet de calculer
// des puissances de 2 sans tout écrire à la main.
const (
	_  = iota // 0 ignoré volontairement (blank identifier)
	Ko = 1 << (10 * iota)
	Mo
	Go
	To
)

// iotaAvecExpression montre que chaque constante peut réutiliser la même
// expression que celle de la première ligne, iota changeant à chaque fois.
func iotaAvecExpression() {
	fmt.Println(Ko, Mo, Go, To)
}

// iotaAvecSaut montre qu'on peut ignorer certaines valeurs d'iota
// avec le blank identifier _.
const (
	A = iota // 0
	_        // 1, ignoré
	C        // 2
	_        // 3, ignoré
	E        // 4
)

func iotaAvecSaut() {
	fmt.Println(A, C, E)
}

// Run exécute toutes les démonstrations du package.
func Run() {
	constanteSimple()
	constantesMultiples()
	iotaBasique()
	iotaAvecExpression()
	iotaAvecSaut()
}
