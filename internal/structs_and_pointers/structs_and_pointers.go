// Package structs_and_pointers explore les structs en Go : déclaration,
// initialisation, structs imbriqués, méthodes à receiver valeur/pointeur,
// pointeurs vers structs, embedding (composition) et comparaison, ainsi que
// le piège classique des slices de structs modifiées par valeur dans un range.
package structs_and_pointers

import "fmt"

// Personne est un struct simple avec deux champs exportés.
type Personne struct {
	Nom string
	Age int
}

// declarationStruct montre la déclaration d'un type struct simple et
// l'accès à ses champs via la notation pointée.
func declarationStruct() {
	var p Personne
	p.Nom = "Yann"
	p.Age = 30

	fmt.Println(p.Nom, p.Age)
}

// initialisationLitteralePositionnelle montre l'initialisation d'un struct
// en respectant l'ordre des champs tels que déclarés dans le type.
func initialisationLitteralePositionnelle() {
	p := Personne{"Alice", 25} // l'ordre doit correspondre à Nom puis Age

	fmt.Println(p)
}

// initialisationLitteraleNommee montre l'initialisation d'un struct en
// nommant chaque champ : l'ordre n'a alors plus d'importance.
func initialisationLitteraleNommee() {
	p := Personne{Age: 40, Nom: "Bob"}

	fmt.Println(p)
}

// valeurZeroStruct montre que, sans initialisation, chaque champ d'un struct
// reçoit la valeur zéro de son propre type ("" pour string, 0 pour int).
func valeurZeroStruct() {
	var p Personne

	fmt.Printf("%+v\n", p) // {Nom: Age:0}
}

// Adresse est un struct destiné à être imbriqué dans un autre struct.
type Adresse struct {
	Rue   string
	Ville string
}

// PersonneAvecAdresse montre un struct imbriqué : le champ Adresse est
// lui-même un struct complet.
type PersonneAvecAdresse struct {
	Nom     string
	Age     int
	Adresse Adresse
}

// structImbrique montre la déclaration et l'accès à un struct contenant
// un autre struct comme champ.
func structImbrique() {
	p := PersonneAvecAdresse{
		Nom: "Claire",
		Age: 35,
		Adresse: Adresse{
			Rue:   "12 rue des Lilas",
			Ville: "Bordeaux",
		},
	}

	fmt.Println(p.Nom, "habite", p.Adresse.Rue, p.Adresse.Ville)
}

// AnniversaireParValeur reçoit une copie de Personne : incrémenter Age ici
// ne modifie pas la variable d'origine passée par l'appelant.
func (p Personne) AnniversaireParValeur() {
	p.Age++
	fmt.Println("dans la méthode (valeur):", p.Age)
}

// AnniversaireParPointeur reçoit l'adresse de la Personne d'origine :
// l'incrémentation de Age modifie bien la variable de l'appelant.
func (p *Personne) AnniversaireParPointeur() {
	p.Age++ // équivalent à (*p).Age++, Go déréférence automatiquement
	fmt.Println("dans la méthode (pointeur):", p.Age)
}

// methodeValeurVsPointeur illustre la différence essentielle entre un
// receiver par valeur (ne modifie pas l'original) et un receiver par
// pointeur (modifie l'original).
func methodeValeurVsPointeur() {
	p := Personne{Nom: "Diego", Age: 20}

	p.AnniversaireParValeur()
	fmt.Println("après receiver valeur:", p.Age) // toujours 20

	p.AnniversaireParPointeur()
	fmt.Println("après receiver pointeur:", p.Age) // 21
}

// pointeurVersStruct montre qu'on peut créer un pointeur vers un struct et
// accéder/modifier ses champs directement via ce pointeur, sans écrire
// (*p).Champ : Go effectue le déréférencement automatiquement.
func pointeurVersStruct() {
	p := &Personne{Nom: "Emma", Age: 28}

	p.Age = 29 // raccourci pour (*p).Age = 29
	fmt.Println(p.Nom, p.Age)

	nouvellePersonne := &Personne{} // pointeur vers un struct en valeur zéro
	nouvellePersonne.Nom = "Fabien"
	fmt.Println(*nouvellePersonne)
}

// Employe embarque Personne par composition (embedding) : Employe hérite
// directement des champs et méthodes de Personne, sans héritage classique.
type Employe struct {
	Personne // champ anonyme : embedding
	Poste    string
}

// embeddingStruct montre l'accès direct aux champs et méthodes du struct
// embarqué, comme s'ils appartenaient au struct englobant.
func embeddingStruct() {
	e := Employe{
		Personne: Personne{Nom: "Gaspard", Age: 45},
		Poste:    "Développeur",
	}

	// accès direct, sans passer par e.Personne.Nom
	fmt.Println(e.Nom, e.Age, e.Poste)

	e.AnniversaireParPointeur() // méthode de Personne, promue sur Employe
	fmt.Println("age après promotion:", e.Age)
}

// comparaisonStructs montre que deux structs sont comparables avec == si
// tous leurs champs le sont (ici string et int), champ par champ.
func comparaisonStructs() {
	p1 := Personne{Nom: "Hugo", Age: 22}
	p2 := Personne{Nom: "Hugo", Age: 22}
	p3 := Personne{Nom: "Hugo", Age: 23}

	fmt.Println(p1 == p2) // true, tous les champs sont égaux
	fmt.Println(p1 == p3) // false, Age diffère
}

// slicePiegeRange montre le piège classique : range copie chaque élément,
// donc modifier la variable de boucle ne modifie pas le slice d'origine.
func slicePiegeRange() {
	personnes := []Personne{
		{Nom: "Ines", Age: 30},
		{Nom: "Jules", Age: 40},
	}

	for _, p := range personnes {
		p.Age++ // modifie seulement la copie locale p, pas le slice
	}
	fmt.Println("après range par valeur (sans effet):", personnes)

	for i := range personnes {
		personnes[i].Age++ // modifie bien l'élément du slice via son index
	}
	fmt.Println("après modification par index:", personnes)
}

// Run exécute toutes les démonstrations du package.
func Run() {
	declarationStruct()
	initialisationLitteralePositionnelle()
	initialisationLitteraleNommee()
	valeurZeroStruct()
	structImbrique()
	methodeValeurVsPointeur()
	pointeurVersStruct()
	embeddingStruct()
	comparaisonStructs()
	slicePiegeRange()
}
