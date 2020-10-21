# sammy
Sammy is a simple tool for normalizing audio sample filenames.

# Functionality
Sammy tries to normalize filenames by following a few simple rules:
- Naked keys are assumed to be major and will be suffixed with "maj" (`A` becomes `Amaj`).
- Minor keys on the format `Am` will have their suffix changed to "min" (`Am` becomes `Amin`).
- Keys with flat key signatures will be converted into the corresponding sharp keys (`Eb` becomes `D#`).

Below are a few examples of normalizations that sammy performs.

| Original             | Normalized           |
| :--------------------| :------------------- |
| Chords_A_120.wav     | Chords_Amaj_120.wav  |
| Chords_Am_120.wav    | Chords_Amin_120.wav  |
| CHORDS-AM-120.wav    | CHORDS-Amin-120.wav  |
| Chords_Db_120.wav    | Chords_C#maj_120.wav |
| Chords_Ebmin_120.wav | Chords_D#min_120.wav |


# Building
```sh
./build.sh
```

To enable debug console, build using
```sh
./build.sh debug
```

# Supported file extensions
- .wav, .wave
- .flac
- .mp3, .mp4
- .aiff
- .ogg, .ogv, .oga, .ogx, .ogm, .spx, .opus
