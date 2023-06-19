# gopherconeu23img
This repo contains a CLI that produces some proposals for images to use during gopherconeu23

## How to use it

Get it with `go install github.com/mcaci/gopherconeu23img@latest` then run it as `gopherconeu23img [flags] text` to get a first example with the text of your choice.

There are some flags that can be used to customize the image that is produced.

This is the flag list:

- bgHex: this flag gets a string in the form of "0x" followed by 3 or 6 hexadecimal digits. The value is used to select the color of the background.
- fgHex: this flag gets a string in the form of "0x" followed by 3 or 6 hexadecimal digits. The value is used to select the color of the text in the foreground.
- h: this flag is just the height of the image. It's an integer.
- l: this flag is just the width of the image. It's an integer.
- o: name of the output image
- fontPath: path for the font to use to draw the text in the foreground. It should be a path to a valid TrueType font.
- fontSize: size of the font to use to draw the text in the foreground.
- xPtFactor: this number is a factor used to determine the width of the character box for each character. It is used to adjust the alignment of each character of the ASCII art text drawn.
- yPtFactor: this number is a factor used to determine the height of the character box for each character. It is used to adjust the alignment of each character of the ASCII art text drawn.
- figlet: name of the figlet font: figlets are fonts used to convert text into ASCII art. See https://github.com/common-nighthawk/go-figure/tree/master/fonts for the possible values and http://www.figlet.org/examples.html to see examples of what are the effects of these fonts.
- banner: if set to true it will produce a gif of a banner that shows the text sliding on the background, if false it will produce a png.
- blink: if set to true it will produce a gif of the text blinking on the background, if false it will produce a png.
- alt: if set to true it will produce a gif of the text blinking and switching colors with the background, if false it will produce a png.
- delay: used with `banner`, `blink` or `alt`, it indicates the delay between each frame of the gif.

These flags can be listed using the `--help` flag.

## Examples

Here are some examples with the commands run to create them:

1. gopherconeu23img -o talk.png Go Beyond the Console: Developing 2D Games With Go
![Example 1](./examples/talk.png)

2. gopherconeu23img -o examples/talk2.png -figlet banner3 -bgHex 0xb4d9ef -fgHex 0x000 Towards Modern Development of Cloud Applications With Service Weaver
![Example 2](./examples/talk2.png)

3. gopherconeu23img -o examples/talk3.png -figlet larry3d -l 7500 -h 500 -bgHex 0xff9e99 -fgHex 0x325e5e Go Right Ahead! Simple Hacks To Cut Memory Usage by 80%
![Example 3](./examples/talk3.png)

4. gopherconeu23img -o examples/talk4.png -fontSize=48 -figlet banner -bgHex 0xc9c -fgHex 0x000 Gentle Introduction To eBPF
![Example 4](./examples/talk4.png)

5. gopherconeu23img -o examples/talk5.png -figlet roman -bgHex 0xabc -fgHex 0x000 Go Right Ahead! Simple Hacks To Cut Memory Usage by 80%
![Example 5](./examples/talk5.png)

6. gopherconeu23img -o examples/talk6.png -figlet roman -bgHex 0xf66 -fgHex 0x121 Why Integration Tests Might Be Better Than Unit Tests
![Example 6](./examples/talk6.png)

7. gopherconeu23img -o examples/talk7.png -figlet 3-d -bgHex 0x423 -fgHex 0xeee Coffee Break
![Example 7](./examples/talk7.png)

8. gopherconeu23img -o examples/talk8.png -figlet dotmatrix -bgHex 0x500 -fgHex 0xeee Coffee Break 
![Example 8](./examples/talk8.png)

9. gopherconeu23img -o examples/talk9.png -figlet alligator2 -fgHex 0x121 -bgHex 0xada Useful Functional-Options Tricks For Better Libraries
![Example 9](./examples/talk9.png)

10. gopherconeu23img -o examples/talk10.png -figlet speed -fgHex 0xdad GoTime Podcast Live
![Example 10](./examples/talk10.png)

### Banner Gif

1. gopherconeu23img -o examples/talk.gif -banner Go Beyond the Console: Developing 2D Games With Go
![Example 1](./examples/talk.gif)

2. gopherconeu23img -o examples/talk2.gif -figlet epic -banner -bgHex 0xada -fgHex 0x121 Why Integration Tests Might Be Better Than Unit Tests
![Example 2](./examples/talk2.gif)

3. gopherconeu23img -o examples/talk4.gif -figlet cosmic -banner -bgHex 0x044 Reaching the Unix Philosophys Logical Extreme With WebAssembly
![Example 3](./examples/talk4.gif)

4. gopherconeu23img -o examples/talk8.gif -figlet smkeyboard -banner -delay 10 Panel With the Go Team
![Example 4](./examples/talk8.gif)

### Blinking Gif

1. gopherconeu23img -o examples/talk3.gif -figlet roman -blink -bgHex 0xffd364 -fgHex 0x000 Why Integration Tests Might Be Better Than Unit Tests
![Example 1](./examples/talk3.gif)

2. gopherconeu23img -o examples/talk9.gif -figlet speed -blink -bgHex 0xb4d9ef -fgHex 0x284b4b GoTime Podcast Live
![Example 2](./examples/talk9.gif)

3. gopherconeu23img -o examples/talk10.gif -blink -delay 125 A Fast Structured Logging Package
![Example 3](./examples/talk10.gif)

### Blinking Gif with alternating colors 

1. gopherconeu23img -o examples/talk5.gif -figlet smkeyboard -alt -bgHex 0x00ADD8 -fgHex 0xFFF Panel With the Go Team
![Example 1](./examples/talk5.gif)

2. gopherconeu23img -o examples/talk6.gif -figlet computer -alt -delay 85 -bgHex 0x78C475 -fgHex 0x030303 API Optimization Tale: Monitor, Fix and Deploy
![Example 2](./examples/talk6.gif)

3. gopherconeu23img -o examples/talk7.gif -figlet isometric3 -alt -delay 150 -bgHex 0xffbeda -fgHex 0x000 Coffee Break
![Example 3](./examples/talk7.gif)

4. gopherconeu23img -o examples/talk11.gif -figlet banner3-D -alt -bgHex 0x5DC9E2 -fgHex 0xCE3262 -delay 125 Keynote - State of the Go Nation
![Example 4](./examples/talk11.gif)

5. gopherconeu23img -o examples/talk12.gif -figlet contrast -alt How To Avoid Breaking Changes in Your Go Modules
![Example 5](./examples/talk12.gif)