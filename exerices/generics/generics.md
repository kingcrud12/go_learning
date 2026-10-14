# Exercices — Génériques (generics)

Ces exercices couvrent les types génériques en Go : fonctions paramétrées par
type, contraintes personnalisées, contrainte `comparable`, structures
génériques et le package standard `cmp`.

## Exercice 1 — Fonction générique simple

Écris une fonction générique `minimum[T cmp.Ordered](a, b T) T` qui renvoie
la plus petite des deux valeurs. Teste-la avec des `int`, des `float64` et
des `string`.

## Exercice 2 — Contrainte personnalisée

Définis une contrainte `Numerique` qui regroupe `~int` et `~float64`
(comme dans le package `generics`). Écris une fonction générique
`moyenne[T Numerique](valeurs []T) float64` qui calcule la moyenne d'un
slice de valeurs numériques. Teste-la avec un slice d'`int` et un slice de
`float64`.

## Exercice 3 — Fonction générique sur slice (Filtrer)

Écris ta propre version de `filtrer[T any](slice []T, predicate func(T) bool) []T`.
Utilise-la pour extraire, d'un slice de `string`, uniquement les mots de
plus de 3 caractères.

## Exercice 4 — Fonction générique sur slice (Map)

Écris ta propre version de `appliquer[T, U any](slice []T, transformation func(T) U) []U`.
Utilise-la pour transformer un slice de `string` en un slice d'`int`
représentant la longueur de chaque mot.

## Exercice 5 — Reduce générique

Écris une fonction générique `reduire[T, U any](slice []T, initial U, combiner func(U, T) U) U`
qui réduit un slice à une seule valeur. Utilise-la pour calculer la somme
d'un slice d'`int`, puis pour concaténer un slice de `string` en une seule
chaîne séparée par des virgules.

## Exercice 6 — Struct générique (Pile)

En t'inspirant de `Pile[T any]` du package `generics`, crée une structure
générique `File[T any]` (une file, FIFO cette fois) avec deux méthodes :
`Enfiler(valeur T)` qui ajoute à la fin, et `Defiler() (T, bool)` qui retire
et renvoie le premier élément. Teste-la avec une `File[int]` puis une
`File[string]`.

## Exercice 7 — Contrainte comparable

Écris une fonction générique `indexDe[T comparable](slice []T, element T) int`
qui renvoie l'index de la première occurrence de `element` dans `slice`,
ou `-1` si l'élément n'est pas trouvé. Teste-la sur un slice d'`int` et un
slice de `string`.

## Exercice 8 — Dédupliquer un slice générique

Écris une fonction générique `dedupliquer[T comparable](slice []T) []T` qui
renvoie un nouveau slice sans doublons, en conservant l'ordre d'apparition.
Teste-la sur `[]int{1, 2, 2, 3, 1, 4}` (résultat attendu : `[1 2 3 4]`).

## Exercice 9 — Struct générique à deux paramètres de type

Crée une structure générique `Paire[K comparable, V any]` avec deux champs
`Cle K` et `Valeur V`, et une méthode `Afficher()` qui imprime la paire au
format `clé: valeur`. Crée un slice de `Paire[string, int]` représentant
quelques âges (ex. `{"Yann", 30}`) et affiche chaque paire.

## Exercice bonus — Pile générique avec contrainte sur le sommet

Reprends la structure `Pile[T any]` du package `generics` et ajoute une
méthode `Sommet() (T, bool)` qui renvoie l'élément du sommet sans le
retirer de la pile (contrairement à `Pop`). Ajoute ensuite une fonction
libre `PileVersSlice[T any](p *Pile[T]) []T` qui renvoie une copie de tous
les éléments de la pile sous forme de slice, du fond vers le sommet.
