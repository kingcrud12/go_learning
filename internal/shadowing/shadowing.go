// Package shadowing explore le masquage de variables (variable shadowing) en
// Go : une variable déclarée dans un bloc interne qui porte le MÊME NOM
// qu'une variable d'un bloc englobant, et qui la masque le temps du bloc.
//
// Attention à ne pas confondre avec la simple portée de bloc (voir la
// fonction portee() du package variables, qui utilise des noms différents
// x/y) : ici, c'est le même identifiant qui est réutilisé via := dans un
// bloc interne, ce qui crée une NOUVELLE variable distincte plutôt que de
// modifier celle du bloc externe. C'est ce piège spécifique qui est illustré
// ici, pas la portée en général.
package shadowing

import "fmt"

// shadowingBasique montre qu'une variable déclarée avec := dans un bloc if
// masque la variable externe de même nom, sans la modifier.
func shadowingBasique() {
	message := "externe"

	if true {
		message := "interne" // nouvelle variable, masque celle du dessus
		fmt.Println("dans le bloc :", message)
	}

	fmt.Println("après le bloc :", message) // toujours "externe"
}

// doSomething simule un appel qui peut échouer, pour illustrer le piège
// classique du shadowing avec err.
func doSomething() error {
	return nil
}

// autreChose simule un second appel qui renvoie une valeur et une erreur.
func autreChose() (int, error) {
	return 42, fmt.Errorf("erreur dans autreChose")
}

// piegeErrShadowing montre le bug classique en Go : un err interne, déclaré
// avec := dans un bloc if, masque le err externe. L'erreur réelle est donc
// ignorée car on teste et on affiche le mauvais err.
func piegeErrShadowing() {
	err := doSomething()

	if x, err := autreChose(); x > 0 {
		// ce err est une NOUVELLE variable, locale au bloc if,
		// différente du err externe déclaré juste au-dessus
		fmt.Println("err interne (masqué) :", err)
	}

	// ici, err vaut toujours nil : l'erreur renvoyée par autreChose()
	// n'a jamais été vue par le code externe ! C'est le piège classique.
	fmt.Println("err externe (bug : toujours nil) :", err)
}

// shadowingDansUneBoucle montre qu'une variable de boucle peut masquer une
// variable externe de même nom.
func shadowingDansUneBoucle() {
	valeur := "avant la boucle"

	for i := 0; i < 3; i++ {
		valeur := fmt.Sprintf("iteration %d", i) // masque la variable externe
		fmt.Println("dans la boucle :", valeur)
	}

	fmt.Println("après la boucle :", valeur) // toujours "avant la boucle"
}

// declarationVsAffectation montre la différence entre := (déclare une
// NOUVELLE variable, donc peut créer du shadowing) et = (réaffecte la
// variable existante, sans shadowing).
func declarationVsAffectation() {
	compteur := 10

	if true {
		compteur := 99 // := : nouvelle variable locale, masque l'externe
		fmt.Println("avec := dans le bloc :", compteur)
	}
	fmt.Println("après le bloc avec := :", compteur) // 10, inchangé

	if true {
		compteur = 99 // = : réaffecte la variable externe existante
		fmt.Println("avec = dans le bloc :", compteur)
	}
	fmt.Println("après le bloc avec = :", compteur) // 99, bien modifié
}

// commentDetecterLeShadowing rappelle les outils permettant de détecter ce
// genre de bug avant qu'il ne cause des problèmes en production.
//
// go vet intègre une analyse "shadow" (package golang.org/x/tools/go/analysis
// /passes/shadow), installable via :
//
//	go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
//	go vet -vettool=$(which shadow) ./...
//
// Le linter golangci-lint propose aussi un linter nommé "shadow" (via
// govet avec l'option shadow activée) qui peut être activé dans la config
// .golangci.yml pour détecter ces masquages automatiquement dans la CI.
func commentDetecterLeShadowing() {
	fmt.Println("voir le commentaire de commentDetecterLeShadowing pour les outils (go vet -shadow, golangci-lint)")
}

// Run exécute toutes les démonstrations du package.
func Run() {
	shadowingBasique()
	piegeErrShadowing()
	shadowingDansUneBoucle()
	declarationVsAffectation()
	commentDetecterLeShadowing()
}
