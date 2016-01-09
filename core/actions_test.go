package core

import (
  "testing"
  "errors"
)

func TestAddInt(t *testing.T) {
  beforeEach();
  expected := "2"
  Push("1");
  Push("1");
  add();
  actual := Pop();
  if expected != actual {
    t.Error("Unexpected answer");
    t.Error(actual);
  }
}

func TestAddString(t *testing.T) {
  beforeEach();
  expected := "Hello World!"
  Push(" World!");
  Push("Hello");
  add();
  actual := Pop();
  if expected != actual {
    t.Error("Unexpected answer");
    t.Error(actual);
  }
}

func TestAddHasRequiredParameters(t *testing.T) {
  beforeEach();

  expected := errors.New("Requires 2 parameters on the stack");
  actual := add();

  if expected.Error() != actual.Error() {
    t.Error("Expecting an error for not enough parameters.");
    t.Error("EXPECT");
    t.Error(expected.Error());
    t.Error("ACTUAL");
    t.Error(actual.Error());
  }
}

func TestPrintHasRequiredParameters(t *testing.T) {
  beforeEach();
  expected := errors.New("Requires 1 parameter on the stack");
  _, actual := print();

  if expected.Error() != actual.Error() {
    t.Error("Expecting an error for not enough parameters.");
    t.Error("EXPECT");
    t.Error(expected.Error());
    t.Error("ACTUAL");
    t.Error(actual.Error());
  }
}

func TestPrint(t *testing.T) {
  beforeEach();
  expected := "Hello World!";

  Push("Hello World!");
  actual, _ := print();

  if expected != actual {
    t.Error("Expecting: "+expected+" but was: "+ actual);
  }
}

func TestSubtract(t *testing.T) {
  beforeEach();
  expected := "1"
  Push("1");
  Push("2");
  subtract();
  actual := Pop();
  if expected != actual {
    t.Error("Expecting: "+expected+" but was: "+ actual);
  }
}

func TestSubtractString(t *testing.T) {
  beforeEach();
  expected := errors.New("Both values should be of type int");
  Push(" World!");
  Push("Hello");
  actual := subtract();
  if expected.Error() != actual.Error() {
    t.Error("Expecting: ",expected.Error()," but was: ",actual.Error());
  }
}
