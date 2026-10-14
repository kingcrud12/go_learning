// Package tests explore comment tester du code en Go avec le package
// "testing" de la bibliothèque standard. Contrairement aux autres packages
// du dépôt, le contenu pédagogique ne se trouve pas dans ce fichier (via une
// fonction Run), mais dans tests_test.go : on ne peut illustrer les tests
// qu'en écrivant de vrais tests.
//
// Ce fichier ne contient que quelques petites fonctions simples, volontairement
// faciles à tester, qui serviront de "code sous test" pour tests_test.go.
package tests

import "errors"

// Addition retourne la somme de deux entiers.
func Addition(a, b int) int {
	return a + b
}

// EstPair indique si un entier est pair.
func EstPair(n int) bool {
	return n%2 == 0
}

// Inverser retourne une chaîne de caractères inversée.
func Inverser(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ErrDivisionParZero est retournée par Diviser lorsque le diviseur vaut 0.
var ErrDivisionParZero = errors.New("division par zéro")

// Diviser retourne a/b, ou une erreur si b vaut 0.
func Diviser(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionParZero
	}
	return a / b, nil
}
