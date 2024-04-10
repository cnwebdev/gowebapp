## Useful Project Notes

### Nil Map  

It is a good practice to always check if the map is not nil before storing data to the Map, to avoid program panic.

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


