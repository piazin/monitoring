package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	monitorings = 3
	delay = 5
)

func main() {
	welcome()

 	for {
 		showMenu()
		command := getSelectedCommand()
		executeSelectedCommand(command)
	}
	// fmt.Printf("O comando recebido foi %d", command)

}

func welcome() {
	name, version := "Lucas", 1.1

	fmt.Println("Olá,", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("\n\n1 - Iniciar o Monitoramento \n2 - Exibir os Logs \n0 - Sair do programa\n ")
}

func getSelectedCommand() int {
	var command int
	
	/* o & faz a referencia ao ponteiro de onde está armazenado o valor da variavel command
	fmt.Println("O poiteiro de var command é", &command)*/
	fmt.Scanf("%d", &command)
	return command
}

func executeSelectedCommand(command int) {
	switch command {
	case 1:
		startMonitoring()
	case 2:
		fmt.Println("Exibindo Logs... \n ")
		printLogs()
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Opção invalida")	
		os.Exit(-1)
	}
} 

func startMonitoring() {
	fmt.Println("-> Iniciando monitoramento\n ")
	// urls := [...]string{""} tipo inferido nesse caso 1

	urls := readUrlsFromFile()

	for i:=0; i < monitorings; i++ {
		for _, url := range urls {
			testUrl(url)
		}
		time.Sleep(time.Second*delay)
	}

	/* for classico
	for i:=0; i < len(urls); i++ {
		response, _ := http.Get(urls[i])
		if response.StatusCode == 200 {
			fmt.Printf("Site: %s foi carregado com sucesso!\n", urls[i])
		} else {
			fmt.Printf("Site %s esta com problemas. Status Code: %d \n", urls[i], response.StatusCode)
		}
	} */
}

func printLogs() {
	file, err := os.ReadFile("./log.txt")
	checkError(err)

	fmt.Println(string(file))
}

func testUrl(url string) {
	response, err := http.Get(url)
	checkError(err)
	if response.StatusCode == 200 {
		fmt.Printf("Site: %s foi carregado com sucesso!\n", url)
		writeLog(url, true)
	} else {
		fmt.Printf("Site %s esta com problemas. Status Code: %d \n", url, response.StatusCode)
		writeLog(url, false)
	}
}

func readUrlsFromFile() []string {
	// data, err := os.ReadFile("./urls.txt")
	file, err := os.Open("./urls.txt")
	checkError(err)

	var urlsReadFromFile []string
	
	reader := bufio.NewReader(file)

	for {
		row, err := reader.ReadString('\n')
		urlsReadFromFile = append(urlsReadFromFile, strings.TrimSpace(row))

		if err == io.EOF {
			break
		}

	}
	
	file.Close()
	return urlsReadFromFile
}

func writeLog(url string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkError(err)

	currentTime := time.Now().Format("02-01-2006 15:04:05")

	file.WriteString(url + " - online: " + strconv.FormatBool(status) + " - datetime: " + currentTime + "\n")
	file.Close()
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

/* 
func loops() {
	fruits := [3]string{"apple", "orange", "banana"}
	for i:= 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}
}*/
/* devolver dois valores em uma função
func returnNameAndAge() (string, int) {
	name := "Lucas"
	age := 20
	return name, age
}
name, age := returnNameAndAge()
fmt.Println(name, age)
*/
