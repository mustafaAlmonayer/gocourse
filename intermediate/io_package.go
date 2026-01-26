package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf[:n]))
}

func wirteToWiter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln(err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe!"))
		pw.Close()
	}()
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(pr)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.String())
}

func ioPackage() {
	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello Reader!"))
	var writer bytes.Buffer
	wirteToWiter(&writer, "Hello Writer!")
	fmt.Println(writer.String())
	bufferExample()
	multiReaderExample()
	pipeExample()
}
