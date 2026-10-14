# Programme 3 — Analyseur de texte

**Difficulté : intermédiaire**

> ⚠️ Ce programme doit être écrit **en dehors de ce repo**. Crée un nouveau dossier (ex: `~/programmes_go/03_analyseur`), initialise-le avec `go mod init analyseur`, et travaille dedans. Reviens ici seulement pour relire la consigne.

## Consigne

Écris un programme qui prend un texte (saisi, ou lu depuis un fichier `.txt` passé en argument) et affiche :

- le nombre total de mots
- le nombre de caractères (hors espaces)
- les 5 mots les plus fréquents, avec leur nombre d'occurrences
- la longueur moyenne des mots

## Une fois que ça marche : refactoriser

- As-tu une fonction qui fait "tout" le comptage, ou des fonctions séparées et testables indépendamment pour chaque statistique ?
- Comment comptes-tu les occurrences de chaque mot ? Ta structure de données est-elle adaptée (recherche en boucle dans un slice à chaque mot, ou autre chose de plus direct) ?
- Pour trouver les 5 mots les plus fréquents, as-tu dû trier quelque chose ? Comment gères-tu les ex-æquo ?
- Que se passe-t-il sur un texte vide, ou un fichier qui n'existe pas ?

## Notions que tu vas probablement découvrir en butant

- Compter des occurrences efficacement → `internal/types` (map) et `internal/data_structures`
- Découper une chaîne en mots, nettoyer la ponctuation → `internal/tools` (package strings)
- Lire le contenu d'un fichier, gérer l'erreur si le fichier n'existe pas → `internal/error_handling`
- Trier une collection selon un critère personnalisé (le package standard `sort` ou `slices`) — notion encore non couverte dans `internal/`, bon moment pour explorer la documentation officielle par toi-même
- Écrire des tests avec des cas limites (texte vide, un seul mot...) → `internal/tests`

## Pour aller plus loin (optionnel)

- Ignorer la casse et les accents pour le comptage ("Le" et "le" comptent pour le même mot).
- Permettre d'analyser plusieurs fichiers à la suite et d'agréger les statistiques.
