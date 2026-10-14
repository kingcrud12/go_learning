// Package concurrency explore les bases de la concurrence en Go : goroutines,
// synchronisation avec sync.WaitGroup et sync.Mutex, channels (bufferisés ou
// non), select et un pattern simple de worker pool. Chaque exemple utilise un
// mécanisme de synchronisation explicite (WaitGroup, channel bloquant) pour
// produire une sortie stable et reproductible : aucune démonstration ne
// s'appuie sur un time.Sleep comme seul moyen d'attendre une goroutine.
package concurrency

import (
	"fmt"
	"sync"
)

// goroutineBasique lance une goroutine et utilise un sync.WaitGroup pour
// attendre proprement sa fin, sans time.Sleep arbitraire.
func goroutineBasique() {
	var attente sync.WaitGroup

	attente.Add(1) // on annonce une goroutine à attendre
	go func() {
		defer attente.Done() // signale la fin du travail
		fmt.Println("goroutine basique : travail terminé")
	}()

	attente.Wait() // bloque jusqu'à ce que Done() ait été appelé
	fmt.Println("goroutine basique : la fonction principale continue")
}

// channelNonBufferise montre l'envoi et la réception d'une valeur entre deux
// goroutines via un channel non bufferisé. L'envoi (channel <- valeur) bloque
// tant qu'aucune réception n'est prête, et inversement : c'est ce blocage
// mutuel qui sert de synchronisation.
func channelNonBufferise() {
	messages := make(chan string) // channel non bufferisé

	go func() {
		messages <- "bonjour depuis la goroutine" // bloque jusqu'à réception
	}()

	recu := <-messages // bloque jusqu'à ce qu'une valeur soit envoyée
	fmt.Println("channel non bufferisé :", recu)
}

// channelBufferise montre la différence avec un channel non bufferisé :
// un channel bufferisé permet d'envoyer jusqu'à sa capacité sans bloquer,
// même si personne ne reçoit encore.
func channelBufferise() {
	nombres := make(chan int, 3) // capacité de 3

	// Ces trois envois ne bloquent pas car le buffer n'est pas plein.
	nombres <- 1
	nombres <- 2
	nombres <- 3

	// Le channel est maintenant plein : un 4e envoi bloquerait tant
	// qu'aucune valeur n'est lue. On le ferme pour pouvoir le vider avec range.
	close(nombres)

	for n := range nombres {
		fmt.Println("channel bufferisé :", n)
	}
}

// selectMultiCanaux montre comment select attend sur plusieurs channels à la
// fois et exécute le cas correspondant au premier channel prêt. Chaque
// goroutine écrit une seule fois dans son propre channel non bufferisé, et on
// lit exactement deux fois pour obtenir un résultat déterministe (les deux
// messages seront bien reçus, même si leur ordre d'arrivée peut varier).
func selectMultiCanaux() {
	canalA := make(chan string)
	canalB := make(chan string)
	var attente sync.WaitGroup

	attente.Add(2)
	go func() {
		defer attente.Done()
		canalA <- "message du canal A"
	}()
	go func() {
		defer attente.Done()
		canalB <- "message du canal B"
	}()

	// On reçoit exactement deux messages (un par channel), quel que soit
	// l'ordre dans lequel select les choisit.
	for i := 0; i < 2; i++ {
		select {
		case msgA := <-canalA:
			fmt.Println("select :", msgA)
		case msgB := <-canalB:
			fmt.Println("select :", msgB)
		}
	}

	attente.Wait() // attend la fin propre des deux goroutines
}

// workerPool illustre un pool de workers : plusieurs goroutines consomment
// des tâches depuis un même channel jusqu'à sa fermeture, et un WaitGroup
// garantit qu'on attend la fin de tous les workers avant de continuer.
func workerPool() {
	const nombreWorkers = 3
	taches := make(chan int, 10)
	resultats := make(chan int, 10)
	var attente sync.WaitGroup

	// Démarrage des workers : chacun lit des tâches jusqu'à fermeture du channel.
	for id := 1; id <= nombreWorkers; id++ {
		attente.Add(1)
		go func(idWorker int) {
			defer attente.Done()
			for tache := range taches {
				resultats <- tache * tache // le "travail" : mettre au carré
			}
		}(id)
	}

	// Envoi des tâches puis fermeture : plus aucun envoi n'est possible après close.
	for t := 1; t <= 5; t++ {
		taches <- t
	}
	close(taches)

	attente.Wait()   // attend que tous les workers aient fini de consommer
	close(resultats) // plus aucune goroutine n'écrit dans resultats désormais

	somme := 0
	for r := range resultats {
		somme += r
	}
	fmt.Println("worker pool : somme des carrés de 1 à 5 =", somme) // 1+4+9+16+25 = 55
}

// compteurSansMutex documente (en commentaire, sans l'exécuter) le problème
// classique de race condition : si plusieurs goroutines font `compteur++`
// sans protection, l'opération n'est pas atomique (lecture, incrémentation,
// écriture) et des mises à jour peuvent être perdues. Exemple dangereux
// (NE PAS exécuter tel quel, résultat imprévisible et non reproductible) :
//
//	compteur := 0
//	var attente sync.WaitGroup
//	for i := 0; i < 1000; i++ {
//		attente.Add(1)
//		go func() {
//			defer attente.Done()
//			compteur++ // race condition : plusieurs goroutines lisent/écrivent en même temps
//		}()
//	}
//	attente.Wait()
//	fmt.Println(compteur) // résultat variable, souvent < 1000
//
// La fonction compteurAvecMutex ci-dessous montre la solution correcte.
func compteurSansMutex() {
	fmt.Println("compteur sans mutex : voir le commentaire du code source pour l'exemple dangereux (non exécuté ici)")
}

// compteurAvecMutex protège l'accès concurrent à un compteur partagé grâce à
// sync.Mutex : Lock/Unlock garantit qu'une seule goroutine à la fois modifie
// la variable, ce qui rend le résultat final déterministe.
func compteurAvecMutex() {
	compteur := 0
	var mutex sync.Mutex
	var attente sync.WaitGroup

	for i := 0; i < 1000; i++ {
		attente.Add(1)
		go func() {
			defer attente.Done()
			mutex.Lock()
			compteur++ // section critique protégée par le mutex
			mutex.Unlock()
		}()
	}

	attente.Wait()
	fmt.Println("compteur avec mutex :", compteur) // toujours 1000
}

// fermetureEtRange montre comment close() signale qu'un channel n'enverra
// plus aucune valeur, et comment range sur un channel s'arrête
// automatiquement dès que celui-ci est fermé et vidé.
func fermetureEtRange() {
	valeurs := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			valeurs <- i
		}
		close(valeurs) // plus aucun envoi après close ; range saura s'arrêter
	}()

	for v := range valeurs { // s'arrête automatiquement à la fermeture du channel
		fmt.Println("fermeture et range :", v)
	}
}

// Run exécute toutes les démonstrations du package.
func Run() {
	goroutineBasique()
	channelNonBufferise()
	channelBufferise()
	selectMultiCanaux()
	workerPool()
	compteurSansMutex()
	compteurAvecMutex()
	fermetureEtRange()
}
