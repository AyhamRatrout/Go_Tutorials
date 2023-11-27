package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("Level", "CRITICAL", "Log level to search for.\n Valid options[CRITICAL|DEBUG|INFO|EMERGENCY]")
	flag.Parse()

	file, err := os.Open("./log.txt")
	if err != nil {
		fmt.Printf("Error ocurred while trying to open log file %s\n", err)
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}
