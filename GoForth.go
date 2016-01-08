package main 

import (
  "fmt"
  "bufio"
  "os"
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

/** Builds a prompt */
func prompt() string {
  return "FORTH> ";
}

func parse(input string) {
  //TODO parse the input
}

/** Loops collecting information until `die` */
func loop() {
  var prompt string = prompt();
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print(prompt);
  for scanner.Scan() {
      var input = scanner.Text();
      var symbols = core.Parse(input);
      fmt.Println(symbols);
      fmt.Print(prompt);
  }

}

/** Starts the FORTH interrupter */
func main() {
  setup();
  loop();
}
