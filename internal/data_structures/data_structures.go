// Package data_structures explore des structures de données implémentées
// à la main en Go : pile (Stack), file (Queue), liste chaînée (linked list),
// ensemble (Set) et matrice 2D (slice de slices). L'objectif est d'aller
// au-delà des slices et maps de base (voir le package types) et des structs
// simples (voir le package structs_and_pointers), en construisant de vraies
// structures de données à partir de ces briques.
package data_structures

import "fmt"

// Stack représente une pile LIFO (Last In, First Out) construite au-dessus
// d'un slice d'entiers.
type Stack struct {
	elements []int
}

// Push ajoute un élément au sommet de la pile.
func (s *Stack) Push(valeur int) {
	s.elements = append(s.elements, valeur)
}

// Pop retire et renvoie l'élément au sommet de la pile. ok vaut false si la
// pile est vide.
func (s *Stack) Pop() (valeur int, ok bool) {
	if len(s.elements) == 0 {
		return 0, false
	}

	dernierIndex := len(s.elements) - 1
	valeur = s.elements[dernierIndex]
	s.elements = s.elements[:dernierIndex] // on raccourcit le slice d'un cran

	return valeur, true
}

// Peek renvoie l'élément au sommet de la pile sans le retirer.
func (s *Stack) Peek() (valeur int, ok bool) {
	if len(s.elements) == 0 {
		return 0, false
	}

	return s.elements[len(s.elements)-1], true
}

// pileDemonstration illustre une pile : le dernier élément empilé est le
// premier dépilé (LIFO).
func pileDemonstration() {
	var pile Stack

	pile.Push(1)
	pile.Push(2)
	pile.Push(3)

	sommet, _ := pile.Peek()
	fmt.Println("sommet de la pile :", sommet) // 3

	for {
		valeur, ok := pile.Pop()
		if !ok {
			break
		}
		fmt.Println("dépilé :", valeur) // 3, puis 2, puis 1
	}
}

// Queue représente une file FIFO (First In, First Out) construite au-dessus
// d'un slice d'entiers.
type Queue struct {
	elements []int
}

// Enqueue ajoute un élément à la fin de la file.
func (q *Queue) Enqueue(valeur int) {
	q.elements = append(q.elements, valeur)
}

// Dequeue retire et renvoie l'élément en tête de la file. ok vaut false si
// la file est vide.
func (q *Queue) Dequeue() (valeur int, ok bool) {
	if len(q.elements) == 0 {
		return 0, false
	}

	valeur = q.elements[0]
	q.elements = q.elements[1:] // on avance le début du slice d'un cran

	return valeur, true
}

// fileDemonstration illustre une file : le premier élément enfilé est le
// premier défilé (FIFO).
func fileDemonstration() {
	var file Queue

	file.Enqueue(1)
	file.Enqueue(2)
	file.Enqueue(3)

	for {
		valeur, ok := file.Dequeue()
		if !ok {
			break
		}
		fmt.Println("défilé :", valeur) // 1, puis 2, puis 3
	}
}

// Noeud représente un maillon d'une liste chaînée simple : une valeur et un
// pointeur vers le maillon suivant (nil pour le dernier maillon).
type Noeud struct {
	Valeur  int
	Suivant *Noeud
}

// insererEnTete ajoute une nouvelle valeur au début d'une liste chaînée et
// renvoie le nouveau premier maillon (la nouvelle tête).
func insererEnTete(tete *Noeud, valeur int) *Noeud {
	nouveauNoeud := &Noeud{Valeur: valeur, Suivant: tete}
	return nouveauNoeud
}

// parcourirListe affiche toutes les valeurs d'une liste chaînée, du premier
// au dernier maillon.
func parcourirListe(tete *Noeud) {
	for noeudCourant := tete; noeudCourant != nil; noeudCourant = noeudCourant.Suivant {
		fmt.Print(noeudCourant.Valeur, " ")
	}
	fmt.Println()
}

// listeChaineeDemonstration illustre la construction d'une liste chaînée par
// insertions successives en tête, puis son parcours.
func listeChaineeDemonstration() {
	var tete *Noeud // liste vide : tete est nil

	tete = insererEnTete(tete, 3) // liste : 3
	tete = insererEnTete(tete, 2) // liste : 2 -> 3
	tete = insererEnTete(tete, 1) // liste : 1 -> 2 -> 3

	fmt.Print("liste chaînée : ")
	parcourirListe(tete)
}

// ensembleDemonstration illustre un ensemble (Set) construit avec une map.
// On utilise map[string]struct{} plutôt que map[string]bool : struct{} est un
// type de taille zéro (zero-size type), il n'occupe aucun octet en mémoire
// pour la valeur, contrairement à bool qui occupe un octet par entrée. Seule
// la présence de la clé compte, pas la valeur associée.
func ensembleDemonstration() {
	villesVisitees := make(map[string]struct{})

	villesVisitees["Bordeaux"] = struct{}{}
	villesVisitees["Paris"] = struct{}{}
	villesVisitees["Lyon"] = struct{}{}

	_, dejaVisitee := villesVisitees["Paris"]
	fmt.Println("Paris déjà visitée ?", dejaVisitee) // true

	_, dejaVisitee = villesVisitees["Marseille"]
	fmt.Println("Marseille déjà visitée ?", dejaVisitee) // false

	delete(villesVisitees, "Lyon") // retirer un élément de l'ensemble
	fmt.Println("nombre de villes visitées :", len(villesVisitees))
}

// matriceDemonstration illustre une matrice 2D construite comme un slice de
// slices : chaque ligne est elle-même un slice indépendant.
func matriceDemonstration() {
	matrice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("élément [1][2] :", matrice[1][2]) // 6

	for indexLigne, ligne := range matrice {
		for indexColonne, valeur := range ligne {
			fmt.Printf("matrice[%d][%d]=%d ", indexLigne, indexColonne, valeur)
		}
	}
	fmt.Println()

	// Construction dynamique d'une matrice vide puis remplissage ligne par
	// ligne : chaque ligne doit être allouée séparément avec make.
	lignes, colonnes := 2, 3
	grille := make([][]int, lignes)
	for i := range grille {
		grille[i] = make([]int, colonnes)
		for j := range grille[i] {
			grille[i][j] = i*colonnes + j
		}
	}
	fmt.Println("grille construite dynamiquement :", grille)
}

// Run exécute toutes les démonstrations du package.
func Run() {
	pileDemonstration()
	fileDemonstration()
	listeChaineeDemonstration()
	ensembleDemonstration()
	matriceDemonstration()
}
