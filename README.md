# Leven - Fast implementation of Levenstein distance written in Go
This implementation have equal or fastest implementation when using:
- Prefixed words
- Postfixed words
- Fullscan

# Installation
```
go get github.com/m1ome/leven
```

# Usage
```
package main

import (
    "fmt"
    "github.com/m1ome/leven"
)

func main() {
    s1 := "kitten"
    s2 := "sitten"
    fmt.Printf("The distance between %v and %v is %v\n", s1, s2, leven.Distance(s1, s2))
    // The distance between kitten and sitten is 1
}
```

# Behchmarks
```
BenchmarkDefaultM1omeDistance-8                   	 5000000	       255 ns/op
BenchmarkDefaultAgextLevenshteinDistance-8        	 5000000	       261 ns/op
BenchmarkDefaultAgnivadeLevenshteinDistance-8     	 5000000	       325 ns/op
BenchmarkDefaultArbovmLevenshteinDistance-8       	 5000000	       238 ns/op

BenchmarkPrefixedM1omeDistance-8                  	 5000000	       357 ns/op
BenchmarkPrefixedAgextLevenshteinDistance-8       	 5000000	       385 ns/op
BenchmarkPrefixedAgnivadeLevenshteinDistance-8    	 2000000	       981 ns/op
BenchmarkPrefixedArbovmLevenshteinDistance-8      	 2000000	       741 ns/op

BenchmarkPostfixedM1omeDistance-8                 	 5000000	       372 ns/op
BenchmarkPostfixedAgextLevenshteinDistance-8      	 5000000	       396 ns/op
BenchmarkPostfixedAgnivadeLevenshteinDistance-8   	 1000000	      1186 ns/op
BenchmarkPostfixedArbovmLevenshteinDistance-8     	 2000000	       815 ns/op

BenchmarkFullscanM1omeDistance-8                  	10000000	       231 ns/op
BenchmarkFullscanAgextLevenshteinDistance-8       	 5000000	       245 ns/op
BenchmarkFullscanAgnivadeLevenshteinDistance-8    	 5000000	       337 ns/op
BenchmarkFullscanArbovmLevenshteinDistance-8      	10000000	       227 ns/op
```
