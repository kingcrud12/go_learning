# Exercices — Interfaces

Ces exercices couvrent les interfaces en Go : déclaration, implémentation
implicite (duck typing), polymorphisme, interface vide (`any`), assertion de
type, type switch et composition d'interfaces. L'interface `error` est
abordée en profondeur dans les exercices du package `error_handling`.

## Exercice 1 — Déclarer une interface simple

Déclare une interface `Animal` avec une seule méthode `Crier() string`.
Crée un type `Chien` (struct) qui implémente `Crier()` en renvoyant `"Wouf"`.
Déclare une variable de type `Animal` contenant un `Chien` et affiche le
résultat de `Crier()`.

## Exercice 2 — Duck typing avec plusieurs types

Crée un second type `Chat` (struct) qui implémente lui aussi `Crier()`
(renvoyant `"Miaou"`), sans jamais écrire explicitement que `Chat` ou `Chien`
"implémentent" `Animal`. Vérifie que les deux types satisfont bien
l'interface `Animal` en les assignant chacun à une variable de type `Animal`.

## Exercice 3 — Fonction polymorphique

Écris une fonction `presenter(a Animal)` qui affiche
`"cet animal dit : <résultat de Crier()>"`. Crée un slice `[]Animal`
contenant un `Chien` et un `Chat`, puis boucle dessus en appelant
`presenter` sur chaque élément.

## Exercice 4 — Interface vide

Écris une fonction `decrire(v any)` qui affiche la valeur reçue ainsi que
son type grâce à `fmt.Printf` et le verbe `%T`. Appelle-la avec un `int`,
un `string`, un `bool` et une valeur de type `Chien`.

## Exercice 5 — Assertion de type

Reprends une variable de type `Animal` contenant un `Chat`. Utilise une
assertion de type à deux valeurs (`v, ok := a.(Chat)`) pour vérifier si
cette variable contient bien un `Chat`. Affiche un message différent selon
que `ok` vaut `true` ou `false`. Fais ensuite le test avec `Chien` pour
observer le cas où l'assertion échoue.

## Exercice 6 — Type switch

Écris une fonction `decrireAnimal(a Animal)` qui utilise un
`switch v := a.(type)` pour afficher un message différent selon que `v`
est un `Chien`, un `Chat`, ou un autre type (cas `default`). Teste-la sur
plusieurs animaux.

## Exercice 7 — Interface composée (embedding)

Déclare deux interfaces, `Demarrable` (méthode `Demarrer()`) et
`Arretable` (méthode `Arreter()`), puis une interface composée `Machine`
qui embarque les deux. Crée un type `Moteur` (struct) qui implémente les
deux méthodes et vérifie qu'une variable `Moteur` peut être utilisée comme
`Demarrable`, comme `Arretable`, et comme `Machine`.

## Exercice 8 — Polymorphisme avec plusieurs formes

En t'inspirant de l'interface `Forme` du package `interfaces` (méthode
`Aire() float64`), crée un nouveau type `Triangle` (avec `Base` et
`Hauteur`) qui l'implémente. Écris une fonction `surfaceTotale(formes []Forme) float64`
qui additionne l'aire de toutes les formes d'un slice, puis teste-la avec un
`Cercle`, un `Rectangle` et ton `Triangle`.

## Exercice 9 — Interface vide dans une collection hétérogène

Crée un slice `[]any` contenant au moins quatre valeurs de types différents
(par exemple un `int`, un `string`, un `bool` et une struct personnalisée).
Parcours ce slice avec un `type switch` et affiche, pour chaque élément,
un message adapté à son type réel.

## Exercice bonus — Mini-système de notifications

Déclare une interface `Notifiable` avec une méthode
`Notifier(message string) string` qui renvoie le message formaté pour le
canal concerné. Crée deux types qui l'implémentent, `NotificationEmail` et
`NotificationSMS`, chacun formatant le message différemment (par exemple
`"[Email] <message>"` et `"[SMS] <message>"`). Écris une fonction
`envoyerATous(destinataires []Notifiable, message string)` qui appelle
`Notifier` sur chaque destinataire et affiche le résultat.
