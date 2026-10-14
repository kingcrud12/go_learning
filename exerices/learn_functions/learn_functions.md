# Exercices — Fonctions

Ces exercices couvrent les fonctions en Go : paramètres et retour simple,
retours multiples, retour nommé, variadique, fonction en tant que valeur,
higher-order function, closure, fonction anonyme, récursivité et defer.

## Exercice 1 — Fonction simple

Écris une fonction `soustraire(a int, b int) int` qui renvoie `a - b`.
Appelle-la avec deux valeurs de ton choix et affiche le résultat.

## Exercice 2 — Plusieurs valeurs de retour

Écris une fonction `minEtMax(nombres []int) (int, int)` qui renvoie le plus
petit et le plus grand nombre du slice. Teste-la sur un slice d'au moins
cinq valeurs et affiche les deux résultats.

## Exercice 3 — Retour avec erreur

Écris une fonction `racineCarree(n float64) (float64, error)` qui renvoie
une erreur si `n` est négatif (pas de racine carrée réelle), sinon la racine
carrée de `n` (tu peux utiliser `math.Sqrt`). Teste-la avec une valeur
positive puis une valeur négative, et affiche le résultat ou l'erreur selon
le cas.

## Exercice 4 — Retour nommé

Écris une fonction `statistiques(notes []float64) (moyenne float64, total float64)`
qui utilise des retours nommés pour calculer la somme et la moyenne d'un
slice de notes. Affiche les deux valeurs après l'appel.

## Exercice 5 — Paramètre variadique

Écris une fonction `concatener(separateur string, mots ...string) string`
qui assemble un nombre variable de chaînes en les séparant par
`separateur`. Appelle-la avec zéro, un puis plusieurs mots, et une dernière
fois en passant un slice existant avec `...`.

## Exercice 6 — Fonction comme valeur

Déclare deux fonctions `estPair(n int) bool` et `estImpair(n int) bool`.
Assigne l'une des deux à une variable `filtre`, puis utilise cette variable
pour tester plusieurs nombres et afficher le résultat.

## Exercice 7 — Higher-order function

Écris une fonction `filtrer(nombres []int, predicat func(int) bool) []int`
qui renvoie uniquement les éléments du slice pour lesquels `predicat`
renvoie `true`. Utilise-la pour garder uniquement les nombres pairs d'un
slice, en passant une fonction anonyme en argument.

## Exercice 8 — Closure

Écris une fonction `creerAccumulateur() func(int) int` qui renvoie une
fonction capturant une variable `total`. Chaque appel de la fonction
retournée doit ajouter l'argument reçu à `total` et renvoyer la nouvelle
valeur. Crée deux accumulateurs indépendants et montre qu'ils ne
s'influencent pas l'un l'autre.

## Exercice 9 — Fonction anonyme et récursivité

1. Écris une fonction anonyme immédiatement invoquée (IIFE) qui affiche
   un message de ton choix.
2. Écris une fonction récursive `fibonacci(n int) int` qui renvoie le
   n-ième terme de la suite de Fibonacci (0, 1, 1, 2, 3, 5, 8...).
   Affiche les dix premiers termes.

## Exercice 10 — Defer

1. Écris une fonction qui place trois `defer fmt.Println(...)` à la suite
   avec des messages différents, puis affiche un message "normal" avant la
   fin de la fonction. Note l'ordre d'affichage obtenu et explique pourquoi
   (LIFO).
2. Écris une fonction `traiterFichier(nom string)` qui simule l'ouverture
   d'un fichier (un simple `fmt.Println("ouverture de", nom)`), utilise
   `defer` pour simuler sa fermeture, puis affiche un message de
   traitement entre les deux.

## Exercice bonus — Combiner closure et higher-order function

Écris une fonction `creerMultiplicateur(facteur int) func(int) int` qui
renvoie une closure multipliant son argument par `facteur`. Utilise ensuite
`appliquerATous` (ou une fonction équivalente que tu écris toi-même,
prenant un slice et une fonction `func(int) int`) pour appliquer un
multiplicateur par 10 à un slice de nombres, et affiche le résultat.
