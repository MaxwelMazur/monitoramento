package main

import (
	"fmt"
	"net/http"
	"os"
)

var sites = []string{
	"https://venturasinformatica.com.br/",
	"https://www.gov.br/pt-br",
	"https://www.playlostark.com/",
	"https://grupointegrado.instructure.com/",
	"http://www.github.com/",
	"http://www.alura.com.br/",
	"https://downdetector.com.br/",
	"https://grupointegrado.br/",
	"https://medarcade.grupointegrado.br/",
	"https://www.professorburnes.com.br/",
	"https://vizer.tv/",
	"https://naoexite.nois/",
	"http://www.orkut.com/",
	"https://discord.com/login",
	"https://projectcheats.com/ui/#",
	"https://www.facebook.com/",
	"https://www.mercadolivre.com.br/",
	"https://outlook.live.com/owa/",
	"https://www.kabum.com.br/",
}

func main() {
	for {
		menu()
		var arg string
		fmt.Scan(&arg)

		switch arg {
		case "1":
			adicionarSite()
		case "2":
			verificaSites()
		default:
			os.Exit(0)
		}
	}
}

func adicionarSite() {
	fmt.Println("adicionar site")
	sites = append(sites, "https://www.github.com/MaxwelMazur")
}

func menu() {
	fmt.Println("### Seja bem-vindo ###")
	fmt.Println("Escolha a opção:")
	fmt.Println("1 - adicione site para verificar")
	fmt.Println("2 - verificar todos sites")
	fmt.Println("3 - deixar monitorando")
	fmt.Println("4 - 'ctrl-c' para sair")
}

func verificaSites() {
	for _, v := range sites {
		resp, err := http.Get(v)
		if err != nil {
			fmt.Println("Resposta: Este site não teve resposta")
			continue
		}
		fmt.Println("Resposta: ", v, " - ", resp.Status)
	}
}
