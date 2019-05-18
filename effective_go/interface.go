package main

type IntPtr *int

type Intf interface {

}

type Int int

func (ip *Int) doit() {

}

// receiver cant be a pointer
//func (ip IntPtr) do() {
//
//}
// receiver cant be interface
//func (itf Intf) do() {
//
//}

/**
 * created: 2019/5/18 14:42
 * By Will Fan
 */
func main() {
	
}
