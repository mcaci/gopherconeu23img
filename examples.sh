#!/bin/sh 

#png
go run main.go -o examples/talk.png Go Beyond the Console: Developing 2D Games With Go
go run main.go -o examples/talk2.png -figlet banner3 -l 9500 -bgHex 0xabc -fgHex 0x000 Towards Modern Development of Cloud Applications With Service Weaver
go run main.go -o examples/talk3.png -figlet banner -l 3100 Gentle Introduction To eBPF
go run main.go -o examples/talk4.png -fontSize=48 -figlet larry3d -l 9500 -h 500 -bgHex 0xc0c -fgHex 0x000 Go Right Ahead! Simple Hacks To Cut Memory Usage by 80%
go run main.go -o examples/talk5.png -figlet roman -l 9500 -bgHex 0xabc -fgHex 0x000 Go Right Ahead! Simple Hacks To Cut Memory Usage by 80%
go run main.go -o examples/talk6.png -figlet roman -bgHex 0xa23 -fgHex 0x121 Why Integration Tests Might Be Better Than Unit Tests

#gif
go run main.go -o examples/talk.gif -banner Go Beyond the Console: Developing 2D Games With Go
go run main.go -o examples/talk2.gif -figlet roman -banner Why Integration Tests Might Be Better Than Unit Tests # TO REVIEW
go run main.go -o examples/talk3.gif -figlet roman -blink -bgHex 0xabc -fgHex 0x000 Why Integration Tests Might Be Better Than Unit Tests