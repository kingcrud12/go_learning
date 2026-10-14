# Programme 5 — Serveur HTTP de notes (API REST minimale)

**Difficulté : avancé**

> ⚠️ Ce programme doit être écrit **en dehors de ce repo**. Crée un nouveau dossier (ex: `~/programmes_go/05_api_notes`), initialise-le avec `go mod init api_notes`, et travaille dedans. Reviens ici seulement pour relire la consigne.

## Consigne

Écris un petit serveur HTTP (package standard `net/http`, pas de framework) qui expose une API pour gérer des notes (titre + contenu) :

- `GET /notes` : liste toutes les notes
- `GET /notes/{id}` : retourne une note précise
- `POST /notes` : crée une note (corps de requête en JSON)
- `DELETE /notes/{id}` : supprime une note

Les notes peuvent être stockées en mémoire (pas besoin de base de données). Teste ton serveur avec `curl` ou Postman.

## Une fois que ça marche : refactoriser

- Ton code d'accès aux données (la liste de notes en mémoire) est-il mélangé avec le code qui gère les requêtes HTTP, ou séparé dans un type dédié ?
- Que se passe-t-il si deux requêtes arrivent en même temps (par exemple deux `POST` simultanés) ? As-tu un problème d'accès concurrent à ta liste de notes ?
- Gères-tu proprement les cas d'erreur HTTP (id inexistant → 404, JSON invalide en entrée → 400) avec les bons codes de statut ?

## Notions que tu vas probablement découvrir en butant

- Le package `net/http` : routes, handlers, `http.ResponseWriter`, `*http.Request` — notion non couverte dans `internal/`, c'est l'occasion de lire la documentation officielle (`net/http`) par toi-même
- Protéger une donnée partagée accédée par plusieurs requêtes simultanées → `internal/concurrency` (sync.Mutex)
- Définir une interface pour le "stockage" des notes (pour pouvoir la remplacer plus tard par une vraie base de données sans changer le reste du code) → `internal/interfaces`
- Décoder/encoder du JSON dans les deux sens → à relier avec ce que tu as découvert au programme 4
- Tester des handlers HTTP (le package `net/http/httptest`) → prolonge `internal/tests`

## Pour aller plus loin (optionnel)

- Ajouter une route `PUT /notes/{id}` pour modifier une note existante.
- Ajouter un middleware simple qui log chaque requête (méthode, chemin, durée).
