## Useful Project Notes

### Go Basic Data Type  

**The error type**  
Error is a special data type for representing error conditions and error messages, this means Go treats errors as values.
Go comes with the errors package in the standard library, some of the ways to return a custom error is to use **errors.New("error message")**.
To format error messages use **fmt.Errorf()** which allows the creation of the formatted message and returns the error value just like **errors.New()**.   

```go
if err != nil {
   return errors.New("error message")
}
```
or using **fmt.Errorf()  
```go
if err != nil {
   return fmt.Errorf("error message") 
}
```  

Custom error message with errors.New()  
```go
func check(a, b int) error {
   if a == 0 && b == 0 {
      return errorw.New("this is a custom error message")
   }
   return nil
}
```  

Custom error message with fmt.Errorf()
```go 
func formatedError(a, b int) error {
   if a == 0 && b == 0 {
      return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid() )
   }
   return nil
}
```

### Nil Map  

It is a good practice to always **check** if the map is not nil before storing data to the Map, to avoid program panic.

```Go 
if aMap == nil {
   //... do something
}
``` 

### Iterate over map  

Iterate over map works in go, but the output will be in random order (by design)

```Go
for key, val := range aMap {
   fmt.Printf("Key: %v, value: %v\n", key, val)
}
``` 

### Struct Data Type 

Create a new data struct using the type keyword   
```go 
type Person struct {
   Name string 
   SurName string
   BirthYear int
   Age int
}
```
When initializing an object without any data assigned to the struct fields, Go will assign a default value to each field depending on its data type.   
- string default = "" // empty string
- int, float will be = 0 
- bool will be = nil
  
**Using the new keyword** 
The new keyword allows the creation of a new structure, it works on all data types except channel and map.

Example:
```go
p := new(Person) // return a new pointer to Person struct 

