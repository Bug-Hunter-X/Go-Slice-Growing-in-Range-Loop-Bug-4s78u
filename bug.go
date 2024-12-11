func myFunc(a []int) {
  for i := range a {
    a = append(a, i)
  }
}