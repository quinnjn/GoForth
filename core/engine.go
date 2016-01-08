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
  var err error = nil;

  switch symbol {
    case "+":
      err = add();
    case ".":
      symbol, err = print();
      if symbol != "" {
        fmt.Println(symbol);
      }
    case "bye":
      panic("bye!");
    default:
      Push(symbol);
  }

  if err != nil {
    fmt.Print("ERROR: ");
    fmt.Println(err);
  }
}

func cleanup() {
  if r := recover(); r != nil {
    fmt.Println(r);
  }
}

/** Loops collecting information until `die` */
func loop() {
  defer cleanup();
  prompt := prompt();
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print(prompt);
  for scanner.Scan() {
      var input = scanner.Text();
      var symbols = Parse(input);
      for i:=0; i<len(symbols); i++ {
        action(symbols[i]);
      }
      fmt.Print(prompt);
  }
}


func Run(){
  loop()
}
