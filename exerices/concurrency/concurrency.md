# Exercices — Concurrency (goroutines, channels, sync)

Ces exercices couvrent les bases de la concurrence en Go : goroutines,
synchronisation avec `sync.WaitGroup` et `sync.Mutex`, channels bufferisés ou
non, `select` et un pattern simple de worker pool.

Attention : la concurrence n'a de sens que si on peut observer un résultat
stable et reproductible. N'utilise jamais `time.Sleep` comme unique moyen
d'attendre qu'une goroutine termine : utilise toujours un `sync.WaitGroup` ou
un channel bloquant.

## Exercice 1 — Première goroutine

Écris une fonction qui lance une goroutine affichant `"je suis une goroutine"`,
puis utilise un `sync.WaitGroup` pour attendre sa fin avant d'afficher
`"terminé"` dans la fonction principale. Vérifie que `"terminé"` s'affiche
toujours après le message de la goroutine, même en exécutant le programme
plusieurs fois.

## Exercice 2 — Plusieurs goroutines et WaitGroup

Lance 5 goroutines, chacune affichant son numéro (de 0 à 4). Utilise un seul
`sync.WaitGroup` pour attendre que les 5 aient terminé avant de continuer.
Remarque : l'ordre d'affichage des numéros n'est pas garanti, explique
pourquoi en une phrase.

## Exercice 3 — Channel non bufferisé

Crée un channel non bufferisé de type `string`. Dans une goroutine, envoie un
message dessus. Dans la fonction principale, reçois ce message et affiche-le.
Explique pourquoi cet envoi/réception sert aussi de synchronisation (pas
besoin de `WaitGroup` ici).

## Exercice 4 — Channel bufferisé

Crée un channel bufferisé de capacité 5 contenant des `int`. Envoie-y les
valeurs de 1 à 5 sans qu'aucun envoi ne bloque, ferme le channel, puis
parcours-le avec `range` pour afficher chaque valeur.
Que se passerait-il si tu essayais d'envoyer une 6e valeur avant toute
lecture ?

## Exercice 5 — select sur plusieurs channels

Crée deux channels non bufferisés. Lance deux goroutines, chacune envoyant un
message différent sur l'un des deux channels. Utilise un `select` qui reçoit
exactement deux messages (un par channel, dans n'importe quel ordre) et les
affiche. N'oublie pas d'attendre la fin des goroutines avec un `WaitGroup`.

## Exercice 6 — Worker pool

Crée un channel de tâches (des entiers de 1 à 10) et un channel de résultats.
Lance 3 workers (goroutines) qui lisent les tâches depuis le channel avec
`range` et écrivent dans le channel de résultats le double de chaque valeur
reçue. Utilise un `sync.WaitGroup` pour attendre la fin de tous les workers,
ferme le channel de résultats, puis calcule la somme totale des résultats.

## Exercice 7 — Race condition et mutex

1. Dans un commentaire (ne l'exécute pas), écris le code d'un compteur
   partagé incrémenté par 500 goroutines sans aucune protection, et explique
   pourquoi le résultat final est imprévisible (race condition).
2. Écris ensuite le code réel, fonctionnel, qui protège ce même compteur avec
   un `sync.Mutex` (`Lock`/`Unlock`) autour de l'incrémentation, et vérifie
   que le résultat final vaut toujours exactement 500.

## Exercice 8 — Fermeture de channel et range

Lance une goroutine qui envoie les carrés des nombres de 1 à 6 dans un
channel, puis ferme ce channel une fois terminé. Dans la fonction principale,
utilise `range` sur ce channel pour afficher chaque valeur reçue, et explique
pourquoi la boucle `range` s'arrête toute seule sans condition explicite.

## Exercice 9 — Combiner WaitGroup et Mutex

Simule un compteur de visites partagé entre 10 "pages" (goroutines). Chaque
goroutine incrémente le compteur 100 fois (donc 1000 incréments au total),
protégé par un `sync.Mutex`. Utilise un `sync.WaitGroup` pour attendre la fin
de toutes les goroutines, puis affiche le total (qui doit toujours valoir
1000).

## Exercice bonus — Pipeline de channels

Crée un pipeline en trois étapes reliées par des channels :
1. une goroutine génère les nombres de 1 à 10 et les envoie dans un premier
   channel, puis le ferme ;
2. une deuxième goroutine lit ce channel avec `range`, met chaque nombre au
   carré, envoie le résultat dans un second channel, puis le ferme ;
3. la fonction principale lit le second channel avec `range` et affiche
   chaque valeur.

Vérifie que le résultat est toujours le même (1, 4, 9, ..., 100) quelle que
soit l'exécution.
