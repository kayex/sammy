package sammy

func naturalNotes() []string {
	return []string{"A", "B", "C", "D", "E", "F", "G"}
}

func sharps() []string {
	return []string{"A#", "C#", "D#", "F#", "G#"}
}

func flats() []string {
	return []string{"Ab", "Bb", "Db", "Eb", "Gb"}
}

func notes() []string {
	n := naturalNotes()
	n = append(n, sharps()...)
	n = append(n, flats()...)
	return n
}
