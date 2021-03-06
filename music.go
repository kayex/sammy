package sammy

import "fmt"

var naturalNotes = []string{"A", "B", "C", "D", "E", "F", "G"}
var sharps = []string{"A#", "C#", "D#", "F#", "G#"}
var flats = []string{"Ab", "Bb", "Db", "Eb", "Gb"}

var sharpMirror = map[string]string{
	"C#": "Db",
	"D#": "Eb",
	"F#": "Gb",
	"G#": "Ab",
	"A#": "Bb",
}

var flatMirror = map[string]string{
	"Db": "C#",
	"Eb": "D#",
	"Gb": "F#",
	"Ab": "G#",
	"Bb": "A#",
}

func keys() []string {
	n := naturalNotes
	n = append(n, sharps...)
	n = append(n, flats...)
	return n
}

func major(key string) string {
	return fmt.Sprintf("%smaj", key)
}

func minor(key string) string {
	return fmt.Sprintf("%smin", key)
}

func flipKeySignature(key string) string {
	if s, ok := sharpMirror[key]; ok {
		return s
	}

	if f, ok := flatMirror[key]; ok {
		return f
	}

	return key
}
