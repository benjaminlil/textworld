package main

import (
	"fmt"
	"time"
)

var harstav bool = false
var swedish = true

var harYxa = false

func main() {
	fmt.Printf("a) svenska")
	fmt.Printf(" ")
	fmt.Printf("b) English")
	fmt.Println()

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		swedish = true
	} else if svar == "b" {
		swedish = false
	}

	utanförHuset()
}

func utanförHuset() {
	messages := []string{"Hello\nHello again\n", "hej\nhej igen"}

	var message string
	if swedish {
		message = messages[1]
	} else {
		message = messages[0]
	}

	fmt.Printf(message)

	fmt.Println()
	fmt.Println("Utanför huset")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("Du står utanför ett grönt mögligt hus.")
	fmt.Println("Vad gör du?")
	fmt.Println("a) Går in")
	fmt.Println("b) Går runt huset och kollar")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		iHallen()
	} else if svar == "b" {
		gåRunt()
	}

}

func iHallen() {
	fmt.Println()
	fmt.Println("I hallen")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("Hallen var smutsig och luktade skunk.")
	fmt.Println("Du ser ett kök och en trappa ner.")
	fmt.Println("Vad gör du?")
	fmt.Println("a) Går ner för trappan.")
	fmt.Println("b) Går in i köket.")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		iKällaren()
	} else if svar == "b" {
		iKöket()
	}

}

func gåRunt() {
	fmt.Println()
	fmt.Println("På baksidan")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("På baksidan såg du en stege.")
	fmt.Println("Vad gör du?")
	fmt.Println()
	fmt.Println("a) Går tillbaks till framsidan.")
	fmt.Println("b) Tar stegen och klättrar upp på taket.")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		utanförHuset()
	} else if svar == "b" {
		påTaket()
	}

}

func påTaket() {
	fmt.Println()
	fmt.Println("På taket")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("På taket finns en skorsten.")
	fmt.Println("Vad gör du?")
	fmt.Println()
	fmt.Println("a) Går ner igen.")
	fmt.Println("b) Gör rent skorstenen.")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		gåRunt()
	} else if svar == "b" {
		görRent()
	}

}

func görRent() {
	fmt.Println()
	fmt.Println("Städar")
	fmt.Println("_______________")
	fmt.Println()

	if harYxa == true {
		fmt.Println("Värst vad nogran du är, du har ju redan städat")
	} else {
		fmt.Println("Du städar skorstenen.")
		fmt.Println("Du hittar en yxa i skorstenen.")
		fmt.Println("Grattis!!!")
		harYxa = true
	}
	fmt.Println("Vad gör du nu?")
	fmt.Println("a) Går ner igen.")
	fmt.Println("b) Leker på taket.")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		gåRunt()
	} else if svar == "b" {
		fmt.Println("Du halkar och yxan åker upp i luften, medans du faller ner.")
		fmt.Println("Du landar och känner:\n - Skönt jag dog inte.")
		time.Sleep(time.Second * 7)
		fmt.Println("Sen landar yxan, den landar rakt i bröstkorgen. Trist, du förlora..")
	}

}

func iKällaren() {
	fmt.Println()
	fmt.Println("I källaren")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("källaren är mörk och stinker.")
	fmt.Println("Du ser ett monster springa mot dig.")

	if harYxa {
		// Gör vad som skall hända om du har yxa
		fmt.Println("Du dödar monstret med yxan, grattis!!!")
		fmt.Println("Du går vidare i källaren och ser en väg ner i kloakerna.")
		fmt.Println()
		fmt.Println("Du går ner i kloakerna och det luktar död fisk.")
		iKloakerna()
	} else {
		// Om man inte har yxa
		fmt.Println("Vad gör du?")
		fmt.Println("a)Slåss mot den med knytnävarna.")
		fmt.Println("b)Springer upp igen.")
		var svar string
		fmt.Scanln(&svar)

		if svar == "a" {
			fmt.Println("Du dog för att monstret åt upp dig.")
		} else if svar == "b" {
			iHallen()
		}

	}

}

func iKöket() {
	fmt.Println()
	fmt.Println("I köket")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("Du har lite smutsiga händer.")
	fmt.Println("Vad gör du?")
	fmt.Println()
	fmt.Println("a) Går ut ur köket.")
	fmt.Println("b) Tvättar händerna.")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		iHallen()
	} else if svar == "b" {
		fmt.Println("När du tvättar händerna så exploderar en bomb som aktiverades när du tvättade händerna.")
		fmt.Println("Tyvär du dog...")
	}

}

func iKloakerna() {
	fmt.Println()
	fmt.Println("I kloakerna")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("Det var mörkt nere i kloakerna.")
	fmt.Println("Vad gör du?")
	fmt.Println()
	fmt.Println("a) Går åt vänster.")
	fmt.Println("b) Går åt höger.")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		gåVänster()
	} else if svar == "b" {
		gåHöger()
	}

}

func gåHöger() {
	fmt.Println()
	fmt.Println("Högra delen")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("I den högra delen av kloaken finns en stav.")
	fmt.Println("Vad gör du?")
	fmt.Println()
	fmt.Println("a) Tar staven.")
	fmt.Println("b) går tillbaks.")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		harstav = true
		tarStaven()
	} else if svar == "b" {
		iKloakerna()
	}

}

func tarStaven() {
	fmt.Println()
	fmt.Println("Har staven")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("Du tar staven.")
	fmt.Println("Vad gör du?")
	fmt.Println()
	fmt.Println("a) Leker att staven har krafer.")
	fmt.Println("b) går tillbaks.")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		fmt.Println("Trollstaven har visst krafter så du sjuter i väggen och stog nära.")
		fmt.Println("Du dog för att du stog för nära. Synd")
	} else if svar == "b" {
		iKloakerna()
	}

}

func gåVänster() {
	fmt.Println()
	fmt.Println("Går vänster")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("Du går vänster i kloakerna.")
	fmt.Println("Nu är du i vänstra delen av kloakerna.")
	fmt.Println("Du ser ett troll med guldrustning på.")
	fmt.Println("Vad gör du?")
	if harstav {
		fmt.Println()
		fmt.Println("a) Kasta din yxa mot den")
		fmt.Println("b) Rikta staven mot den.")
	} else {
		fmt.Println("Det var ett troll där som dödar dig.")
		fmt.Println("Du dog... Trist.")
	}

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		fmt.Println("trollets rustning gjorde så att din yxa inte skadade något.")
		fmt.Println("Då sjuter trollet på dig. Du dog. Synd")
	} else if svar == "b" {
		fmt.Println("staven sjuter ett ljus mot trollet.")
		fmt.Println("Ljuset dödar trollet och du fortsätter.")
		gåVidare()
	}

}

func gåVidare() {
	fmt.Println()
	fmt.Println("Går vidare")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("Du går vidare i kloakerna.")
	fmt.Println("Du ser en väg ut och en väg åt vänster.")
	fmt.Println("Vad gör du?")
	fmt.Println()
	fmt.Println("a)gå mot utvägen")
	fmt.Println("b)gå vänster")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		fmt.Println("Nu är du ute.")
		därUte()
	} else if svar == "b" {
		fmt.Println("Du går vänster och faller ner i en fallucka.")
		fmt.Println("Tyvär så dog du.")
	}

}

func därUte() {
	fmt.Println()
	fmt.Println("Utomhus")
	fmt.Println("_______________")
	fmt.Println()

	fmt.Println("Du ser ett troll till väster och ett monster till höger.")
	fmt.Println("Vart går du?")
	fmt.Println("a)Rakt fram")
	fmt.Println("b)Gå höger")

	var svar string
	fmt.Scanln(&svar)

	if svar == "a" {
		fmt.Println("Du klarade det sista testet.")
		fmt.Println("Grattis!!!")
	} else if svar == "b" {
		fmt.Println("Du går höger och trollet till höger sjöt dig i ryggen, med sin trollstav.")
		fmt.Println("Du dog... Trist)=")
	}

}
