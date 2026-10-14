# Exercices — Variables

Ces exercices couvrent uniquement les variables : déclaration, inférence,
valeurs zéro, conversion, portée, affectation multiple et pointeurs.
Les constantes, iota et le shadowing sont traités dans leurs propres exercices.

## Exercice 1 — Déclaration explicite

Déclare trois variables avec leur type explicite :
- un `int` nommé `annee` valant `2026`
- un `string` nommé `ville` valant `"Bordeaux"`
- un `bool` nommé `estEtudiant` valant `true`

Affiche les trois valeurs avec `fmt.Println`.

## Exercice 2 — Inférence de type

Déclare les mêmes variables que l'exercice 1, mais sans préciser le type
(laisse Go l'inférer). Utilise `fmt.Printf` avec `%T` pour afficher le type
réel de chaque variable.

## Exercice 3 — Déclaration courte

Réécris l'exercice 1 avec l'opérateur `:=`.
Que se passe-t-il si tu essaies d'utiliser `:=` en dehors d'une fonction ?
Teste-le et note l'erreur obtenue.

## Exercice 4 — Valeurs zéro

Déclare, sans les initialiser, une variable de chacun des types suivants :
`int`, `float64`, `string`, `bool`, `[]int` (slice).
Affiche chacune et explique ce qu'est une "valeur zéro" en Go.

## Exercice 5 — Conversion de type

1. Déclare un `int` valant `17`.
2. Convertis-le en `float64` et divise-le par `2`.
3. Reconvertis le résultat en `int`.
4. Affiche les trois valeurs et explique pourquoi Go n'autorise pas la
   conversion implicite entre `int` et `float64`.

## Exercice 6 — Portée (scope)

Écris une fonction qui déclare une variable `total` au niveau de la fonction,
puis un bloc `for` qui déclare une variable `i` utilisée uniquement à
l'intérieur de la boucle. Montre que `i` n'est pas accessible après le `for`.

## Exercice 7 — Déclaration multiple

Déclare trois variables (`largeur`, `hauteur`, `profondeur`) sur une seule
ligne avec des valeurs différentes, puis calcule et affiche leur produit
(le volume).

## Exercice 8 — Échanger deux valeurs

Déclare deux variables `a` et `b` avec des valeurs différentes.
Échange leurs valeurs en une seule ligne, sans variable temporaire.
Affiche `a` et `b` avant et après l'échange.

## Exercice 9 — Pointeurs

1. Déclare une variable `score` valant `10`.
2. Crée un pointeur vers `score`.
3. Modifie la valeur de `score` uniquement via le pointeur, pour qu'elle
   devienne `20`.
4. Affiche `score` et la valeur pointée pour vérifier qu'elles sont identiques.

## Exercice bonus — Fonction qui modifie une variable

Écris une fonction `doubler(n *int)` qui prend un pointeur vers un `int` et
double sa valeur directement en mémoire (sans `return`).
Appelle-la sur une variable `nombre := 5` et affiche `nombre` après l'appel
pour vérifier qu'il vaut `10`.
