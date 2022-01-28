package Rivercrossing

var state = "[kylling rev korn hs ---\\ \\__/ _________________/---]"
var name = "Ardian"

func ViewState() string {
	// Sjekke data som er lagret ("kylling til venstre", "rev til venstre")
	return state
}
func Putinboat() {
	state = "[kylling rev hs ---\\ \\_korn_/ _________________/---]"

}
func Putinrev() {
	state = "[kylling korn hs ---\\ \\_rev_/ _________________/---]"

}
func Getname() string {
	return name

}
