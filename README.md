# webp2png

**Note that this project has been considered deprecated in favor of using
[img-convert](https://github.com/xubiod/img-convert) for conversions from and to
multiple other formats, also made in Go and works the same way.**

## Usage

`go run main.go image1.webp [image2.webp] [..]`

The output file will be in the same directory as the inputs, appending `.webp`
onto the original file. For an input of `/path/to/image.webp`, the program will
convert and put the output as `/path/to/image.webp.png`.
