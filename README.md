# GoForth
Forth interpreter in Go.

## What is Forth?

[Forth](https://en.wikipedia.org/wiki/Forth_(programming_language)) is a stack based programming language. It works by placing `symbols` onto a stack, and `actions` that mutate the `symbols`.

Forth | Stack        | Output
------|--------------|-------
4     | [4]          |
3     | [4, 3]       |
2     | [4, 3, 2]    |
1     | [4, 3, 2, 1] |
+     | [4, 3, 3]    |
*     | [4, 9]       |
-     | [36]         |
.     | []           | 36

## Currently supported actions:
### *

Multiplies two symbols on the stack. 

Examples:

```
$ ./GoForth
GoForth 0.1.0
FORTH> 4 3 * .
12
FORTH> 3 4 * .
12
FORTH> hi 3 * .
hihihi
FORTH> 4 hi * .
hihihihi
FORTH>
```

### +

Adds two symbols on the stack.

Examples:

```
$ ./GoForth
GoForth 0.1.0
FORTH> 4 3 + .
7
FORTH> 3 4 + .
7
FORTH> hi 3 + .
3hi
FORTH> 4 hi + .
hi4
FORTH>
```

### -

Subtracts two symbols on the stack.

Examples:
```
$ ./GoForth
GoForth 0.1.0
FORTH> 4 3 - .
-1
FORTH> 3 4 - .
1
FORTH>
```

### .

Prints out the last symbol on the stack.

Examples:
```
$ ./GoForth
GoForth 0.1.0
FORTH> 4 3 .
3
FORTH> .
4
FORTH> hi .
hi
FORTH>
```

### bye

Exits the interpreter 

Example:
```
$ ./GoForth
GoForth 0.1.0
FORTH> bye
bye!
```
