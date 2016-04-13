package main

func main() {
  foo := newMyStruct()
  foo.myMap["bar"] = "buz"

  foo.printValue("bar")
}

type myStruct struct {
  myMap map[string]string
}

func newMyStruct() *myStruct {
  result := myStruct{}
  result.myMap = map[string]string{}

  return &result
}

func (ms *myStruct) printValue(key string) {
  println(ms.myMap[key])
}
