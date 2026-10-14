# Exercices — Structs et pointeurs

Ces exercices couvrent les structs : déclaration, initialisation, structs
imbriqués, méthodes avec receiver par valeur ou par pointeur, pointeurs vers
structs, embedding (composition), comparaison et slices de structs.

## Exercice 1 — Déclaration et initialisation

Déclare un struct `Livre` avec les champs `Titre` (string), `Auteur` (string)
et `Pages` (int). Crée une variable de ce type en utilisant :
1. le littéral positionnel (`Livre{"Titre", "Auteur", 300}`)
2. le littéral nommé (`Livre{Titre: "...", Auteur: "...", Pages: ...}`)

Affiche les deux valeurs avec `fmt.Println`.

## Exercice 2 — Valeur zéro

Déclare une variable de type `Livre` sans l'initialiser (`var l Livre`).
Affiche-la avec `fmt.Printf("%+v\n", l)` et explique ce que valent ses champs
par défaut.

## Exercice 3 — Struct imbriqué

Crée un struct `Editeur` avec les champs `Nom` (string) et `Pays` (string).
Ajoute un champ `Editeur` de type `Editeur` au struct `Livre` de l'exercice 1.
Initialise un `Livre` complet avec son éditeur et affiche le nom de
l'éditeur via `livre.Editeur.Nom`.

## Exercice 4 — Méthode avec receiver par valeur

Écris une méthode `Resumer() string` sur `Livre` (receiver par valeur) qui
renvoie une phrase du type `"Titre par Auteur (Pages pages)"`.
Appelle-la et affiche le résultat.

## Exercice 5 — Méthode avec receiver par pointeur

Écris une méthode `AjouterPages(n int)` sur `*Livre` (receiver par pointeur)
qui ajoute `n` au champ `Pages`. Appelle-la sur une variable `Livre` et
vérifie que `Pages` a bien changé après l'appel.

## Exercice 6 — Valeur vs pointeur, la vraie différence

Écris une méthode `RenommerParValeur(nouveauTitre string)` (receiver par
valeur) et une méthode `RenommerParPointeur(nouveauTitre string)` (receiver
par pointeur), toutes deux modifiant le champ `Titre`.
Appelle les deux sur la même variable `Livre` et affiche `Titre` après
chaque appel. Explique pourquoi l'une fonctionne et pas l'autre.

## Exercice 7 — Pointeur vers struct

Crée un `Livre` via `l := &Livre{...}` (donc `l` est un pointeur).
Modifie `l.Pages` directement (sans écrire `(*l).Pages`) et affiche le
résultat. Explique en une phrase pourquoi Go permet cette écriture simplifiée.

## Exercice 8 — Embedding (composition)

Crée un struct `LivreNumerique` qui embarque `Livre` (champ anonyme) et
ajoute un champ `TailleMo` (float64).
Initialise un `LivreNumerique`, puis affiche `Titre` et `Auteur` en y
accédant directement (sans passer par `.Livre.Titre`).
Appelle aussi la méthode `Resumer()` de l'exercice 4 directement sur le
`LivreNumerique`.

## Exercice 9 — Comparaison de structs

Crée deux variables `Livre` avec exactement les mêmes valeurs de champs, et
une troisième avec une valeur différente. Compare-les avec `==` et affiche
les résultats. Que se passerait-il si `Livre` avait un champ de type slice
(`[]string`) ? (Réponds par un commentaire, pas besoin de coder ce cas.)

## Exercice 10 — Slice de structs et le piège du range

Crée un `[]Livre` contenant au moins trois livres.
1. Essaie d'incrémenter `Pages` de chaque livre avec
   `for _, l := range livres { l.Pages++ }`, puis affiche le slice :
   constate que rien n'a changé.
2. Corrige en utilisant `for i := range livres { livres[i].Pages++ }` et
   affiche à nouveau le slice pour vérifier que les valeurs ont bien changé.
3. Explique en une phrase pourquoi la première version ne fonctionne pas.

## Exercice bonus — Bibliothèque

Écris une fonction `AjouterPagesATous(livres []Livre, n int)` qui ajoute `n`
pages à chaque livre d'un slice (en utilisant l'accès par index, pas par
copie de range). Teste-la sur un slice de plusieurs `Livre` et affiche le
résultat avant/après pour vérifier que les livres d'origine ont bien été
modifiés.
