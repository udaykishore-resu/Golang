package bestpractices

import (
	"fmt"
	"os"
)

type Writer interface {
	Write([]byte) (int, error)
}

type FileWriter struct {
	file *os.File
}

func (fw *FileWriter) Write(data []byte) (int, error) {
	return fw.file.Write(data)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	return fmt.Fprint(os.Stdout, string(data))
}

func WriteData(w Writer, data string) {
	w.Write([]byte(data))
}

func DisplayInterfacesFlexibility() {
	file, err := os.Create("output.txt")

	if err != nil {
		fmt.Println("error in file creating", err)
		return
	}

	defer file.Close()

	fileWriter := &FileWriter{file: file}
	WriteData(fileWriter, "I am file writer")

	consoleWriter := ConsoleWriter{}
	WriteData(consoleWriter, "I am console writer")

}
