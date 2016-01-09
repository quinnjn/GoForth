package core

import (
  "strconv"
  "errors"
)

func add() error {
  if stackSize() < 2 {
    return errors.New("Requires 2 parameters on the stack");
  }

  firstString := Pop();
  secondString := Pop();

  first, firstErr := strconv.Atoi(firstString);
  second, secondErr := strconv.Atoi(secondString);

  if(firstErr != nil || secondErr != nil){
    Push(firstString + secondString);
  } else {
    value := first + second;
    valueString := strconv.Itoa(value);
    Push(valueString);
  }

  return nil;
}

func multiply() error {
  if stackSize() < 2 {
    return errors.New("Requires 2 parameters on the stack");
  }

  firstString := Pop();
  secondString := Pop();

  first, firstErr := strconv.Atoi(firstString);
  second, secondErr := strconv.Atoi(secondString);

  if(firstErr == nil && secondErr == nil){ // Both are digits
    Push(strconv.Itoa(first * second));
  } else if(firstErr == nil && secondErr != nil){ // First is digit, second is string
    if(first > 0){
      Push(secondString);
      first--;
    }
    for i:=0; i<first; i++ {
      Push(secondString);
      add();
    }
  } else if(firstErr != nil && secondErr == nil){ // First is string, second is digit
    if(second > 0){
      Push(firstString);
      second--;
    }
    for i:=0; i<second; i++ {
      Push(firstString);
      add();
    }
  } else { // Both are string.
    Push(secondString);
    Push(firstString);
    return errors.New("Requires at least one Int.");
  }

  return nil;
}

func subtract() error {
  if stackSize() < 2 {
    return errors.New("Requires 2 parameters on the stack");
  }

  firstString := Pop();
  secondString := Pop();

  first, firstErr := strconv.Atoi(firstString);
  second, secondErr := strconv.Atoi(secondString);

  if(firstErr != nil || secondErr != nil){
    Push(secondString);
    Push(firstString);
    return errors.New("Both values should be of type int");
  } else {
    value := first - second;
    valueString := strconv.Itoa(value);
    Push(valueString);
  }

  return nil;
}

func print() (string, error) {
  if stackSize() < 1 {
    return "", errors.New("Requires 1 parameter on the stack");
  }

  return Pop(), nil;
}
