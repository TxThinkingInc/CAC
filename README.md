# CAC

> **C**ommand-line **A**rguments **C**onfiguration

A configuration file format that is the same as command-line arguments.

## Specification

1. Write all the arguments into a file, and then run the command with the filename. The effect will be the same
2. Supports multiple arguments on separate lines
3. Ignore blank lines and comments lines starting with `#`

## Example

### The original command and arguments

```
command argument1 argument2 argument3
```

### Write all arguments to CAC file

```
argument1 argument2 argument3
```
or
```
argument1 argument2
argument3
```
or
```
# this is comment
argument1
# this is comment too
argument2
argument3
```

### Run command with CAC file

```
command /path/to/cac
```

## Reference implementation

PR your implementation here
