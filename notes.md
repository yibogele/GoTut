
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
### black identifier _
- interface checks
  - type assertions
    ```javascript
    if _, ok := val.(json.Marshaler); ok {
        fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
    }
    ```
  - To guarantee that the implementation is correct, a global declaration using the blank identifier can be used in the package:
    ```javascript
    var _ json.Marshaler = (*RawMessage)(nil)
    ```
### Embedding
- it's easier and more evocative to embed the two interfaces to form the new one
```javascript
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```
Only interfaces can be embedded within interfaces.
- The same basic idea applies to structs
```javascript
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```
The embedded elements are pointers to structs and of course must be initialized to point to valid structs before they can be used. The ReadWriter struct could be written as
```javascript
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
```
but then to promote the methods of the fields and to satisfy the io interfaces, we would also need to provide forwarding methods, like this:
```javascript
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
```
By embedding the structs directly, we avoid this bookkeeping

There's an important way in which embedding differs from subclassing. When we embed a type, the methods of that type become methods of the outer type, but when they are invoked the receiver of the method is the inner type, not the outer one
- Embedding can also be a simple convenience
```javascript
type Job struct {
    Command string
    *log.Logger
}
```
once initialized, we can log to the Job:
```javascript
job.Println("starting now...")
```
The Logger is a regular field of the Job struct, so we can initialize it in the usual way inside the constructor for Job, like this,
```javascript
func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
```
or with a composite literal,
```javascript
job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
```
If we need to refer to an embedded field directly, the type name of the field, ignoring the package qualifier, serves as a field name
```javascript
func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
```
- Embedding types introduces the problem of name conflicts but the rules to resolve them are simple
- there is no problem if a field is added that conflicts with another field in another subtype if neither field is ever used.

### Concurrency
- Do not communicate by sharing memory; instead, share memory by communicating.
>Concurrent programming in many environments is made difficult by the subtleties required to implement correct access to shared variables. Go encourages a different approach in which shared values are passed around on channels and, in fact, never actively shared by separate threads of execution. Only one goroutine has access to the value at any given time. Data races cannot occur, by design. 
