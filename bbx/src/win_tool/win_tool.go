package win_tool

import(
	"os"
	"log"
	"bufio"
)

func Wait_input()  {
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
		}
		log.Println("command", command)
	}
}
