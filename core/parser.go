package core

import (
  "strings"
)

func Parse(input string) []string {
  var symbols = strings.Fields(input);
  return symbols;
}
