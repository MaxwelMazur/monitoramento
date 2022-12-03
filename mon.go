package main

import (
	"fmt"
	"math"
)

type link struct {
	Link string `json:"link"`
}

//func main() {
//	menu := tui.NewMenu(tui.DefaultStyle())
//	menu.Title = "Test"
//	menu.Description = "Test app"
//
//	menu.Commands = []tui.Command{
//		tui.Command{
//			Title:       "No Args",
//			Cli:         "echo hello!",
//			Description: "just being polite",
//			Success:     "Yey it works",
//			PrintOut:    true,
//		},
//	}
//
//	menu.Show()
//	go menu.EventManager()
//	<-menu.Wait
//	menu.Quit()
//}

func main() {

	a()

	//checkErr := func(err error) {
	//	if err != nil {
	//		Info("Err: ", err.Error())
	//		return
	//	}
	//}
	//
	//file, err := os.ReadFile("./links.json")
	//checkErr(err)
	//
	//var l []link
	//err = json.Unmarshal(file, &l)
	//checkErr(err)
	//
	//newItem := link{
	//	Link: "http://sitenotexist.bf",
	//}
	//
	//l = append(l, newItem)
	//
	//for _, v := range l {
	//	Info(v.Link)
	//}
	//
	//newByte, err := json.MarshalIndent(l, " ", "	")
	//checkErr(err)
	//
	//err = os.WriteFile("./links.json", newByte, os.ModePerm)
	//checkErr(err)

	//for {
	//	menu()
	//	var arg string
	//	fmt.Scan(&arg)
	//
	//	switch arg {
	//	case "1":
	//		adicionarSite()
	//	case "2":
	//		verificaSites()
	//	default:
	//		os.Exit(0)
	//	}
	//}
}

func adicionarSite() {
	//fmt.Println("adicionar site")
	//sites = append(sites, "https://www.github.com/MaxwelMazur")
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
	//for _, v := range sites {
	//	resp, err := http.Get(v)
	//	if err != nil {
	//		fmt.Println("Resposta: Este site não teve resposta")
	//		continue
	//	}
	//	fmt.Println("Resposta: ", v, " - ", resp.Status)
	//}
}

func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

type geo interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

type circulo struct {
	raio float64
}

func (r retangulo) area() float64 {
	return r.largura * r.altura
}

type B struct{}

func (b B) area() float64 {
	return 0
}

//	func (r rect) perim() float64 {
//		return 2*r.width + 2*r.height
//	}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

//func (c circle) perim() float64 {
//	return 2 * math.Pi * c.radius
//}

func measure(g geo) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func a() {
	r := retangulo{34.1, 12.2}
	c := circulo{raio: 11.1}

	measure(r)
	measure(c)
	measure(B{})
}
