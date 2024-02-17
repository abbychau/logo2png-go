# Logo2PNG

This is a simple Python script that converts logo(UCBLogo / Berkeley Logo) files to PNG images.

## Usage

```bash
go run . logo_commands.txt 
# or `go build` and run `logo2png logo_commands.txt`
```
Then, you will get a PNG file named `logo_commands.png` under the same directory.

## Example

```log
setpencolor [255 0 0]
setpentype circle
setpensize [16 5]
fd 200
rt 90
fd 50
rt 90
fd 200
rt 90
fd 50
rt 90
pu
rt 45
fd 60
fill [55 35 45]
```

Output:

![logo_commands.png](logo_commands.png)

## Supported Commands

- setpencolor
- setpensize
- fd
- bk
- rt
- pu
- pd
- fill

## Customized Commands

- setpencolor is now supporting RGBA values. Alpha value is the forth value in the list. For example, [255 0 0 255] is red with full opacity.
- setpentype (New). You can do setpentype circle or setpentype square. The default is square(same as traditional logo).

## Unsupported Commands (TODO)

- repeat, repcount
- setscreencolor
- setxy
- home
- to, end (define a procedure)
- cs (clearscreen, probably not needed)
- label (I will want it)

They are not supported yet, just because I didn't need them. I may add them if I need them.

## Architectual Todo

- Make use of a lexer and token parser

## Why I made this

I fill so lazy to draw images in mspaint or Photoshop when I am making a game. So I wanted to have a simple tool to draw images with code. I found that Logo is a good tool for this purpose instead of other very complex syntaxes like Processing, SVG, etc. So I made this tool to convert Logo commands to PNG images.

## License
MIT License

