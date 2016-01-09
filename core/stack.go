package core

var useQ1 = true
var queue1 = make(chan string, 100);
var queue2 = make(chan string, 100);

func stackSize() int{
  return len(queue1) + len(queue2);
}

func Push(symbol string) {
  if(useQ1){
    normalizeStack(queue1, queue2, symbol);
  } else {
    normalizeStack(queue2, queue1, symbol);
  }
  useQ1 = !useQ1
}

func normalizeStack(q1 chan string, q2 chan string, symbol string){
  q1 <- symbol
  for len(q2) > 0 {
    temp := <- q2
    q1 <- temp
  }
}

func Pop() string{
  if len(queue1) > 0 {
    return <- queue1;
  } else {
    return <- queue2;
  }
}

func clearStack() {
  for i :=stackSize(); i>0; i--{
    Pop();
  }
}
