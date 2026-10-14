# Programmes à faire

Ce dossier contient des consignes de programmes à écrire en Go, du plus simple au plus compliqué. Ce ne sont pas des exercices isolés comme dans `exerices/` : ce sont de vrais petits programmes, à faire fonctionner puis à **refactoriser**.

## Important : où écrire le code

**Le code de ces programmes ne doit jamais être écrit dans ce repo (`go_learning`).** Ce repo est une référence de notions (`internal/`) et d'exercices ciblés (`exerices/`), pas un espace de travail pour des projets.

Pour chaque programme :

1. Crée un nouveau dossier, en dehors de `go_learning`, par exemple `~/programmes_go/01_nom_du_programme`.
2. Initialise ton module à l'intérieur :
   ```bash
   mkdir -p ~/programmes_go/01_nom_du_programme
   cd ~/programmes_go/01_nom_du_programme
   go mod init nom_du_programme
   ```
3. Écris et fais tourner ton code avec `go run .` depuis ce nouveau dossier.
4. Si tu bloques sur une notion, reviens consulter `internal/<notion>` et `exerices/<notion>` dans `go_learning`, puis retourne travailler dans ton propre dossier.

## La méthode (rappel)

Pour chaque programme :

1. **Faire fonctionner** une première version, même maladroite, sans se soucier du style.
2. **Refactoriser** : relire son propre code, repérer ce qui est répété, illisible, ou fragile.
3. **Buter volontairement** sur ce qui ne va pas, et chercher la notion qui débloque (voir `summary/2.comment_apprendre_go.md` dans `go_learning`).
4. **Vérifier** avec `gofmt -l .`, `go vet ./...`, `go build ./...`, et `go test ./...` si tu as écrit des tests.

## Liste des programmes

| N° | Fichier | Difficulté |
|----|---------|------------|
| 1 | [01_calculatrice_cli.md](01_calculatrice_cli.md) | Débutant |
| 2 | [02_gestionnaire_taches.md](02_gestionnaire_taches.md) | Débutant/Intermédiaire |
| 3 | [03_analyseur_texte.md](03_analyseur_texte.md) | Intermédiaire |
| 4 | [04_carnet_contacts_fichier.md](04_carnet_contacts_fichier.md) | Intermédiaire |
| 5 | [05_serveur_http_notes.md](05_serveur_http_notes.md) | Avancé |
| 6 | [06_crawler_concurrent.md](06_crawler_concurrent.md) | Avancé |
