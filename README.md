# gopherconeu23img
This repo contains a CLI that produces some proposals for images to use during gopherconeu23

## How to use it

Get it with `go install gitHub.com/mcaci/gopherconeu23img@latest` then run it as `gophercon23img $TEXT` to get a first example with the text of your choice.

There are some flags that can be used to customize the image that is produced.

This is the flag list:

- bgHex: this flag gets a string in the form of "0x" followed by 3 or 6 hexadecimal digits. The value is used to select the color of the background.
- l: this flag is just the width of the image. It's an integer.
- h: this flag is just the height of the image. It's an integer.
- fgHex: this flag gets a string in the form of "0x" followed by 3 or 6 hexadecimal digits. The value is used to select the color of the text in the foreground.
- fontPath: path for the font to use to draw the text in the foreground. It should be a path to a valid TrueType font.
- fontSize: size of the font to use to draw the text in the foreground.
- xPtFactor: this number is a factor used to determine the width of the character box for each character. It is used to adjust the alignment of each character of the ASCII art text drawn.
- yPtFactor: this number is a factor used to determine the height of the character box for each character. It is used to adjust the alignment of each character of the ASCII art text drawn.
- o: name of the output image
- figlet: name of the figlet font: figlets are fonts used to convert text into ASCII art. See https://github.com/common-nighthawk/go-figure/tree/master/fonts for the possible values and http://www.figlet.org/examples.html to see examples of what are the effects of these fonts.

These flags can be listed using the `--help` flag.

## Examples

Here are some examples with the commands run to create them:

1. go run main.go -o talk.png Go Beyond the Console: Developing 2D Games With Go
![Example 1](./examples/talk.png)

2. go run main.go -o examples/talk2.png -figlet banner3 -l 9500 -bgHex 0xabc -fgHex 0x000 Towards Modern Development of Cloud Applications With Service Weaver
![Example 2](./examples/talk2.png)

3. go run main.go -o examples/talk3.png -figlet banner -l 3100 Gentle Introduction To eBPF
![Example 3](./examples/talk3.png)

4. go run main.go -o examples/talk4.png -fontSize=48 -figlet larry3d -l 9500 -h 500 -bgHex 0xc9c -fgHex 0x000 Go Right Ahead! Simple Hacks To Cut Memory Usage by 80%
![Example 4](./examples/talk4.png)

5. go run main.go -o examples/talk5.png -figlet roman -l 9500 -bgHex 0xabc -fgHex 0x000 Go Right Ahead! Simple Hacks To Cut Memory Usage by 80%
![Example 5](./examples/talk5.png)

6. go run main.go -o examples/talk6.png -figlet roman -bgHex 0xa23 -fgHex 0x121 Why Integration Tests Might Be Better Than Unit Tests
![Example 6](./examples/talk6.png)

7. go run main.go -o examples/talk7.png -figlet 3-d -bgHex 0x423 -fgHex 0xeee Coffee Break
![Example 7](./examples/talk7.png)

8. go run main.go -o examples/talk8.png -figlet dotmatrix -bgHex 0x500 -fgHex 0xeee Coffee Break 
![Example 8](./examples/talk8.png)

### Banner Gif

1. go run main.go -o examples/talk.gif -gif Go Beyond the Console: Developing 2D Games With Go
![Example 1](./examples/talk.gif)

2. go run main.go -o examples/talk2.gif -banner -bgHex 0xada -fgHex 0x121 Why Integration Tests Might Be Better Than Unit Tests
![Example 2](./examples/talk2.gif)

3. go run main.go -o examples/talk4.gif -figlet cosmic -banner -bgHex 0x044 -l 5656 Reaching the Unix Philosophys Logical Extreme With WebAssembly
![Example 3](./examples/talk4.gif)

### Blinking Gif

1. go run main.go -o examples/talk3.gif -figlet roman -blink -bgHex 0xabc -fgHex 0x000 Why Integration Tests Might Be Better Than Unit Tests
![Example 1](./examples/talk3.gif)

2. go run main.go -o examples/talk6.gif -figlet computer -blink -bgHex 0xabc -fgHex 0x000 API Optimization Tale: Monitor, Fix and Deploy
![Example 2](./examples/talk6.gif)

### Blinking Gif with alternating colors 

1. go run main.go -o examples/talk5.gif -figlet smkeyboard -alt Panel With the Go Team
![Example 1](./examples/talk5.gif)

2. go run main.go -o examples/talk7.gif -figlet isometric3 -alt -bgHex 0x500 -fgHex 0xeee Coffee Break
![Example 2](./examples/talk7.gif)
