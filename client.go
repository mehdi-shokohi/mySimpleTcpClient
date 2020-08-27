package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

// Write to Socket Con
func Write(conn net.Conn, content string) (int, error) {
	writer := bufio.NewWriter(conn)
	number, err := writer.WriteString(content)
	if err == nil {
		err = writer.Flush()
	}
	return number, err
}

// Messages Delimeter
const (
	QuitSign = "/quit!"
)

func main() {

	var addr = flag.String("addr", "", "The address to Server.")
	var port = flag.Int("port", 8000, " default is 8000.")
	flag.Parse()
	conn, err := net.DialTimeout("tcp", (*addr)+":"+strconv.Itoa(*port), time.Millisecond*6000)
	if err != nil {
		log.Printf("Sender: DialTimeout Error: %s\n", err)
		os.Exit(1)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		stdin, _ := reader.ReadString('\n')
		fmt.Println(stdin)
		sendText := fmt.Sprintf("%d\n%s%s\n", (len(stdin) + len(QuitSign)), stdin[:len(stdin)-1], QuitSign)

		fmt.Println(sendText)
		num, err := Write(conn, sendText)
		if err != nil {
			log.Printf("Sender: Write Error: %s\n", err)
		}
		log.Printf("Sender: Wrote %d byte(s)\n", num)

	}
}
