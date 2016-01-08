package core

import (
  "fmt"
  "bufio"
  "os"
)

/** Builds a prompt */
func prompt() string {
  return "FORTH> ";
}

/** Loops collecting information until `die` */
func loop() {
  prompt := prompt();
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print(prompt);
  for scanner.Scan() {
      var input = scanner.Text();
      var symbols = Parse(input);
      fmt.Println(symbols);
      Push("a");
      fmt.Println(Pop());
      fmt.Print(prompt);
  }

}


func Run(){
  loop()
}
