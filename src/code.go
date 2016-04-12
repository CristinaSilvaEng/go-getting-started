package main

func main() {
  if foo := 2; foo == 1 {
    println("bar")
  } else {
    println("buz")
  }

  switch foo := 1; foo {
  case 1:
    println("one")
  case 2:
    println("two")
  }

  for index := 0; index < 5; index++ {
    println(index)
  }

  i := 0
  for {
    i++
    println(i)

    if i > 5 {
      break
    }
  }

  s := []string{"foo", "bar", "buz"}
  for idx, v := range s {
    println(idx, v)
  }

  m := make(map[string]string)
  m["first"] = "foo"
  m["second"] = "second"
  m["third"] = "buz"

  for k, v := range m {
    println(k, v)
  }

}
