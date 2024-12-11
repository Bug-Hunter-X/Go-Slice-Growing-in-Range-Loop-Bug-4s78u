func myFuncFixed(a []int) {
  b := make([]int, len(a), 2*len(a)) // Pre-allocate to avoid resizing issues
  copy(b,a)
  for i := range a {
    b = append(b, i)
  }
  a = b //reassign a with the new slice b
} 