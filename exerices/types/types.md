# Exercices — Types

Ces exercices couvrent les types de base et composés de Go : types numériques
et leurs limites, string et rune, bool, array, slice, map, types nommés et
conversions. Les struct et pointeurs sont traités dans le package
`structs_and_pointers`, et les generics dans leur propre package.

## Exercice 1 — Types numériques et limites

1. Déclare une variable de type `int8` et assigne-lui la valeur maximale
   autorisée en utilisant `math.MaxInt8`.
2. Fais de même avec `uint8` et `math.MaxUint8`.
3. Affiche les deux valeurs, puis ajoute `1` à la variable `int8` et affiche
   le résultat : que se passe-t-il (overflow) ?

## Exercice 2 — float32 vs float64

Déclare la même valeur décimale `1.0 / 3.0` dans une variable `float32` et
dans une variable `float64`. Affiche les deux avec `fmt.Printf("%.10f\n", ...)`
et compare le nombre de décimales correctes.

## Exercice 3 — String, byte et rune

Prends le mot `"résumé"`.
1. Affiche sa longueur avec `len(...)`.
2. Convertis-le en `[]rune` et affiche la longueur de ce slice.
3. Explique en une phrase pourquoi les deux longueurs sont différentes.

## Exercice 4 — Itération UTF-8 avec range

Utilise une boucle `for index, caractere := range "café"` pour afficher
chaque rune avec son index. Observe et note pourquoi l'index ne s'incrémente
pas toujours de 1.

## Exercice 5 — Array vs slice

1. Déclare un `array` de 4 entiers avec des valeurs de ton choix.
2. Déclare un `slice` de `string` vide, puis ajoute-y 3 villes avec `append`.
3. Affiche `len` et `cap` du slice après chaque `append` et explique la
   différence entre `len` et `cap`.

## Exercice 6 — Map : ajout et lecture sécurisée

1. Crée une map `map[string]int` représentant le stock de 3 fruits.
2. Lis la quantité d'un fruit qui existe dans la map avec le motif
   `valeur, ok := m[cle]` et affiche le résultat.
3. Fais de même avec un fruit qui n'existe pas et vérifie que `ok` vaut
   `false` et que `valeur` vaut `0`.

## Exercice 7 — Map : suppression

Reprends la map de l'exercice 6. Supprime un fruit avec `delete`, puis
affiche la map avant et après suppression pour vérifier le résultat.
Essaie aussi de supprimer une clé qui n'existe pas : que se passe-t-il ?

## Exercice 8 — Type nommé

Crée un type nommé `Euro` basé sur `float64`. Écris une fonction
`versDollars(e Euro, taux float64) float64` qui convertit un montant en
euros vers des dollars selon un taux de change donné. Teste-la avec
`10.0` euros et un taux de `1.08`.

## Exercice 9 — Conversions numériques

1. Déclare un `int` valant `7`, convertis-le en `float64` et divise-le
   par `2`.
2. Déclare un `int64` valant `500` et convertis-le en `int8`. Affiche le
   résultat et explique pourquoi la valeur affichée n'est pas `500`.

## Exercice 10 — Conversion vers/depuis string

1. Convertis l'entier `42` en `string` avec `strconv.Itoa`.
2. Convertis la chaîne `"123"` en `int` avec `strconv.Atoi` en gérant
   l'erreur retournée.
3. Essaie de convertir la chaîne `"abc"` en `int` et affiche l'erreur
   obtenue au lieu de planter le programme.

## Exercice bonus — Inventaire de magasin

Construis une map `map[string]int` représentant l'inventaire d'un magasin
(nom d'article → quantité). Écris une fonction qui :
1. Ajoute un article ou augmente sa quantité s'il existe déjà.
2. Vérifie si un article est en stock (`quantité > 0`) via le motif
   `valeur, ok := m[cle]`.
3. Supprime un article de l'inventaire quand sa quantité tombe à `0`.

Teste ta fonction avec plusieurs scénarios et affiche l'inventaire à chaque
étape.
