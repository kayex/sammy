package main

import (
	"os"

	"github.com/sqweek/dialog"
)

func main() {
	//filename := "Z:\\Sound Library\\User Library\\Splice Sounds - J-E-T-S Zoospa Sounds\\JETS_tonal\\JETS_bass\\JETS_bass_one_shots\\JETS_bass_one_shot_truth_14_Bb.wav"
	//filename = "D:\\bmx.vpj"

	filename, err := dialog.File().Title("Choose file to examine").Load()
	if err != nil {
		panic(err)
	}

	fi, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	dialog.Message("Successfully read file info.\n\nName: %s\nSize: %v\nPerm: %v\nDirectory: %v\nModified: %v", fi.Name(), fi.Size(), fi.Mode(), fi.IsDir(), fi.ModTime()).Title("File examination").Info()

}
