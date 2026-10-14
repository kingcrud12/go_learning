# Programme 6 — Vérificateur de liens concurrent

**Difficulté : avancé**

> ⚠️ Ce programme doit être écrit **en dehors de ce repo**. Crée un nouveau dossier (ex: `~/programmes_go/06_verif_liens`), initialise-le avec `go mod init verif_liens`, et travaille dedans. Reviens ici seulement pour relire la consigne.

## Consigne

Écris un programme qui prend une liste d'URLs (en dur dans le code, ou lue depuis un fichier texte, une URL par ligne) et vérifie, **en parallèle**, si chaque URL répond (code HTTP 200) ou non. Affiche un résumé à la fin : nombre d'URLs valides, nombre d'URLs en erreur, et le détail.

Contrainte volontaire : limite le nombre de requêtes simultanées (par exemple 5 maximum en même temps), pour ne pas tout lancer d'un coup.

## Une fois que ça marche : refactoriser

- Comment limites-tu le nombre de requêtes simultanées ? Un simple compteur partagé pose-t-il un problème de concurrence ?
- Comment collectes-tu les résultats de toutes les goroutines à la fin, sans en perdre, sans te bloquer, et sans les mélanger entre elles ?
- Que se passe-t-il si une URL met un temps anormalement long à répondre ? Ton programme attend-il indéfiniment ?

## Notions que tu vas probablement découvrir en butant

- Lancer plusieurs goroutines et attendre qu'elles terminent toutes → `internal/concurrency` (sync.WaitGroup)
- Limiter la concurrence avec un pool de workers → `internal/concurrency` (pattern worker pool)
- Faire remonter les résultats de plusieurs goroutines vers la fonction principale → `internal/concurrency` (channels)
- Mettre un timeout sur une requête HTTP (`context.WithTimeout`) — notion non couverte dans `internal/`, bon moment pour explorer le package `context` par toi-même
- Gérer l'erreur réseau proprement (timeout, URL invalide, hôte introuvable) sans faire planter tout le programme → `internal/error_handling`

## Pour aller plus loin (optionnel)

- Mesurer et afficher le temps de réponse de chaque URL.
- Réessayer automatiquement une URL en échec une ou deux fois avant de la considérer en erreur définitive.
- Écrire ce programme une première fois de façon séquentielle (sans aucune concurrence), mesurer le temps total, puis comparer avec la version concurrente : constater toi-même le gain, c'est souvent plus parlant qu'une explication.
