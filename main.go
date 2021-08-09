package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	monitoramentos = 3
	delay          = 5
)

func main() {

	for {
		showMenu()
		option := getOption()

		switch option {
		case 1:
			fmt.Println("Monitorando...")
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Opção não reconhecida")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func getOption() int {
	var option int
	fmt.Scan(&option)
	return option
}

func startMonitoring() {

	sites := readWebsitesFromFile()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			response, err := http.Get(site)
			if err != nil {
				fmt.Println("Site não disponível.", err)
			}

			if response.StatusCode == http.StatusOK {
				fmt.Println("Site", site, "foi carregado com sucesso!")
			} else {
				fmt.Println("Site", site, "está com problemas. Status Code:", response.StatusCode)
			}
		}
		time.Sleep(delay * time.Second)
	}

}

func readWebsitesFromFile() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("erro ao ler o arquivo:", err)
	}
	reader := bufio.NewReader(arquivo)

	var site string
	for {
		site, err = reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		sites = append(sites, strings.TrimSpace(site))
	}

	fmt.Println(sites)

	arquivo.Close()

	return sites
}
