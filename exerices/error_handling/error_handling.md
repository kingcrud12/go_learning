# Exercices — Gestion des erreurs

Ces exercices couvrent la gestion des erreurs en Go : le pattern classique
valeur/erreur, errors.New, fmt.Errorf, les erreurs personnalisées, le
wrapping avec `%w`, errors.Is, errors.As, ainsi que panic et recover.

## Exercice 1 — Pattern classique

Écris une fonction `racineCarree(x float64) (float64, error)` qui retourne
une erreur (via `errors.New`) si `x` est négatif, sinon la racine carrée de
`x` (utilise `math.Sqrt`). Appelle-la avec `9` puis avec `-4` en vérifiant
systématiquement `if err != nil { ... }` avant d'afficher le résultat.

## Exercice 2 — errors.New

Écris une fonction `verifierMotDePasse(mdp string) error` qui retourne une
erreur avec `errors.New` si le mot de passe fait moins de 8 caractères,
sinon `nil`. Teste-la avec un mot de passe trop court et un mot de passe
valide.

## Exercice 3 — fmt.Errorf (sans wrapping)

Écris une fonction `verifierAge(age int) error` qui retourne, via
`fmt.Errorf`, un message incluant la valeur reçue si `age` est supérieur à
`150` (ex. `"âge invalide : 200"`). Affiche l'erreur obtenue pour `age = 200`.

## Exercice 4 — Erreur personnalisée

Crée un type `ErreurStock` (struct) avec les champs `Produit string` et
`Quantite int`, qui implémente l'interface `error` (méthode
`Error() string` renvoyant par exemple `"stock insuffisant pour {Produit}
({Quantite} restant(s))"`). Écris une fonction `verifierStock(produit string,
quantite int) error` qui retourne cette erreur si `quantite == 0`.

## Exercice 5 — Wrapping avec %w

Écris une fonction `chargerConfiguration() error` qui simule une erreur de
lecture de fichier avec `errors.New("fichier introuvable")`, puis une
fonction `demarrerApplication() error` qui appelle `chargerConfiguration` et,
en cas d'erreur, la wrappe avec `fmt.Errorf("démarrage impossible : %w",
err)`. Affiche l'erreur finale.

## Exercice 6 — errors.Unwrap

Reprends l'exercice 5 et utilise `errors.Unwrap` sur l'erreur retournée par
`demarrerApplication` pour retrouver et afficher l'erreur d'origine
(`"fichier introuvable"`), séparément du message englobant.

## Exercice 7 — Erreur sentinelle et errors.Is

Déclare une erreur sentinelle `ErrAccesRefuse = errors.New("accès refusé")`.
Écris une fonction `ouvrirFichier(utilisateurAdmin bool) error` qui wrappe
`ErrAccesRefuse` avec du contexte (`fmt.Errorf("... : %w", ErrAccesRefuse)`)
si `utilisateurAdmin` est `false`. Utilise `errors.Is` pour vérifier si
l'erreur retournée correspond à `ErrAccesRefuse` et affiche un message adapté.

## Exercice 8 — errors.As

En réutilisant le type `ErreurStock` de l'exercice 4, écris une fonction
`passerCommande(produit string, quantite int) error` qui wrappe l'erreur de
`verifierStock` avec `fmt.Errorf("commande refusée : %w", err)`. Utilise
`errors.As` pour récupérer le `*ErreurStock` sous-jacent depuis l'erreur
wrappée et affiche le nom du produit concerné.

## Exercice 9 — Panic et recover

Écris une fonction `acceder(tableau []int, index int) int` qui provoque
volontairement un panic (via un accès hors limites, ex. `tableau[index]`
avec un index trop grand). Enveloppe l'appel dans une fonction englobante
qui utilise `defer` et `recover()` pour intercepter le panic et afficher un
message d'erreur au lieu de laisser le programme planter.

## Exercice 10 — Combiner panic/recover et une valeur de retour

Transforme la fonction englobante de l'exercice 9 en
`accederEnSecurite(tableau []int, index int) (valeur int, err error)` qui
retourne une erreur (plutôt que de paniquer) lorsque l'index est invalide,
en utilisant `recover()` dans un `defer` pour convertir le panic en erreur
assignée à la variable nommée `err`. Teste-la avec un index valide et un
index invalide.

## Exercice bonus — Chaîne d'erreurs à plusieurs niveaux

Écris trois fonctions qui s'appellent en cascade (`niveau1` appelle
`niveau2` qui appelle `niveau3`), où `niveau3` retourne une erreur
sentinelle, chaque niveau intermédiaire la wrappant avec son propre
contexte via `%w`. Au niveau le plus haut, utilise à la fois `errors.Is`
pour confirmer la présence de l'erreur sentinelle d'origine et
`fmt.Println` pour afficher le message d'erreur complet, afin de montrer
que le contexte de chaque niveau est préservé malgré le wrapping successif.
