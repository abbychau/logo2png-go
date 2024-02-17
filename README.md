# Logo2PNG

This is a simple Python script that converts logo(UCBLogo / Berkeley Logo) files to PNG images.

## Usage

```bash
go run main.go logo_commands.txt
```
Then, you will get a PNG file named `logo_commands.png` under the same directory.

## Example

```bash
setpencolor [255 0 0]
setpensize [3 3]
fd 50
rt 90
fd 50
rt 90
fd 50
rt 90
fd 50
rt 90
pu
rt 45
fd 10
fill [55 35 45]
```

Output:

![logo_commands.png](logo_commands.png)

## Supported Commands

- setpencolor
- setpensize
- fd
- rt
- pu
- pd
- fill

## Unsupported Commands (TODO)

- repeat, repcount
- setscreencolor
- setxy
- home
- to, end (define a procedure)

They are not supported yet, just because I didn't need them. I may add them if I need them.

## Architectual Todo

- Make use of a lexer and token parser

## Why I made this

I fill so lazy to draw images in mspaint or Photoshop when I am making a game. So I wanted to have a simple tool to draw images with code. I found that Logo is a good tool for this purpose instead of other very complex syntaxes like Processing, SVG, etc. BTW, Logo is a very old language and it is not very popular now. 

## License
MIT License

