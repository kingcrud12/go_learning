// Package error_handling explore la gestion des erreurs en Go : le pattern
// classique valeur/erreur, la création d'erreurs simples ou formatées, les
// erreurs personnalisées, le wrapping/unwrapping, errors.Is, errors.As, ainsi
// que panic et recover pour les cas exceptionnels.
package error_handling

import (
	"errors"
	"fmt"
)

// diviser illustre le pattern classique résultat, err := f(); if err != nil { ... }.
func diviser(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division par zéro")
	}
	return a / b, nil
}

// patternClassique montre comment on vérifie systématiquement l'erreur
// retournée avant d'utiliser le résultat.
func patternClassique() {
	resultat, err := diviser(10, 2)
	if err != nil {
		fmt.Println("erreur :", err)
	} else {
		fmt.Println("résultat :", resultat)
	}

	resultat, err = diviser(10, 0)
	if err != nil {
		fmt.Println("erreur :", err)
	} else {
		fmt.Println("résultat :", resultat)
	}
}

// erreurSimple montre errors.New pour créer une erreur avec un simple message.
func erreurSimple() {
	err := errors.New("quelque chose s'est mal passé")
	fmt.Println(err)
}

// erreurFormatee montre fmt.Errorf pour construire un message dynamique,
// à la manière de fmt.Sprintf, sans wrapping d'une autre erreur.
func erreurFormatee() {
	utilisateur := "Yann"
	err := fmt.Errorf("utilisateur %q non autorisé", utilisateur)
	fmt.Println(err)
}

// ErreurValidation est une erreur personnalisée : une struct qui implémente
// l'interface error en définissant la méthode Error() string.
type ErreurValidation struct {
	Champ   string
	Message string
}

// Error implémente l'interface error pour ErreurValidation.
func (e *ErreurValidation) Error() string {
	return fmt.Sprintf("validation du champ %q : %s", e.Champ, e.Message)
}

// validerAge retourne une *ErreurValidation si l'âge est invalide.
func validerAge(age int) error {
	if age < 0 {
		return &ErreurValidation{Champ: "age", Message: "ne peut pas être négatif"}
	}
	return nil
}

// erreurPersonnalisee montre l'utilisation d'un type d'erreur qui transporte
// plus d'informations qu'un simple message.
func erreurPersonnalisee() {
	err := validerAge(-5)
	if err != nil {
		fmt.Println(err)
	}
}

// ErrNotFound est une erreur sentinelle : une valeur d'erreur prédéfinie,
// comparable avec errors.Is, qui représente un cas précis et réutilisable.
var ErrNotFound = errors.New("ressource introuvable")

// chercherUtilisateur simule une recherche qui échoue et wrappe l'erreur
// sentinelle ErrNotFound avec du contexte supplémentaire grâce à %w.
func chercherUtilisateur(id int) error {
	if id != 1 {
		return fmt.Errorf("recherche de l'utilisateur %d : %w", id, ErrNotFound)
	}
	return nil
}

// wrapperEtDewrapper montre fmt.Errorf("...: %w", err) pour envelopper une
// erreur tout en conservant l'erreur d'origine accessible via errors.Unwrap.
func wrapperEtDewrapper() {
	err := chercherUtilisateur(42)
	fmt.Println(err) // message complet avec le contexte ajouté

	erreurOriginale := errors.Unwrap(err)
	fmt.Println(erreurOriginale) // erreur d'origine, sans le contexte
}

// comparerAvecErrorsIs montre errors.Is pour vérifier si une erreur (ou l'une
// des erreurs qu'elle wrappe) correspond à une erreur sentinelle précise.
func comparerAvecErrorsIs() {
	err := chercherUtilisateur(42)

	if errors.Is(err, ErrNotFound) {
		fmt.Println("l'utilisateur n'existe pas")
	} else {
		fmt.Println("autre type d'erreur")
	}
}

// recupererAvecErrorsAs montre errors.As pour extraire un type d'erreur
// personnalisé précis depuis une chaîne d'erreurs wrappées, même si
// l'erreur a été enveloppée plusieurs fois.
func recupererAvecErrorsAs() {
	errValidation := validerAge(-5)
	err := fmt.Errorf("échec de la création du profil : %w", errValidation)

	var cible *ErreurValidation
	if errors.As(err, &cible) {
		fmt.Printf("erreur de validation détectée sur le champ %q\n", cible.Champ)
	} else {
		fmt.Println("ce n'est pas une erreur de validation")
	}
}

// diviserOuPaniquer déclenche un panic si b vaut 0, simulant une situation
// jugée irrécupérable localement.
func diviserOuPaniquer(a, b int) int {
	if b == 0 {
		panic("division par zéro impossible")
	}
	return a / b
}

// panicEtRecover montre comment recover(), appelé dans un defer, intercepte
// un panic et empêche le crash du programme.
func panicEtRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover a intercepté un panic :", r)
		}
	}()

	fmt.Println(diviserOuPaniquer(10, 2))
	fmt.Println(diviserOuPaniquer(10, 0)) // déclenche le panic
	fmt.Println("cette ligne n'est jamais exécutée")
}

// Run exécute toutes les démonstrations du package.
func Run() {
	patternClassique()
	erreurSimple()
	erreurFormatee()
	erreurPersonnalisee()
	wrapperEtDewrapper()
	comparerAvecErrorsIs()
	recupererAvecErrorsAs()
	panicEtRecover()
}
