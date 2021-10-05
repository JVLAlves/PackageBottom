package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	functions "github.com/JVLAlves/DinamizeSnipeitDataQuest/Utilities/Functions"
	path "github.com/JVLAlves/PackageBottom/PATH"
)

const (
	funcreg = `(?:^\s*func)\s([[:alnum:]]+)`
	methreg = `(?:^\s*func)\s(?:\([a-z]+\s[[:graph:]]+)\s([[:alnum:]]+)`
)

var (
	File     string
	Location string
)

func main() {

	fmt.Print("Type the directory name: ")
	fmt.Scan(&Location)
	switch Location {

	case "personal", "Personal", "PERSONAL":
		Location = path.PERSONAL
	case "forks", "Forks", "FORKS":
		Location = path.FORK
	case "dataquest", "Dataquest", "DataQuest", "DATAQUEST":
		Location = path.DATAQUEST
	case "datamission", "Datamission", "DataMission", "DATAMISSION":
		Location = path.DATAMISSION
	case "datacenter", "Datacenter", "DataCenter", "DATACENTER":
		Location = path.DATACENTER
	case "DinamizeEMKT-FeedBack", "Dinamizeemktfeedback", "DinamizeEmktFeedback", "DinamizeEmkt-Feedback", "Dinamizeemkt-feedback", "DINAMIZEEMKTFEEDBACK", "DINAMIZEEMKT-FEEDBACK":
		Location = path.DINAMIZEEMKTFEEDBACK
	case "Utilities", "utilities", "UTILITIES":
		Location = path.UTILITIES
	case "SnipeMethods", "Snipemethods", "SNIPEMETHODS":
		Location = path.UTILITIES + "/" + "SnipeMethods"
	case "Functions", "functions", "FUNCTIONS":
		Location = path.UTILITIES + "/" + "Functions"
	}

	fmt.Println(Location)
	fmt.Print("Type the file name: ")
	fmt.Scanln(&File)
	filename := Location + "/" + File
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	f, err := os.OpenFile(File+".txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	fileScanner := bufio.NewScanner(file)

	if err != nil {
		fmt.Println("Error opening or creating file:", err)
	}

	for fileScanner.Scan() {
		FuncIn := functions.RegexThis(funcreg, fileScanner.Text())
		MethIn := functions.RegexThis(methreg, fileScanner.Text())

		if fileScanner.Err() != nil {
			log.Fatalf("Erro SCAN: %v", fileScanner.Err().Error())
		}
		if FuncIn != "" {
			fmt.Fprintln(f, FuncIn)
		}
		if MethIn != "" {
			re := regexp.MustCompile(`func`)
			MethIn := re.ReplaceAllString(MethIn, "method")
			fmt.Fprintln(f, MethIn)

		}
	}

}
