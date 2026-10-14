# Programme 1 — Calculatrice en ligne de commande

**Difficulté : débutant**

> ⚠️ Ce programme doit être écrit **en dehors de ce repo**. Crée un nouveau dossier (ex: `~/programmes_go/01_calculatrice`), initialise-le avec `go mod init calculatrice`, et travaille dedans. Reviens ici seulement pour relire la consigne.

## Consigne

Écris un programme qui lit deux nombres et un opérateur (`+`, `-`, `*`, `/`) saisis par l'utilisateur, et affiche le résultat.

Exemple d'utilisation attendue :
```
Premier nombre : 10
Opérateur (+ - * /) : /
Second nombre : 0
Erreur : division par zéro impossible
```

Contraintes minimales :
- Le programme doit gérer la division par zéro sans planter.
- Le programme doit refuser un opérateur inconnu (ex: `%`) avec un message clair.
- Une seule opération par exécution suffit pour la version 1.

## Une fois que ça marche : refactoriser

Ne t'arrête pas à "ça marche". Relis ton propre code et pose-toi ces questions :

- Ta fonction `main` fait-elle tout (lecture, calcul, affichage) ? Peut-elle être découpée ?
- Que se passe-t-il si l'utilisateur tape du texte à la place d'un nombre ? Ton programme plante-t-il ou gère-t-il l'erreur proprement ?
- Si tu devais ajouter une opération (`%`, puissance...), combien d'endroits du code faudrait-il modifier ?

## Notions que tu vas probablement découvrir en butant

- Lire une entrée utilisateur (`bufio.Scanner` ou `fmt.Scan`)
- Convertir une chaîne en nombre, et gérer l'erreur de conversion → `internal/types` (strconv) et `internal/error_handling`
- Retourner plusieurs valeurs depuis une fonction (résultat + erreur) → `internal/learn_functions`
- Un `switch` pour dispatcher selon l'opérateur → `internal/flow`

## Pour aller plus loin (optionnel)

- Permettre d'enchaîner plusieurs calculs tant que l'utilisateur ne tape pas "quitter".
- Accepter l'expression complète en un seul argument de ligne de commande (`os.Args`), ex: `./calculatrice 10 / 0`.
