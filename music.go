package sammy

var naturalNotes = []string{"A", "B", "C", "D", "E", "F", "G"}
var sharps = []string{"A#", "C#", "D#", "F#", "G#"}
var flats = []string{"Ab", "Bb", "Db", "Eb", "Gb"}

func notes() []string {
	n := naturalNotes
	n = append(n, sharps...)
	n = append(n, flats...)
	return n
}
