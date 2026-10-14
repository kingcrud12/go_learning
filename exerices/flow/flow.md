# Exercices — Structures de contrôle (flow)

Ces exercices couvrent les structures de contrôle en Go : if/else,
switch (classique et type switch), les différentes formes de for,
break, continue, range et les labels.

## Exercice 1 — if / else if / else

Écris une fonction qui prend une température (`int`) et affiche :
- `"gel"` si elle est inférieure ou égale à 0
- `"frais"` si elle est entre 1 et 14
- `"doux"` si elle est entre 15 et 25
- `"chaud"` si elle est supérieure à 25

Teste-la avec plusieurs valeurs.

## Exercice 2 — if avec instruction d'initialisation

Écris une fonction `verifierAcces(age int)` qui utilise la forme
`if variable := expression; condition { ... }` pour calculer une variable
`majeur := age >= 18` directement dans le `if`, puis affiche
`"accès autorisé"` ou `"accès refusé"` selon sa valeur.
Vérifie que la variable créée dans le `if` n'est pas accessible après le bloc.

## Exercice 3 — switch classique avec plusieurs valeurs par case

Écris une fonction qui prend un numéro de mois (`int` de 1 à 12) et affiche
la saison correspondante grâce à un `switch`, en regroupant plusieurs mois
dans un même `case` (ex. `case 12, 1, 2:` pour l'hiver).

## Exercice 4 — switch sans expression (switch true)

Reprends l'exercice de la température (exercice 1) mais réécris-la avec un
`switch` sans expression (la forme `switch { case condition: ... }`) au lieu
du `if/else if`.

## Exercice 5 — type switch

Écris une fonction qui prend un paramètre `interface{}` et utilise un
`switch v := valeur.(type)` pour afficher un message différent selon que la
valeur est un `int`, un `string`, un `bool` ou un `float64`. Ajoute un `default`
pour les autres types. Teste-la avec un slice `[]interface{}` contenant
plusieurs types différents.

## Exercice 6 — for classique et for "while"

1. Écris une boucle `for` classique (`init; condition; post`) qui affiche
   les carrés des nombres de 1 à 10.
2. Écris ensuite une boucle `for` avec seulement une condition (le
   équivalent Go du `while`) qui divise un nombre par 2 tant qu'il est
   supérieur à 1, et affiche chaque étape.

## Exercice 7 — for infini avec break

Écris une boucle infinie (`for {}`) qui tire un compteur commençant à 0 et
l'incrémente à chaque tour. Utilise un `break` pour sortir de la boucle dès
que le compteur atteint 7. Affiche le compteur à chaque itération.

## Exercice 8 — continue

Écris une boucle de 1 à 20 qui affiche uniquement les multiples de 3, en
utilisant `continue` pour sauter les nombres qui ne sont pas des multiples
de 3.

## Exercice 9 — range sur un slice et sur une map

1. Déclare un slice de `string` contenant au moins 4 prénoms et parcours-le
   avec `range` en affichant l'index et le prénom.
2. Déclare une `map[string]int` associant des noms de produits à leur prix,
   et parcours-la avec `range` en affichant chaque clé et sa valeur.

## Exercice 10 — labels avec break et continue

1. Écris deux boucles `for` imbriquées (i et j de 0 à 4) et utilise un label
   avec `break` pour sortir complètement des deux boucles dès que
   `i * j > 6`.
2. Dans une nouvelle fonction, écris deux boucles imbriquées similaires et
   utilise un label avec `continue` pour passer directement à l'itération
   suivante de la boucle externe dès que `j == 2`, sans jamais afficher
   les valeurs de `j` suivantes pour ce `i`.

## Exercice bonus — FizzBuzz

Écris une fonction `fizzBuzz(n int)` qui parcourt les nombres de 1 à `n`
avec un `for` et, pour chaque nombre, utilise un `switch` sans expression
pour afficher :
- `"FizzBuzz"` si le nombre est divisible par 3 et par 5
- `"Fizz"` si le nombre est divisible par 3 uniquement
- `"Buzz"` si le nombre est divisible par 5 uniquement
- le nombre lui-même sinon
