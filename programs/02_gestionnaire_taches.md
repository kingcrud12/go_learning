# Programme 2 — Gestionnaire de tâches en mémoire

**Difficulté : débutant / intermédiaire**

> ⚠️ Ce programme doit être écrit **en dehors de ce repo**. Crée un nouveau dossier (ex: `~/programmes_go/02_taches`), initialise-le avec `go mod init taches`, et travaille dedans. Reviens ici seulement pour relire la consigne.

## Consigne

Écris un programme en ligne de commande qui permet de gérer une liste de tâches, le temps d'une exécution (pas besoin de sauvegarder sur disque pour l'instant) :

- ajouter une tâche (un libellé)
- lister les tâches, avec leur statut (faite / pas faite)
- marquer une tâche comme faite (par son numéro)
- supprimer une tâche

Propose une boucle de menu simple tant que l'utilisateur ne choisit pas de quitter.

## Une fois que ça marche : refactoriser

- Comment représentes-tu une tâche ? Un simple `string` suffit-il, ou as-tu besoin d'un type qui regroupe le libellé et le statut ?
- Ta liste de tâches est-elle un simple slice ? Que se passe-t-il si tu supprimes une tâche au milieu de la liste — ton code gère-t-il bien le décalage des index ?
- As-tu une fonction par action (ajouter, lister, marquer, supprimer) ou tout est-il mélangé dans `main` ?

## Notions que tu vas probablement découvrir en butant

- Regrouper plusieurs champs liés → `internal/structs_and_pointers` (struct)
- Modifier un élément à l'intérieur d'un slice de structs (le piège du `range` qui copie) → `internal/structs_and_pointers` et `internal/types`
- Supprimer un élément d'un slice par index → `internal/data_structures`
- Une boucle infinie avec sortie conditionnelle → `internal/flow`

## Pour aller plus loin (optionnel)

- Ajouter une priorité ou une date d'échéance à chaque tâche, et trier l'affichage en conséquence.
- Écrire des tests (`internal/tests`) pour les fonctions d'ajout/suppression/marquage, indépendamment de la boucle de menu.
