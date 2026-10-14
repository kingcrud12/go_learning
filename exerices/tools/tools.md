# Exercices — Outils standards (tools)

Ces exercices couvrent quelques outils standards utiles au quotidien en Go :
`strings`, `strconv`, `fmt.Sprintf`, `time`, `strings.Builder` et `os`.
Les tests unitaires (package `testing`) sont traités dans leurs propres
exercices (voir le package `tests`).

## Exercice 1 — Manipulation de strings

Déclare une string `phrase := "  Le Langage Go Est Rapide  "`.
1. Vérifie avec `strings.Contains` si elle contient le mot `"Go"`.
2. Découpe-la en mots avec `strings.Split` (attention aux espaces superflus,
   utilise d'abord `strings.TrimSpace`).
3. Rejoins les mots avec `strings.Join` en les séparant par `"_"`.
4. Affiche la version en majuscules et en minuscules.

## Exercice 2 — Remplacer dans une string

À partir de `message := "chat chat chat"`, utilise `strings.Replace` pour :
1. Ne remplacer que la **première** occurrence de `"chat"` par `"chien"`.
2. Remplacer **toutes** les occurrences de `"chat"` par `"chien"`.

Affiche les deux résultats et explique le rôle du dernier paramètre
(le nombre de remplacements).

## Exercice 3 — Conversions avec strconv

1. Convertis la string `"256"` en `int` avec `strconv.Atoi` et multiplie le
   résultat par `2`.
2. Convertis l'`int` `99` en string avec `strconv.Itoa` et concatène-le avec
   `" pourcent"`.
3. Convertis la string `"2.5"` en `float64` avec `strconv.ParseFloat`.
4. Gère l'erreur à chaque étape : si la conversion échoue, affiche un message
   clair plutôt que de laisser planter le programme.

## Exercice 4 — Conversion ratée (gestion d'erreur)

Essaie de convertir la string `"douze"` en `int` avec `strconv.Atoi`.
1. Affiche l'erreur retournée.
2. Écris une fonction `convertirOuZero(s string) int` qui renvoie la valeur
   convertie, ou `0` si la conversion échoue (sans paniquer).

## Exercice 5 — Formater un float

Utilise `strconv.FormatFloat` pour afficher la valeur `12.3456789` avec :
1. 2 chiffres après la virgule.
2. 4 chiffres après la virgule.

Indice : regarde la signature de `strconv.FormatFloat(f float64, fmt byte,
prec int, bitSize int) string`.

## Exercice 6 — fmt.Sprintf et ses verbes

Déclare une struct `Produit` avec les champs `Nom string` et `Prix float64`.
Crée une instance, puis utilise `fmt.Sprintf` pour construire et afficher :
1. Une phrase du type `"Le produit X coûte Y euros"` avec `%s` et `%v`.
2. La struct complète avec `%v` puis avec `%+v` (compare les deux affichages).
3. Le type de la struct avec `%T`.
4. Le nom du produit entre guillemets échappés avec `%q`.

## Exercice 7 — Manipuler le temps

1. Récupère l'instant présent avec `time.Now()` et affiche-le.
2. Formate cet instant au format `"2006-01-02"` (année-mois-jour uniquement).
3. Mesure le temps d'exécution d'une boucle qui tourne de `0` à `500000` en
   utilisant `time.Since`. Affiche la durée obtenue.

## Exercice 8 — strings.Builder vs +=

1. Écris une fonction qui construit une string de 10 000 caractères `"a"` en
   utilisant l'opérateur `+=` dans une boucle.
2. Écris une seconde fonction qui fait la même chose avec un
   `strings.Builder`.
3. Mesure et compare la durée d'exécution des deux fonctions avec
   `time.Since`. Que remarques-tu, et pourquoi `strings.Builder` est-il plus
   efficace ?

## Exercice 9 — Lire l'environnement

1. Lis la variable d'environnement `"HOME"` avec `os.Getenv` et affiche-la.
2. Si la variable n'existe pas (string vide), affiche une valeur par défaut
   de ton choix au lieu d'une chaîne vide.
3. Affiche le contenu de `os.Args` et explique ce qu'il contient lorsqu'on
   exécute le programme avec `go run main.go arg1 arg2`.

## Exercice bonus — Mini-formateur de rapport

Écris une fonction `genererRapport(nom string, score float64) string` qui :
1. Utilise `strings.Builder` pour construire un rapport multi-lignes.
2. Inclut le nom, le score formaté à 2 décimales (`strconv.FormatFloat`),
   et la date/heure actuelle formatée en `"02/01/2006"`.
3. Renvoie la string complète et affiche-la avec `fmt.Println`.
