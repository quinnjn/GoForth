package main 

import (
  "fmt"
  "github.com/quinnjn/GoForth/core"
)

/** Version of GoForth */
const VERSION string = "0.0.0";

/** Prints the MOTD */
func printMotd() {
  fmt.Println("GoForth " + VERSION);
}

/** Setup when the program begins */
func setup() {
  printMotd();
}
func parse(input string) {
  //TODO parse the input
}

/** Starts the FORTH interrupter */
func main() {
  setup();
  core.Run();
}
