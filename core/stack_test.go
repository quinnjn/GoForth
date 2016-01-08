package core

import (
  "testing"
)

func beforeEach() {
  clearStack();
}

func TestPush(t *testing.T){
  beforeEach();
  expected := [100]string{"a"};
  Push("a");
  actual := stack
  if actual != expected {
    t.Error("Test failed")
  }
}

func TestPop(t *testing.T) {
  beforeEach();
  expectedStack := [100]string{};
  expectedSymbol := "a";

  Push("a");
  actualSymbol := Pop();
  actualStack := stack;

  if expectedStack != actualStack {
    t.Error("Stack is different");
    t.Error(actualStack);
  }

  if expectedSymbol != actualSymbol {
    t.Error("Symbol differs");
    t.Error(actualSymbol);
  }
}

func TestClear(t *testing.T) {
  beforeEach();
  expectedStack := [100]string{};
  Push("a");
  Push("b");
  Push("c");
  Push("d");
  clearStack();
  actualStack := stack
  if expectedStack != actualStack {
    t.Error("Stack should be clear");
    t.Error(actualStack);
  }
}
