package file

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	filePath := "1.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("open file %s failed err:%v", filePath, err)
	}
	defer file.Close()
	var buffer = make([]byte, 1024)
	for {
		size, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read file %s failed err:%v", filePath, err)

		}
		fmt.Println(string(buffer[:size]))
	}
}

func TestBufferIo(t *testing.T) {
	filePath := "1.txt"
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("open file %s failed err:%v", filePath, err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	delim := []byte("\n")
	for {
		bytes, err := reader.ReadString(delim[0])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(string(bytes))
	}

}

func TestOsRead(t *testing.T) {
	filePath := "1.txt"
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}

func TestFileWrite(t *testing.T) {
	f, err := os.Create("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write([]byte("a beautiful girl in the room"))
}
