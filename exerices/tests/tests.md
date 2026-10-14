# Exercices — Tests

Ces exercices couvrent les bases du package `testing` en Go : tests
basiques, table-driven tests, sous-tests avec `t.Run`, gestion des erreurs
attendues, benchmarks, fonctions utilitaires de test et couverture de code.
Contrairement aux autres packages du dépôt, il n'y a pas de `main.go` ni de
fonction `Run()` à exécuter : le code pédagogique s'écrit directement dans
un fichier `<nom>_test.go` et s'exécute avec `go test`.

## Exercice 1 — Premier test basique

Dans le package `tests`, écris une fonction `Soustraction(a, b int) int` qui
retourne `a - b`. Puis écris `TestSoustraction` qui vérifie que
`Soustraction(5, 3)` vaut `2`, en utilisant `t.Errorf` en cas d'échec.
Lance-le avec `go test ./internal/tests/...` et vérifie que tu obtiens `ok`.

## Exercice 2 — Faire échouer un test exprès

Modifie temporairement la valeur attendue dans `TestSoustraction` pour
qu'elle soit fausse (ex. attendre `3` au lieu de `2`). Relance `go test` et
observe le message d'échec. Note bien le format du message produit par
`t.Errorf`, puis remets la bonne valeur attendue.

## Exercice 3 — Table-driven test

Écris une fonction `EstPositif(n int) bool`. Teste-la avec un
table-driven test : un slice de struct contenant plusieurs cas (`nom`,
entrée, résultat attendu), une boucle `for` qui parcourt ce slice, et un
appel à `t.Run(cas.nom, ...)` pour chaque cas. Ajoute au moins 4 cas
différents (positif, négatif, zéro, grand nombre).

## Exercice 4 — Lire la sortie verbeuse

Relance les tests de l'exercice 3 avec `go test -v ./internal/tests/...`.
Observe comment chaque sous-test apparaît nommé individuellement
(ex. `TestEstPositif/positif`). Essaie ensuite de ne lancer qu'un seul
sous-test avec `go test -run "TestEstPositif/negatif" ./internal/tests/...`.

## Exercice 5 — Tester une fonction qui retourne une erreur

Écris une fonction `Racine(x float64) (float64, error)` qui retourne la
racine carrée de `x`, ou une erreur si `x` est négatif. Écris un
table-driven test qui vérifie à la fois :
- le cas normal (résultat numérique correct) ;
- le cas où une erreur est attendue (`x` négatif), en vérifiant que
  l'erreur retournée n'est pas `nil`.

## Exercice 6 — Fonction utilitaire avec t.Helper()

Écris une fonction `verifierFlottant(t *testing.T, obtenu, attendu float64)`
qui compare deux `float64` (avec une petite tolérance, car les flottants ne
sont jamais parfaitement égaux) et appelle `t.Errorf` si besoin. Ajoute
`t.Helper()` en première ligne. Utilise cette fonction dans le test de
l'exercice 5, puis fais échouer volontairement un cas pour vérifier que le
numéro de ligne signalé est bien celui du test, pas celui du helper.

## Exercice 7 — Benchmark

Écris `BenchmarkRacine(b *testing.B)` qui appelle `Racine` en boucle
`b.N` fois. Lance-le avec `go test -bench=. -run=^$ ./internal/tests/...`
(le `-run=^$` évite de relancer tous les tests classiques en même temps).
Observe le résultat : combien de nanosecondes par opération ?

## Exercice 8 — Couverture de code

Lance `go test -cover ./internal/tests/...` et note le pourcentage de
couverture affiché. Ajoute une fonction non testée (ex. une fonction
`EstNegatif(n int) bool` sans test associé) et relance la commande :
observe comment le pourcentage diminue. Ajoute ensuite le test manquant et
vérifie que la couverture remonte.

## Exercice 9 — Sous-tests avec cas d'erreur mélangés

Écris une fonction `ConvertirEnMajuscules(s string) (string, error)` qui
retourne une erreur si la chaîne est vide, et la chaîne en majuscules sinon
(indice : package `strings`, fonction `ToUpper`). Écris un table-driven test
avec des cas mélangeant succès et erreurs attendues, et utilise `t.Run` pour
chaque cas. Assure-toi que le test couvre bien les deux branches (avec et
sans erreur).

## Exercice bonus — Filtrer et mesurer en une seule commande

En une seule commande `go test`, combine :
- le filtre `-run` pour ne lancer que les tests dont le nom contient
  `"Diviser"` ;
- l'option `-v` pour voir le détail ;
- l'option `-cover` pour voir la couverture.

Note la commande complète que tu as utilisée, et explique avec tes mots la
différence entre un test (`TestXxx`) et un benchmark (`BenchmarkXxx`).
