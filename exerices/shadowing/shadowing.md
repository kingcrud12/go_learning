# Exercices — Shadowing (masquage de variables)

Ces exercices couvrent le shadowing : le fait qu'une variable déclarée avec
`:=` dans un bloc interne (if, for, etc.) peut porter le même nom qu'une
variable d'un bloc englobant et la masquer, sans la modifier. Ne pas
confondre avec la simple portée de bloc (voir les exercices du package
`variables`) : ici le piège vient du fait qu'on réutilise le MÊME nom.

Pour chaque exercice, essaie de **prédire la valeur affichée avant
d'exécuter le code**, puis vérifie en lançant le programme.

## Exercice 1 — Shadowing basique

```go
nombre := 5

if true {
    nombre := 10
    fmt.Println("dans le bloc :", nombre)
}

fmt.Println("après le bloc :", nombre)
```

Prédis ce qui est affiché par les deux `fmt.Println`, puis exécute le code
pour vérifier. Explique en une phrase pourquoi la seconde ligne n'affiche
pas `10`.

## Exercice 2 — := vs =

Reprends le code de l'exercice 1, mais remplace `nombre := 10` par
`nombre = 10` à l'intérieur du bloc `if`. Prédis à nouveau les deux valeurs
affichées, exécute, puis explique la différence de comportement entre les
deux versions.

## Exercice 3 — Le piège du err

```go
func etape1() error {
    return nil
}

func etape2() (int, error) {
    return 42, fmt.Errorf("quelque chose s'est mal passé")
}

func main() {
    err := etape1()

    if valeur, err := etape2(); valeur > 0 {
        fmt.Println("valeur :", valeur, "err :", err)
    }

    if err != nil {
        fmt.Println("une erreur a eu lieu :", err)
    } else {
        fmt.Println("tout va bien")
    }
}
```

Avant d'exécuter : le programme va-t-il afficher "une erreur a eu lieu" ou
"tout va bien" ? Justifie ta réponse, puis exécute le code pour vérifier.
Explique où se cache le bug.

## Exercice 4 — Corriger le piège du err

Corrige le code de l'exercice 3 pour que l'erreur renvoyée par `etape2()`
soit bien prise en compte par le `if err != nil` final (indice : évite de
redéclarer `err` avec `:=` dans le second bloc `if`, ou vérifie l'erreur
immédiatement à l'intérieur de ce bloc).

## Exercice 5 — Shadowing dans une boucle

```go
statut := "initial"

for i := 0; i < 3; i++ {
    statut := fmt.Sprintf("iteration %d", i)
    fmt.Println(statut)
}

fmt.Println("valeur finale :", statut)
```

Prédis la valeur affichée par le dernier `fmt.Println`. Exécute le code et
vérifie. Qu'aurait-il fallu écrire pour que `statut` soit réellement mis à
jour à chaque itération ?

## Exercice 6 — Shadowing avec un type différent

```go
donnees := "42"

if true {
    donnees, err := strconv.Atoi(donnees)
    if err != nil {
        fmt.Println("erreur de conversion")
    }
    fmt.Println("donnees (int) :", donnees)
}

fmt.Println("donnees (string) :", donnees)
```

Prédis les deux valeurs affichées (types inclus). Exécute le code (pense à
importer `"strconv"`) et vérifie. Ce cas montre que le shadowing fonctionne
même quand le type change.

## Exercice 7 — Détecter le shadowing avec les outils

Sans exécuter de code : renseigne-toi sur l'option `-shadow` de `go vet` ou
le linter `shadow` de `golangci-lint`. En une à deux phrases, explique ce
que ces outils signalent et pourquoi ils sont utiles pour repérer des bugs
comme celui de l'exercice 3 avant qu'ils n'arrivent en production.

## Exercice 8 — Shadowing imbriqué

```go
niveau := "A"

if true {
    niveau := "B"
    if true {
        niveau := "C"
        fmt.Println("plus profond :", niveau)
    }
    fmt.Println("milieu :", niveau)
}
fmt.Println("exterieur :", niveau)
```

Prédis les trois valeurs affichées, dans l'ordre. Exécute le code et
vérifie. Combien de variables distinctes nommées `niveau` existent dans ce
programme ?

## Exercice bonus — Trouver le bug dans une fonction réaliste

```go
func recupererUtilisateur(id int) (nom string, err error) {
    nom = "inconnu"
    err = nil

    if id > 0 {
        resultat, err := rechercherEnBase(id) // attention ici
        if err != nil {
            return "", err
        }
        nom = resultat
    }

    return nom, err
}
```

En supposant que `rechercherEnBase` renvoie parfois une erreur, explique
pourquoi cette fonction peut renvoyer `err == nil` même quand
`rechercherEnBase` a échoué à l'intérieur du bloc `if`. Réécris la fonction
pour supprimer le shadowing et renvoyer correctement l'erreur.
