# Exercices — Structures de données

Ces exercices couvrent des structures de données implémentées à la main en
Go : pile (Stack), file (Queue), liste chaînée (linked list), ensemble (Set)
et matrice 2D (slice de slices). Ils vont au-delà des slices et maps de base
déjà vus ailleurs : ici, on construit de vraies structures à partir de ces
briques.

## Exercice 1 — Pile (Stack) basique

Déclare un struct `Pile` avec un champ `elements []int`. Écris les méthodes
`Push(valeur int)` et `Pop() (int, bool)` (le booléen indique si la pile
était vide). Empile 1, 2, 3 puis dépile tout en affichant chaque valeur.
Vérifie que l'ordre de sortie est bien 3, 2, 1 (LIFO).

## Exercice 2 — Pile générique sur string

Réutilise ou adapte ton struct `Pile` pour qu'il stocke des `string` au lieu
d'`int`. Empile des noms de villes et dépile-les pour vérifier le
comportement LIFO.

## Exercice 3 — File (Queue) basique

Déclare un struct `File` avec un champ `elements []int`. Écris les méthodes
`Enqueue(valeur int)` et `Dequeue() (int, bool)`. Enfile 1, 2, 3 puis défile
tout en affichant chaque valeur. Vérifie que l'ordre de sortie est bien
1, 2, 3 (FIFO), contrairement à la pile de l'exercice 1.

## Exercice 4 — Liste chaînée : insertion en tête

Déclare un struct `Noeud` avec les champs `Valeur int` et `Suivant *Noeud`.
Écris une fonction `insererEnTete(tete *Noeud, valeur int) *Noeud` qui crée
un nouveau nœud pointant vers l'ancienne tête et renvoie ce nouveau nœud.
Construis une liste de 5 éléments par insertions successives en tête, puis
parcours-la et affiche toutes les valeurs.

## Exercice 5 — Liste chaînée : recherche et longueur

Sur la liste chaînée de l'exercice 4, écris :
1. une fonction `longueur(tete *Noeud) int` qui compte le nombre de nœuds ;
2. une fonction `contient(tete *Noeud, valeur int) bool` qui parcourt la
   liste et renvoie `true` si la valeur est présente.

Teste les deux fonctions sur ta liste.

## Exercice 6 — Liste chaînée : suppression d'une valeur

Écris une fonction `supprimer(tete *Noeud, valeur int) *Noeud` qui retire le
premier nœud contenant `valeur` de la liste et renvoie la nouvelle tête
(attention au cas où c'est la tête elle-même qu'il faut supprimer). Teste en
supprimant une valeur au milieu, puis la tête.

## Exercice 7 — Ensemble (Set) avec struct{}

Déclare un ensemble de prénoms avec `make(map[string]struct{})`. Ajoute
"Alice", "Bob" et "Yann". Écris une fonction `contientSet(ensemble
map[string]struct{}, valeur string) bool` qui teste la présence d'une
valeur avec le motif `_, ok := ensemble[valeur]`.

En commentaire, explique pourquoi on utilise `struct{}` plutôt que `bool`
comme type de valeur (indice : taille en mémoire d'une valeur `struct{}`
contre une valeur `bool`).

## Exercice 8 — Ensemble : intersection de deux ensembles

Crée deux ensembles de prénoms (`map[string]struct{}`) avec quelques
prénoms en commun. Écris une fonction `intersection(a, b
map[string]struct{}) map[string]struct{}` qui renvoie un nouvel ensemble
contenant uniquement les prénoms présents dans les deux. Affiche le
résultat.

## Exercice 9 — Matrice 2D

Déclare une matrice 3x3 d'entiers avec un littéral `[][]int{...}`. Affiche
l'élément à la position [2][0]. Puis, avec deux boucles `for ... range`
imbriquées, parcours toute la matrice et affiche chaque valeur précédée de
ses coordonnées (ligne, colonne).

Ensuite, construis dynamiquement une matrice 4x4 où chaque case [i][j]
contient la valeur `i*4 + j`, en utilisant `make` pour chaque ligne.

## Bonus — Pile pour vérifier des parenthèses équilibrées

En utilisant ta `Pile` de l'exercice 1 (adaptée pour stocker des `rune` ou
des `byte`), écris une fonction `parenthesesEquilibrees(expression string)
bool` qui vérifie qu'une chaîne contenant des `(`, `)`, `[`, `]`, `{`, `}`
est bien équilibrée (chaque parenthèse ouvrante a sa fermante correspondante,
dans le bon ordre).

Teste avec `"({[]})"` (vrai) et `"({[}])"` (faux).
