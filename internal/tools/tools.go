// Package tools explore quelques outils standards utiles au quotidien en Go :
// manipulation de strings, conversions avec strconv, formatage avec fmt,
// gestion du temps avec time, concaténation efficace avec strings.Builder,
// et lecture de l'environnement avec os. Les tests (package testing) ont
// leur propre package dédié : tests.
package tools

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// manipulationStrings montre les fonctions les plus courantes du package
// strings : Contains, Split, Join, ToUpper/ToLower, TrimSpace, Replace.
func manipulationStrings() {
	phrase := "  Bonjour tout le monde  "

	fmt.Println(strings.Contains(phrase, "tout"))           // true
	fmt.Println(strings.Split("a,b,c", ","))                // [a b c]
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-")) // a-b-c
	fmt.Println(strings.ToUpper(phrase))
	fmt.Println(strings.ToLower(phrase))
	fmt.Printf("%q\n", strings.TrimSpace(phrase)) // enlève les espaces en début/fin
	fmt.Println(strings.Replace(phrase, "o", "0", -1))
}

// conversionsStrconv montre comment convertir entre string et types numériques,
// avec une gestion d'erreur basique à chaque fois.
func conversionsStrconv() {
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("erreur de conversion :", err)
	} else {
		fmt.Println(n + 1)
	}

	texte := strconv.Itoa(123)
	fmt.Println(texte + " est une string")

	f, err := strconv.ParseFloat("3.14", 64)
	if err != nil {
		fmt.Println("erreur de conversion :", err)
	} else {
		fmt.Println(f * 2)
	}

	fTexte := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Println(fTexte) // "3.14"

	// Exemple d'erreur : "abc" n'est pas un nombre valide.
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("conversion échouée comme prévu :", err)
	}
}

// formatageSprintf illustre les verbes les plus courants de fmt.Sprintf
// pour construire des strings formatées.
func formatageSprintf() {
	nom := "Yann"
	age := 30
	coordonnees := struct{ X, Y int }{X: 1, Y: 2}

	s1 := fmt.Sprintf("%s a %d ans", nom, age)
	s2 := fmt.Sprintf("%v", coordonnees)  // affichage simple
	s3 := fmt.Sprintf("%+v", coordonnees) // affichage avec noms de champs
	s4 := fmt.Sprintf("%T", coordonnees)  // type de la valeur
	s5 := fmt.Sprintf("%q", nom)          // string entre guillemets échappés

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}

// gestionDuTemps montre time.Now(), le formatage d'une date et le calcul
// d'une durée écoulée avec time.Since.
func gestionDuTemps() {
	maintenant := time.Now()
	fmt.Println(maintenant)

	// Go formate les dates avec une date de référence fixe : 02/01/2006 15:04:05.
	formatee := maintenant.Format("02/01/2006 15:04:05")
	fmt.Println(formatee)

	debut := time.Now()
	for i := 0; i < 1_000_000; i++ {
		_ = i * i
	}
	duree := time.Since(debut)
	fmt.Println("durée écoulée :", duree)
}

// concatenationEfficace compare la concaténation avec += (qui recopie toute
// la string à chaque fois) et strings.Builder (qui écrit dans un buffer).
func concatenationEfficace() {
	// Avec += : chaque itération recrée une nouvelle string en mémoire.
	var resultatPlusEgal string
	for i := 0; i < 5; i++ {
		resultatPlusEgal += "x"
	}
	fmt.Println(resultatPlusEgal)

	// Avec strings.Builder : écriture directe dans un buffer, plus efficace.
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		builder.WriteString("x")
	}
	fmt.Println(builder.String())
}

// lectureEnvironnement illustre la lecture d'une variable d'environnement
// avec os.Getenv et la lecture des arguments du programme avec os.Args.
func lectureEnvironnement() {
	utilisateur := os.Getenv("USER")
	if utilisateur == "" {
		utilisateur = "inconnu"
	}
	fmt.Println("utilisateur système :", utilisateur)

	fmt.Println("arguments du programme :", os.Args)
}

// Run exécute toutes les démonstrations du package.
func Run() {
	manipulationStrings()
	conversionsStrconv()
	formatageSprintf()
	gestionDuTemps()
	concatenationEfficace()
	lectureEnvironnement()
}
