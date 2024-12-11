# Go Slice Growing in Range Loop Bug

This repository demonstrates a subtle bug in Go involving the modification of a slice's length within a range loop. The code appears to simply append values to a slice, but due to how range loops are implemented in Go, it results in an infinite loop.

## Bug Description

The `myFunc` function takes a slice of integers as input and aims to append each index to the slice. However, appending within the loop causes the loop condition to be continuously true, resulting in infinite growth and a runtime panic (eventually).