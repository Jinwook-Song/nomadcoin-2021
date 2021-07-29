package main

import (
	"github.com/Jinwook-Song/nomadcoin-2021/explorer"
	"github.com/Jinwook-Song/nomadcoin-2021/rest"
)



func main() {
	go explorer.Start(5000)
	rest.Start(4000)
}