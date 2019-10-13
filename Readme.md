# mp3duration

A Golang package for calculating the duration of an mp3 file. Ported from <https://github.com/ddsol/mp3-duration>.

## Usage

```golang
duration, err := mp3duration.Calculate("src/mp3duration/testdata/demo - cbr.mp3")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Duration %v\n", duration)
```
