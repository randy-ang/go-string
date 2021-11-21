# go-string

some library for simple string utility functions

## Installing

```
go get github.com/randy-ang/go-string
```


## Format
formatted print, using values of map to replace the named arguments in curly brackets

### Example

```
import gostring "github.com/randy-ang/go-string"

func main() {
	mapArgs := map[string]interface{}{
		"test": 123,
	}
	fmt.Println(gostring.Format("go {test}", mapArgs))
}
```
