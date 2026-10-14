// Package generics explore les types génériques en Go (introduits en 1.18) :
// fonctions paramétrées par type, contraintes personnalisées, contrainte
// comparable, structures génériques et utilisation du package standard cmp.
package generics

import (
	"cmp"
	"fmt"
)

// Numerique est une contrainte maison regroupant les types sur lesquels
// les opérateurs arithmétiques ont un sens. Le symbole ~ autorise aussi
// les types nommés dont le type sous-jacent est int ou float64.
type Numerique interface {
	~int | ~float64
}

// maximum est une fonction générique simple : T doit respecter cmp.Ordered,
// la contrainte standard pour tout type comparable avec < > <= >=.
func maximum[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// fonctionGeneriqueSimple illustre maximum avec différents types instanciés.
func fonctionGeneriqueSimple() {
	fmt.Println(maximum(3, 7))         // int
	fmt.Println(maximum(2.5, 1.1))     // float64
	fmt.Println(maximum("abc", "abd")) // string
}

// somme utilise la contrainte maison Numerique pour additionner deux valeurs.
func somme[T Numerique](a, b T) T {
	return a + b
}

// contrainteMaison illustre la contrainte Numerique définie plus haut.
func contrainteMaison() {
	fmt.Println(somme(2, 3))     // int
	fmt.Println(somme(2.5, 1.5)) // float64
}

// filtrer renvoie les éléments de slice qui satisfont predicate.
// any (alias de interface{}) signifie ici "aucune contrainte particulière".
func filtrer[T any](slice []T, predicate func(T) bool) []T {
	resultat := make([]T, 0, len(slice))
	for _, element := range slice {
		if predicate(element) {
			resultat = append(resultat, element)
		}
	}
	return resultat
}

// appliquer transforme chaque élément de slice via la fonction transformation,
// et peut changer de type (T -> U), par exemple []int -> []string.
func appliquer[T, U any](slice []T, transformation func(T) U) []U {
	resultat := make([]U, 0, len(slice))
	for _, element := range slice {
		resultat = append(resultat, transformation(element))
	}
	return resultat
}

// fonctionGeneriqueSlice illustre filtrer et appliquer sur un slice d'entiers.
func fonctionGeneriqueSlice() {
	nombres := []int{1, 2, 3, 4, 5, 6}

	pairs := filtrer(nombres, func(n int) bool { return n%2 == 0 })
	fmt.Println(pairs) // [2 4 6]

	carres := appliquer(nombres, func(n int) int { return n * n })
	fmt.Println(carres) // [1 4 9 16 25 36]

	motsVersLongueurs := appliquer([]string{"go", "generics"}, func(s string) int { return len(s) })
	fmt.Println(motsVersLongueurs) // [2 8]
}

// Pile est une structure générique de pile (LIFO) qui fonctionne avec
// n'importe quel type d'élément grâce au paramètre de type T.
type Pile[T any] struct {
	elements []T
}

// Push ajoute un élément au sommet de la pile.
func (p *Pile[T]) Push(valeur T) {
	p.elements = append(p.elements, valeur)
}

// Pop retire et renvoie l'élément au sommet de la pile.
// ok vaut false si la pile est vide.
func (p *Pile[T]) Pop() (valeur T, ok bool) {
	if len(p.elements) == 0 {
		var zero T
		return zero, false
	}
	dernier := len(p.elements) - 1
	valeur = p.elements[dernier]
	p.elements = p.elements[:dernier]
	return valeur, true
}

// structGenerique illustre la pile générique avec des int puis des string.
func structGenerique() {
	var pileEntiers Pile[int]
	pileEntiers.Push(1)
	pileEntiers.Push(2)
	pileEntiers.Push(3)

	for {
		valeur, ok := pileEntiers.Pop()
		if !ok {
			break
		}
		fmt.Println(valeur) // 3, 2, 1
	}

	var pileMots Pile[string]
	pileMots.Push("un")
	pileMots.Push("deux")
	valeur, _ := pileMots.Pop()
	fmt.Println(valeur) // deux
}

// contient vérifie si element est présent dans slice, grâce à la contrainte
// comparable qui garantit que l'opérateur == est utilisable sur T.
func contient[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// contrainteComparable illustre contient sur un slice d'entiers et de chaînes.
func contrainteComparable() {
	fmt.Println(contient([]int{1, 2, 3}, 2))       // true
	fmt.Println(contient([]int{1, 2, 3}, 9))       // false
	fmt.Println(contient([]string{"a", "b"}, "a")) // true
}

// paquetCmp montre l'utilisation du package standard cmp, qui fournit
// cmp.Ordered (la contrainte utilisée par maximum) et des fonctions utilitaires
// comme cmp.Compare pour comparer deux valeurs ordonnées.
func paquetCmp() {
	fmt.Println(cmp.Compare(3, 7)) // -1 (3 < 7)
	fmt.Println(cmp.Compare(7, 3)) // 1  (7 > 3)
	fmt.Println(cmp.Compare(5, 5)) // 0  (égaux)
	fmt.Println(cmp.Less(3, 7))    // true
}

// Run exécute toutes les démonstrations du package.
func Run() {
	fonctionGeneriqueSimple()
	contrainteMaison()
	fonctionGeneriqueSlice()
	structGenerique()
	contrainteComparable()
	paquetCmp()
}
