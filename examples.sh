#!/bin/sh 

#png
go run main.go -o examples/talk.png Go Beyond the Console: Developing 2D Games With Go
go run main.go -o examples/talk2.png -figlet banner3 -l 9500 -bgHex 0xabc -fgHex 0x000 Towards Modern Development of Cloud Applications With Service Weaver
go run main.go -o examples/talk3.png -figlet banner -l 3100 Gentle Introduction To eBPF
go run main.go -o examples/talk4.png -fontSize=48 -figlet larry3d -l 9500 -h 500 -bgHex 0xc9c -fgHex 0x000 Go Right Ahead! Simple Hacks To Cut Memory Usage by 80%
go run main.go -o examples/talk5.png -figlet roman -l 9500 -bgHex 0xabc -fgHex 0x000 Go Right Ahead! Simple Hacks To Cut Memory Usage by 80%
go run main.go -o examples/talk6.png -figlet roman -bgHex 0xa23 -fgHex 0x121 Why Integration Tests Might Be Better Than Unit Tests
go run main.go -o examples/talk7.png -figlet 3-d -bgHex 0x423 -fgHex 0xeee Coffee Break
go run main.go -o examples/talk8.png -figlet dotmatrix -bgHex 0x500 -fgHex 0xeee Coffee Break

#gif
go run main.go -o examples/talk.gif -banner Go Beyond the Console: Developing 2D Games With Go
go run main.go -o examples/talk2.gif -banner -bgHex 0xada -fgHex 0x121 Why Integration Tests Might Be Better Than Unit Tests
go run main.go -o examples/talk3.gif -figlet roman -blink -bgHex 0xabc -fgHex 0x000 Why Integration Tests Might Be Better Than Unit Tests
go run main.go -o examples/talk4.gif -figlet cosmic -banner -bgHex 0x044 -l 5656 Reaching the Unix Philosophys Logical Extreme With WebAssembly
go run main.go -o examples/talk5.gif -figlet smkeyboard -alt Panel With the Go Team
go run main.go -o examples/talk6.gif -figlet computer -blink -bgHex 0xabc -fgHex 0x000 API Optimization Tale: Monitor, Fix and Deploy
go run main.go -o examples/talk7.gif -figlet isometric3 -alt -bgHex 0x500 -fgHex 0xeee Coffee Break
