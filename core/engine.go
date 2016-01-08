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

func action(symbol string) {
  switch symbol {
    case "+":
      add();
    default:
      Push(symbol);
  }
}

/** Loops collecting information until `die` */
func loop() {
  prompt := prompt();
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print(prompt);
  for scanner.Scan() {
      var input = scanner.Text();
      var symbols = Parse(input);
      for i:=0; i<len(symbols); i++ {
        action(symbols[i]);
      }
      fmt.Println(Pop());
      fmt.Println(stack);
      fmt.Print(prompt);
  }

}


func Run(){
  loop()
}
