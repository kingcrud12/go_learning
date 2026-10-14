// Package interfaces explore les interfaces en Go : déclaration, implémentation
// implicite (duck typing), polymorphisme, interface vide (any), assertion de
// type, type switch et composition d'interfaces par embedding. L'interface
// error est volontairement traitée en bref ici, car le package error_handling
// lui est entièrement dédié.
package interfaces

import "fmt"

// Forme est une interface simple : tout type qui possède une méthode
// Aire() float64 implémente automatiquement Forme, sans mot-clé "implements".
type Forme interface {
	Aire() float64
}

// Cercle représente un cercle défini par son rayon.
type Cercle struct {
	Rayon float64
}

// Aire calcule l'aire du cercle. Cette seule méthode suffit à faire de
// Cercle un Forme : c'est le duck typing, l'implémentation est implicite.
func (c Cercle) Aire() float64 {
	return 3.14159 * c.Rayon * c.Rayon
}

// Rectangle représente un rectangle défini par sa largeur et sa hauteur.
type Rectangle struct {
	Largeur float64
	Hauteur float64
}

// Aire calcule l'aire du rectangle. Rectangle implémente Forme de la même
// façon que Cercle, sans aucun lien déclaré entre les deux types.
func (r Rectangle) Aire() float64 {
	return r.Largeur * r.Hauteur
}

// declarationEtImplementation montre que Cercle et Rectangle, bien que des
// types totalement indépendants, satisfont tous les deux l'interface Forme.
func declarationEtImplementation() {
	var f1 Forme = Cercle{Rayon: 2}
	var f2 Forme = Rectangle{Largeur: 3, Hauteur: 4}

	fmt.Printf("aire cercle=%.2f aire rectangle=%.2f\n", f1.Aire(), f2.Aire())
}

// afficherAire est une fonction polymorphique : elle accepte n'importe quel
// type qui implémente Forme, sans savoir à l'avance s'il s'agit d'un Cercle,
// d'un Rectangle ou d'un futur type qui n'existe pas encore.
func afficherAire(f Forme) {
	fmt.Printf("l'aire de cette forme est %.2f\n", f.Aire())
}

// polymorphisme illustre l'appel de afficherAire avec des types concrets
// différents : le même code fonctionne pour toutes les formes.
func polymorphisme() {
	formes := []Forme{
		Cercle{Rayon: 1},
		Rectangle{Largeur: 2, Hauteur: 5},
		Cercle{Rayon: 3},
	}

	for _, f := range formes {
		afficherAire(f)
	}
}

// interfaceVide montre que `any` (alias de interface{}) peut contenir une
// valeur de n'importe quel type, ce qui la rend utile pour des conteneurs
// génériques ou des fonctions acceptant tout type de valeur.
func interfaceVide() {
	var quelconque any

	quelconque = 42
	fmt.Println(quelconque)

	quelconque = "un texte"
	fmt.Println(quelconque)

	quelconque = Cercle{Rayon: 5}
	fmt.Println(quelconque)

	valeurs := []any{1, "deux", 3.0, true, Rectangle{Largeur: 1, Hauteur: 1}}
	fmt.Println(valeurs)
}

// assertionDeType montre v, ok := i.(TypeConcret) pour récupérer la valeur
// concrète stockée dans une interface, sans paniquer si le type ne correspond
// pas grâce à la forme à deux valeurs.
func assertionDeType() {
	var f Forme = Cercle{Rayon: 4}

	cercle, ok := f.(Cercle)
	if ok {
		fmt.Printf("c'est bien un cercle de rayon %.1f\n", cercle.Rayon)
	}

	rectangle, ok := f.(Rectangle)
	if !ok {
		fmt.Println("ce n'est pas un rectangle, valeur zéro obtenue :", rectangle)
	}
}

// typeSwitch montre switch v := i.(type) pour exécuter un traitement
// différent selon le type concret réellement stocké dans l'interface.
func typeSwitch() {
	formes := []Forme{
		Cercle{Rayon: 2},
		Rectangle{Largeur: 3, Hauteur: 3},
	}

	for _, f := range formes {
		switch v := f.(type) {
		case Cercle:
			fmt.Printf("cercle de rayon %.1f, aire %.2f\n", v.Rayon, v.Aire())
		case Rectangle:
			fmt.Printf("rectangle %gx%g, aire %.2f\n", v.Largeur, v.Hauteur, v.Aire())
		default:
			fmt.Println("forme inconnue")
		}
	}
}

// Lecteur sait lire des données sous forme de texte.
type Lecteur interface {
	Lire() string
}

// Ecrivain sait écrire des données sous forme de texte.
type Ecrivain interface {
	Ecrire(donnees string)
}

// LecteurEcrivain compose Lecteur et Ecrivain par embedding : tout type qui
// implémente les deux méthodes Lire() et Ecrire() satisfait automatiquement
// cette interface composée, sans rien déclarer de plus.
type LecteurEcrivain interface {
	Lecteur
	Ecrivain
}

// Memoire est un tampon de texte simple qui implémente LecteurEcrivain.
type Memoire struct {
	contenu string
}

// Lire renvoie le contenu actuellement stocké.
func (m *Memoire) Lire() string {
	return m.contenu
}

// Ecrire remplace le contenu stocké.
func (m *Memoire) Ecrire(donnees string) {
	m.contenu = donnees
}

// interfaceComposee montre qu'une Memoire peut être utilisée partout où l'on
// attend un Lecteur, un Ecrivain, ou l'interface composée LecteurEcrivain.
func interfaceComposee() {
	var rw LecteurEcrivain = &Memoire{}

	rw.Ecrire("première version")
	fmt.Println(rw.Lire())

	rw.Ecrire("version mise à jour")
	fmt.Println(rw.Lire())
}

// rappelInterfaceError rappelle brièvement que error est elle-même une
// interface standard, définie par une seule méthode Error() string. Tout type
// qui implémente cette méthode peut être utilisé comme une erreur. Le package
// error_handling détaille ce sujet en profondeur.
func rappelInterfaceError() {
	var err error = fmt.Errorf("exemple d'erreur : %s", "panne réseau")
	fmt.Println(err)
}

// Run exécute toutes les démonstrations du package.
func Run() {
	declarationEtImplementation()
	polymorphisme()
	interfaceVide()
	assertionDeType()
	typeSwitch()
	interfaceComposee()
	rappelInterfaceError()
}
