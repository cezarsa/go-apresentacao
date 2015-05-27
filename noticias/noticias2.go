package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Noticia struct {
	Url    string
	Titulo string
	Estado string
	Foto   string
}

func (n *Noticia) String() string {
	return n.Estado + ": " + n.Titulo
}

func buscarNoticia(estado string, noticiasChan chan<- *Noticia) {
	url := "http://c.api.globo.com/news/" + estado + ".json"
	rsp, err := http.Get(url)
	if err != nil {
		noticiasChan <- nil
		return
	}
	var noticias []Noticia
	json.NewDecoder(rsp.Body).Decode(&noticias)
	if len(noticias) > 0 {
		noticias[0].Estado = estado
		noticiasChan <- &noticias[0]
	} else {
		noticiasChan <- nil
	}
}

// Atenção, para fins de simplificação esse código não está completamente
// correto. Essa página parecerá estar "carregando" para sempre, mesmo após
// todas as notícias terem sido exibidas. Veja o noticias3.go para a solução
// desse problema usando WaitGroup.
func main() {
	estados := []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		noticias := make(chan *Noticia)
		for _, estado := range estados {
			go buscarNoticia(estado, noticias)
		}
		t, _ := template.ParseFiles("noticias/view.html")
		t.Execute(w, noticias)
	})
	http.ListenAndServe(":9090", nil)
}
