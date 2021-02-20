package support2

import (
	"fmt"       // Pour permettre d'utiliser fmt.print..
	"io/ioutil" // Pour permettre d'utiliser ioutil.ReadFile
)

func AsciiArtWeb2(chaine string, banner string) string {

	chaine = texte(chaine, banner)
	return chaine
}

func texte(mot string, banner string) string {

	lire, err := ioutil.ReadFile("standard.txt")

	if banner == "Standard" {
		lire, err = ioutil.ReadFile("standard.txt")
	} else if banner == "Shadow" {
		lire, err = ioutil.ReadFile("shadow.txt")
	} else if banner == "Thinkertoy" {
		lire, err = ioutil.ReadFile("thinkertoy.txt")
	}

	if err != nil {
		return ""
	}

	//lire, fichier := ioutil.ReadFile("standard.txt") // Dans ce bloc, on demande de lire le fichier qui se situe dans les guillemets
	//if fichier != nil {                              // On instaure une condition, dans laquelle si le fichier ne contient rien alors
	//	fmt.Print(fichier) // Il nous renverra une erreur
	//}

	read := string(lire)     // On converti le fichier en type string
	standard := []rune(read) // On crée un tableau de rune pour la variable read
	word := []rune(mot)      // On crée un tableau de rune pour la variable mot
	var result [10]string
	var ligne int = 0
	var finalstr string

	var ascii int = 32 // Pour permettre de commencer aux caractères imprimable de la table ASCII --> 32 : Space

	for _, c := range word { // On parcourt entièrement la variable word qu'on connait en tant que rune(mot), soit le texte qu'on veut imprimer
		for i := range standard { // On parcourt entièrement la variable file qu'on connait en tant que rune(str), & str était une string(b), on parcourt alors le fichier Standard.txt
			if ascii == int((c + 1)) { // Permet de commencer aux caractères imprimable de la table ASCII --> 33 : A
				break // Permet de casser la boucle
			}
			if i < len(standard)-2 {
				if standard[i] == 10 && standard[i+1] == 10 {
					ascii++
					ligne = 0
				}
			}
			if int(c) == ascii {
				if standard[i] == 10 {
					ligne++
				} else if ligne < 10 {
					result[ligne] = result[ligne] + string(standard[i])
				}
			}
		}
		ascii = 32
	}
	for i := range result {
		fmt.Println(result[i])
		finalstr = finalstr + result[i] + `<br>`
	}
	return finalstr
}
