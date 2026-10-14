// Package types explore les types de base et les types composés de Go :
// types numériques et leurs limites, string et rune, bool, array et slice,
// map, types nommés et conversions. Les struct, pointeurs (voir le package
// structs_and_pointers) et les generics (voir le package generics) sont
// volontairement exclus.
package types

import (
	"fmt"
	"math"
	"strconv"
)

// typesNumeriquesEntiers montre les différentes tailles d'entiers signés et
// non signés disponibles en Go.
func typesNumeriquesEntiers() {
	var i int = -42
	var i8 int8 = -8
	var i16 int16 = -16
	var i32 int32 = -32
	var i64 int64 = -64

	var u uint = 42
	var u8 uint8 = 8
	var u16 uint16 = 16
	var u32 uint32 = 32
	var u64 uint64 = 64

	fmt.Println(i, i8, i16, i32, i64)
	fmt.Println(u, u8, u16, u32, u64)
}

// limitesNumeriques montre les constantes du package math donnant les
// valeurs min/max de chaque type entier, utiles pour comprendre le
// dépassement de capacité (overflow).
func limitesNumeriques() {
	fmt.Println("int8   :", math.MinInt8, "à", math.MaxInt8)
	fmt.Println("int16  :", math.MinInt16, "à", math.MaxInt16)
	fmt.Println("int32  :", math.MinInt32, "à", math.MaxInt32)
	fmt.Println("int64  :", math.MinInt64, "à", math.MaxInt64)
	fmt.Println("uint8  : 0 à", math.MaxUint8)
	fmt.Println("uint16 : 0 à", math.MaxUint16)
}

// typesNumeriquesFlottants montre les deux types de nombres à virgule
// flottante et leur précision différente.
func typesNumeriquesFlottants() {
	var f32 float32 = 3.1415927
	var f64 float64 = 3.14159265358979

	fmt.Println(f32, f64)
	fmt.Println("précision float32 :", math.MaxFloat32)
	fmt.Println("précision float64 :", math.MaxFloat64)
}

// stringEtRune montre qu'une string est une suite d'octets (bytes) encodés
// en UTF-8, alors qu'une rune représente un point de code Unicode complet.
func stringEtRune() {
	mot := "café" // é est encodé sur 2 bytes en UTF-8

	fmt.Println("longueur en bytes :", len(mot))

	var premiereLettre byte = mot[0] // accès par index = byte brut
	fmt.Printf("premier byte : %v (%c)\n", premiereLettre, premiereLettre)

	runes := []rune(mot) // conversion explicite en slice de runes
	fmt.Println("longueur en runes :", len(runes))
	fmt.Printf("dernière rune : %c\n", runes[len(runes)-1])
}

// iterationUTF8 montre que range sur une string itère par rune (point de
// code Unicode) et non par byte, ce qui est essentiel pour les caractères
// multi-octets comme les accents ou les emojis.
func iterationUTF8() {
	mot := "café"

	for index, caractere := range mot {
		fmt.Printf("index=%d rune=%c\n", index, caractere)
	}
	// Remarquer que l'index saute de 2 entre 'f' et 'é' : 'é' occupe 2 bytes.
}

// typeBool montre le type booléen et les opérateurs logiques de base.
func typeBool() {
	var estMajeur bool = true
	var aPermis = false

	peutConduire := estMajeur && aPermis
	peutVoter := estMajeur || false

	fmt.Println(peutConduire, peutVoter)
}

// arrayTailleFixe montre un array : sa taille fait partie de son type et ne
// peut pas changer après déclaration.
func arrayTailleFixe() {
	var notes [3]int = [3]int{12, 15, 9}
	notes[1] = 18 // modification d'un élément autorisée

	fmt.Println(notes)
	fmt.Println("longueur fixe :", len(notes))
}

// sliceDynamique montre un slice : une vue dynamique au-dessus d'un array
// sous-jacent, qui peut grandir avec append.
func sliceDynamique() {
	villes := []string{"Bordeaux", "Paris"}
	fmt.Println(villes, "len:", len(villes), "cap:", cap(villes))

	villes = append(villes, "Lyon")
	fmt.Println(villes, "len:", len(villes), "cap:", cap(villes))

	villes = append(villes, "Nantes", "Lille")
	fmt.Println(villes, "len:", len(villes), "cap:", cap(villes))
	// cap peut dépasser len : Go alloue de la marge pour amortir les
	// futurs append sans réallouer à chaque fois.
}

// sliceVsArray montre la différence de déclaration entre les deux et le
// fait qu'un slice peut être créé vide puis rempli.
func sliceVsArray() {
	var tableauFixe [5]int       // array : taille 5 figée, valeurs zéro
	var tranche []int            // slice : nil, longueur 0, pas d'array sous-jacent
	tranche = append(tranche, 1) // append fonctionne même sur un slice nil

	fmt.Println(tableauFixe)
	fmt.Println(tranche, "est nil ?", tranche == nil)
}

// mapDeclarationEtAjout montre la création d'une map (table associative)
// et l'ajout de paires clé/valeur.
func mapDeclarationEtAjout() {
	ages := make(map[string]int)
	ages["Yann"] = 30
	ages["Alice"] = 25

	capitales := map[string]string{
		"France":   "Paris",
		"Belgique": "Bruxelles",
	}

	fmt.Println(ages)
	fmt.Println(capitales)
}

// mapLectureAvecVerification montre le motif val, ok := m[cle] qui permet
// de distinguer une clé absente d'une clé présente avec une valeur zéro.
func mapLectureAvecVerification() {
	ages := map[string]int{"Yann": 30}

	age, ok := ages["Yann"]
	fmt.Println("Yann :", age, "trouvé ?", ok)

	age, ok = ages["Bob"]
	fmt.Println("Bob :", age, "trouvé ?", ok) // age vaut 0, ok vaut false
}

// mapSuppression montre comment retirer une clé d'une map avec delete.
func mapSuppression() {
	stock := map[string]int{"pommes": 10, "poires": 5}
	fmt.Println("avant :", stock)

	delete(stock, "poires")
	fmt.Println("après :", stock)

	delete(stock, "clé-inexistante") // ne provoque pas d'erreur
}

// Celsius est un type nommé basé sur float64 : il reste compatible avec les
// opérations float64 mais forme un type distinct, plus sûr sémantiquement.
type Celsius float64

// celsiusVersFahrenheit convertit une température en degrés Celsius vers
// Fahrenheit.
func celsiusVersFahrenheit(c Celsius) float64 {
	return float64(c)*9/5 + 32
}

// typeNomme montre l'utilisation du type nommé Celsius pour éviter de
// confondre une température avec un simple nombre.
func typeNomme() {
	var temperature Celsius = 20

	fmt.Printf("%.1f°C = %.1f°F\n", float64(temperature), celsiusVersFahrenheit(temperature))
}

// conversionNumerique montre les conversions explicites entre types
// numériques, y compris la perte de précision possible.
func conversionNumerique() {
	var entier int = 10
	var flottant float64 = float64(entier) / 3

	var grand int64 = 300
	var petit int8 = int8(grand) // dépassement de capacité : résultat imprévisible

	fmt.Println(flottant)
	fmt.Println("300 converti en int8 :", petit) // overflow volontaire pour illustrer le danger
}

// conversionVersString montre comment transformer un nombre en chaîne de
// caractères grâce au package strconv, et l'opération inverse.
func conversionVersString() {
	age := 30
	ageEnTexte := strconv.Itoa(age) // int vers string

	texte := "42"
	nombre, err := strconv.Atoi(texte) // string vers int, avec gestion d'erreur
	if err != nil {
		fmt.Println("erreur de conversion :", err)
	}

	fmt.Println("age en texte :", ageEnTexte, "(type string)")
	fmt.Println("texte en nombre :", nombre, "(type int)")

	flottantEnTexte := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Println("flottant en texte :", flottantEnTexte)
}

// Run exécute toutes les démonstrations du package.
func Run() {
	typesNumeriquesEntiers()
	limitesNumeriques()
	typesNumeriquesFlottants()
	stringEtRune()
	iterationUTF8()
	typeBool()
	arrayTailleFixe()
	sliceDynamique()
	sliceVsArray()
	mapDeclarationEtAjout()
	mapLectureAvecVerification()
	mapSuppression()
	typeNomme()
	conversionNumerique()
	conversionVersString()
}
