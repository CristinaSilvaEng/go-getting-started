package main

func main() {
  addFunc := func(terms ...int) (numTerms int, sum int) {
    for _, term := range terms {
      sum += term
    }
    numTerms = len(terms)
    return
  }

  numTerms, sum := addFunc(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
  println("Added:", numTerms, "terms for a total of", sum)
}
