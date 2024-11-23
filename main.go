package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

var currentList string

type TodoEntry struct {
	Name        string
	Description string
	Priority    string
	DueDate     string
}

func main() {
	fmt.Println("Starting todo CLI ...")

	args := os.Args[1:]

	if len(args) > 0 {
		selectList(args[0])
	}

	for true {
		acceptCommandFromConsole()
	}

	fmt.Println("Terminating todo CLI ...")
}

func acceptCommandFromConsole() {
	commandName := getCommandInput("Command: ")

	handleCommand(commandName)
}

func handleCommand(command string) {
	switch {
	case strings.EqualFold(command, "init"):
		handleInit()
	case strings.EqualFold(command, "select"):
		handleSelect()
	case strings.EqualFold(command, "exit"):
		os.Exit(0)
	case strings.EqualFold(command, "ls"):
		handleLs()
	case strings.EqualFold(command, "push"):
		handlePush()
	case strings.EqualFold(command, "pop"):
		handlePop()
	default:
		fmt.Println("No such command")
	}

}

func handleInit() {

	var file *os.File

	defer file.Close()

	for file == nil {
		listName := getTextInput("List name: ")
		listPath := getListPath(listName)
		potentialFile, err := os.OpenFile(listPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)

		if err != nil {
			if errors.Is(err, os.ErrExist) {
				fmt.Println("That list already exists.")
				continue
			} else {
				panic(err)
			}
		}
		defer potentialFile.Close()
		file = potentialFile
	}

	file.WriteString(strings.Join(getFields(), ", ") + "\n")

}

func handleSelect() {
	listName := getTextInput("List name: ")
	selectList(listName)
}

func requireSelectList() {
	for !listExists(currentList) {
		fmt.Println("A list must be selected to proceed.")
		handleSelect()
	}
}

func selectList(listName string) {
	if !listExists(listName) {
		fmt.Println("Failed: Could not select list " + listName + " because it does not exist")
		return
	}
	currentList = listName
}

func handleLs() {
	requireSelectList()
	listName := currentList
	list, err := readList(listName)
	if err != nil {
		panic(list)
	}
	for i, item := range list[1:] {
		fmt.Print(i)
		fmt.Print(": ")
		fmt.Println(recordToString(item))
	}
}

func handlePush() {
	requireSelectList()
	listName := currentList
	list, err := readList(listName)
	if err != nil {
		panic(list)
	}
	list = append(list, buildRecord())
	writeList(currentList, list)
}

func handlePop() {
	requireSelectList()
	listName := currentList
	list, err := readList(listName)
	if err != nil {
		panic(list)
	}
	if len(list) <= 1 {
		println("List is empty. Cannot pop.")
		return
	}
	item := list[len(list)-1]
	fmt.Println(recordToString(item))
	list = list[:len(list)-1]
	writeList(currentList, list)
}

func buildRecord() []string {
	name := getTextInput("Name: ")
	desc := getTextInput("Description: ")
	prio := getTextInput("Priority: ")
	date := getTextInput("Date: ")
	return []string{name, desc, prio, date}
}

func getListsDirectory() string {
	exectuable, err := os.Executable()
	if err != nil {
		panic(err)
	}

	executableDirectory := filepath.Dir(exectuable)
	listsDirectory := executableDirectory + string(os.PathListSeparator) + "todoLists"
	return listsDirectory
}

func getListPath(listName string) string {
	return getListsDirectory() + string(os.PathListSeparator) + listName + ".csv"
}

func getFields() []string {
	exampleRow := &TodoEntry{Name: "Example", Description: "Example", Priority: "Example", DueDate: "Example"}
	val := reflect.ValueOf(exampleRow).Elem().Type()
	fieldNames := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fieldNames[i] = val.Field(i).Name
	}
	return fieldNames
}

func getCommandInput(instruction string) string {
	fmt.Println(instruction)
	fmt.Print(currentList + "> ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	trimmedText := strings.TrimSuffix(text, "\n")
	return trimmedText
}
func getTextInput(instruction string) string {
	fmt.Println(instruction)
	fmt.Print(currentList + "> ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// If the command is ever to exit, exit immediately
	trimmedText := strings.TrimSuffix(text, "\n")
	return trimmedText
}

func listExists(listName string) bool {
	return fileExists(getListPath(listName))
}

func fileExists(filePath string) bool {
	potentialFile, err := os.OpenFile(filePath, os.O_RDONLY, 0664)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false
		} else {
			panic(err)
		}
	}
	defer potentialFile.Close()
	return true
}

func recordToString(record []string) string {
	return "[" + record[0] + " (" + record[2] + ")]: " + record[1] + "\n" + record[3]
}

func readList(listName string) ([][]string, error) {
	if !listExists(listName) {
		return nil, os.ErrNotExist
	}
	file, err := os.Open(getListPath(listName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lines, err := csv.NewReader(file).ReadAll()
	return lines, nil
}

func writeList(listName string, list [][]string) {
	if !listExists(listName) {
		panic("Tried to write to list that didn't exist.")
	}
	if len(list) == 0 {
		panic("Corrupted list detected, refused to write.")
	}
	file, err := os.OpenFile(getListPath(listName), os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.WriteAll(list)
}
