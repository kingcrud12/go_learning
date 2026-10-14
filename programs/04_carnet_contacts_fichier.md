# Programme 4 — Carnet de contacts persistant (fichier JSON)

**Difficulté : intermédiaire**

> ⚠️ Ce programme doit être écrit **en dehors de ce repo**. Crée un nouveau dossier (ex: `~/programmes_go/04_carnet`), initialise-le avec `go mod init carnet`, et travaille dedans. Reviens ici seulement pour relire la consigne.

## Consigne

Reprends l'idée du gestionnaire de tâches (programme 2), mais cette fois pour des contacts (nom, téléphone, email), avec une différence essentielle : **les données doivent survivre à l'arrêt du programme**, en étant sauvegardées dans un fichier JSON sur disque.

Fonctionnalités attendues :
- ajouter un contact
- lister les contacts
- rechercher un contact par nom
- supprimer un contact
- à chaque modification, réécrire le fichier ; au démarrage, charger le fichier existant s'il y en a un

## Une fois que ça marche : refactoriser

- Où se trouve la logique de lecture/écriture du fichier ? Est-elle mélangée avec la logique métier (ajouter/supprimer un contact), ou bien séparée ?
- Que se passe-t-il si le fichier JSON est corrompu ou absent au démarrage ? Ton programme plante-t-il ou démarre-t-il avec une liste vide ?
- As-tu dupliqué du code entre "sauvegarder après ajout" et "sauvegarder après suppression" ? Peux-tu factoriser ?

## Notions que tu vas probablement découvrir en butant

- Sérialiser/désérialiser une struct en JSON (package standard `encoding/json`, avec les tags de struct comme `` `json:"nom"` ``) — notion non couverte dans `internal/`, bon moment pour lire la documentation du package par toi-même
- Lire et écrire un fichier sur disque (`os.ReadFile`, `os.WriteFile`) → complète ce que tu as vu dans `internal/tools`
- Gérer l'erreur "fichier non trouvé" différemment d'une vraie erreur de lecture → `internal/error_handling` (tester un type d'erreur précis avec `errors.Is`/`os.IsNotExist`)
- Pointeurs vers struct pour modifier un contact en place dans une liste → `internal/structs_and_pointers`

## Pour aller plus loin (optionnel)

- Valider le format de l'email avant d'ajouter un contact (retourner une erreur personnalisée sinon).
- Ajouter une commande d'export vers un format texte lisible, en plus du JSON.
