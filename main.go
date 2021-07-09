package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Jinwook-Song/nomadcoin-2021/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/",home)
	fmt.Printf("Listening on http://localhost%s ✅\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

