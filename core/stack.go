package core

var stack = [100]string{};
var size = 0;

func Push(symbol string) {
  stack[size] = symbol;
  size++;
}

func Pop() string{
  size--;
  symbol := stack[size];
  stack[size] = "";
  return symbol;
}

func clearStack() {
  for i :=size; i>0; i--{
    Pop();
  }
}
