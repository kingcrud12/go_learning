// Ce fichier illustre comment écrire des tests en Go avec le package
// "testing" de la bibliothèque standard, en testant les petites fonctions
// définies dans tests.go.
//
// Rappels pratiques :
//   - un fichier de test se nomme toujours <nom>_test.go ;
//   - une fonction de test basique a la signature func TestXxx(t *testing.T) ;
//   - pour lancer les tests : go test ./internal/tests/...
//   - pour voir la couverture de code : go test -cover ./internal/tests/...
//   - pour ne lancer que certains tests (filtre par nom/regex) :
//     go test -run TestAddition ./internal/tests/...
//   - pour lancer les benchmarks : go test -bench=. ./internal/tests/...
package tests

import (
	"errors"
	"testing"
)

// TestAddition montre le test le plus basique : on appelle la fonction,
// on compare le résultat obtenu au résultat attendu, et si ça ne correspond
// pas on signale l'échec avec t.Errorf (qui marque le test en échec mais
// laisse le test continuer, contrairement à t.Fatalf qui l'arrête net).
func TestAddition(t *testing.T) {
	resultat := Addition(2, 3)
	attendu := 5

	if resultat != attendu {
		t.Errorf("Addition(2, 3) = %d ; attendu %d", resultat, attendu)
	}
}

// TestEstPair illustre les table-driven tests : plutôt que de dupliquer
// le même test pour chaque cas, on regroupe les cas dans un slice de struct,
// puis on boucle dessus avec t.Run pour créer un sous-test nommé par cas.
// C'est le style idiomatique en Go pour tester plusieurs scénarios proprement.
func TestEstPair(t *testing.T) {
	cas := []struct {
		nom     string
		n       int
		attendu bool
	}{
		{"zero est pair", 0, true},
		{"deux est pair", 2, true},
		{"trois est impair", 3, false},
		{"nombre negatif pair", -4, true},
		{"nombre negatif impair", -7, false},
	}

	for _, c := range cas {
		// t.Run crée un sous-test indépendant, identifiable individuellement
		// dans la sortie de `go test -v` (ex: TestEstPair/deux_est_pair) et
		// filtrable avec -run "TestEstPair/deux_est_pair".
		t.Run(c.nom, func(t *testing.T) {
			resultat := EstPair(c.n)
			if resultat != c.attendu {
				t.Errorf("EstPair(%d) = %v ; attendu %v", c.n, resultat, c.attendu)
			}
		})
	}
}

// TestInverser utilise lui aussi des sous-tests avec t.Run, mais sur un
// tableau de cas plus simple, pour montrer que ce style s'applique à
// n'importe quelle fonction, pas seulement aux booléens.
func TestInverser(t *testing.T) {
	cas := []struct {
		nom     string
		entree  string
		attendu string
	}{
		{"mot simple", "bonjour", "ruojnob"},
		{"chaine vide", "", ""},
		{"un seul caractere", "a", "a"},
		{"palindrome", "radar", "radar"},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := Inverser(c.entree)
			if resultat != c.attendu {
				t.Errorf("Inverser(%q) = %q ; attendu %q", c.entree, resultat, c.attendu)
			}
		})
	}
}

// verifierErreur est une fonction utilitaire de test (un "helper"). On y
// appelle t.Helper() en première ligne : cela indique à "testing" que, en cas
// d'échec, la ligne signalée dans le rapport doit être celle de l'appelant
// (dans TestDiviser) et non celle à l'intérieur de verifierErreur elle-même.
// Cela rend les messages d'erreur beaucoup plus utiles pour localiser le bug.
func verifierErreur(t *testing.T, err error, attendueErreur bool) {
	t.Helper()

	if attendueErreur && err == nil {
		t.Error("une erreur était attendue, mais err vaut nil")
	}
	if !attendueErreur && err != nil {
		t.Errorf("aucune erreur attendue, mais err = %v", err)
	}
}

// TestDiviser illustre comment tester une fonction qui peut retourner une
// erreur, et notamment comment vérifier qu'une erreur précise est bien
// retournée dans le cas attendu (division par zéro), via errors.Is.
func TestDiviser(t *testing.T) {
	cas := []struct {
		nom             string
		a, b            float64
		resultatAttendu float64
		erreurAttendue  bool
	}{
		{"division normale", 10, 2, 5, false},
		{"division par zero", 10, 0, 0, true},
		{"numerateur nul", 0, 5, 0, false},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat, err := Diviser(c.a, c.b)

			// On délègue la vérification de l'erreur au helper : il appelle
			// t.Helper(), donc une erreur ici pointera vers cette ligne,
			// pas vers l'intérieur de verifierErreur.
			verifierErreur(t, err, c.erreurAttendue)

			if c.erreurAttendue {
				if !errors.Is(err, ErrDivisionParZero) {
					t.Errorf("err = %v ; attendu ErrDivisionParZero", err)
				}
				return
			}

			if resultat != c.resultatAttendu {
				t.Errorf("Diviser(%v, %v) = %v ; attendu %v", c.a, c.b, resultat, c.resultatAttendu)
			}
		})
	}
}

// BenchmarkAddition montre comment mesurer la performance d'une fonction.
// b.N est choisi automatiquement par le framework de test pour que la mesure
// dure suffisamment longtemps et soit fiable ; on lance les benchmarks avec :
// go test -bench=. ./internal/tests/...
// (par défaut, `go test` seul n'exécute pas les benchmarks).
func BenchmarkAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Addition(2, 3)
	}
}

// BenchmarkInverser mesure la performance d'Inverser sur une chaîne un peu
// plus longue, pour illustrer qu'on peut benchmarker n'importe quelle fonction
// avec le même procédé.
func BenchmarkInverser(b *testing.B) {
	texte := "le developpement en Go est agreable"
	for i := 0; i < b.N; i++ {
		Inverser(texte)
	}
}
