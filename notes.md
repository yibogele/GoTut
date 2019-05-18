
### Interfaces and other types
- A type can implement multiple interfaces
```javascript
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
    s = s.Copy() // Make a copy; don't overwrite argument.
    sort.Sort(s)
    str := "["
    for i, elem := range s { // Loop is O(N²); will fix that in next example.
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```
- It's an idiom in Go programs to convert the type of an expression to access a different set of methods
```javascript
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```
- Type switches are a form of conversion: they take an interface and, for each case in the switch, in a sense convert it to the type of that case
```javascript
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```
- A type assertion takes an interface value and extracts from it a value of the specified explicit type
```javascript
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}
```
- pointer and interface can not have methods
```javascript
type IntPtr *int

// receiver cant be a pointer
//func (ip IntPtr) do() {
//
//}

type Intf interface {

}

// receiver cant be interface
//func (itf Intf) do() {
//
//}

type Int int

// this is ok
func (ip *Int) doit() {

}
```
- function can have method
```javascript
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
```
