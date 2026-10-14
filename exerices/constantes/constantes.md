# Exercices — Constantes et iota

Ces exercices couvrent les constantes et l'identifiant spécial `iota`.

## Exercice 1 — Constante simple

Déclare une constante `TauxTVA` valant `0.20` (float64) et une constante
`NomEntreprise` valant `"Acme"` (string, type inféré).
Affiche les deux.

## Exercice 2 — Pourquoi une constante ?

Essaie de réassigner une nouvelle valeur à une constante après sa
déclaration (`TauxTVA = 0.25`). Que se passe-t-il à la compilation ?
Note l'erreur obtenue et explique en une phrase la différence entre
`const` et `var`.

## Exercice 3 — Bloc de constantes multiples

Déclare, dans un seul bloc `const (...)`, les constantes suivantes :
- `AgeMinimum` = `18`
- `AgeMaximum` = `65`
- `Devise` = `"EUR"`

Affiche les trois.

## Exercice 4 — iota basique

Crée un type nommé `Saison` (basé sur `int`) et déclare, grâce à `iota`,
les quatre saisons : `Printemps`, `Ete`, `Automne`, `Hiver`
(dans cet ordre, en commençant à 0).
Affiche la valeur numérique de chaque saison.

## Exercice 5 — iota qui ne commence pas à 0

Reprends l'exercice 4, mais fais commencer `Printemps` à `1` au lieu de `0`
(indice : tu peux utiliser une expression du type `iota + 1`).
Affiche les quatre valeurs pour vérifier qu'elles vont de 1 à 4.

## Exercice 6 — iota avec expression (bitmask / puissances)

En t'inspirant de l'exemple `Ko, Mo, Go, To` du package `constantes`,
crée tes propres constantes `Octet`, `Kilooctet`, `Megaoctet` en utilisant
`iota` et le décalage de bits (`<<`). Vérifie que `Kilooctet` vaut bien 1024.

## Exercice 7 — iota avec valeurs ignorées

Crée un bloc `const` avec `iota` où tu ignores volontairement la valeur `0`
et la valeur `2` grâce au blank identifier `_`, pour ne garder que les
constantes valant `1` et `3`. Affiche-les.

## Exercice bonus — Jours ouvrés vs week-end

En réutilisant le type `Jour` du package `constantes` (Lundi à Dimanche),
écris une fonction `estWeekEnd(j Jour) bool` qui renvoie `true` seulement
pour `Samedi` et `Dimanche`. Teste-la sur plusieurs jours et affiche le
résultat.
