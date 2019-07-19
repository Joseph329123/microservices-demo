package main

import (
	"math"
) 

func mean(xs[] float64) float64 {
	total := 0.0
	
	for _, v := range xs {
		total +=v
	}
	
	return total/float64(len(xs))
}

func std_dev(xs[] float64) float64 {
	std := 0.0
	avg := mean(xs)
	
	for j:=0; j<len(xs); j++ {
		std += math.Pow(xs[j] - avg, 2)
	}
	
	std = math.Sqrt(std/float64(len(xs)))
	return std
}

func find_min(xs[] float64) float64 {
	min := xs[0]
	
	for _, v := range xs {
        if (v < min) {
            min = v
        }
	}

	return min
}

func find_max(xs[] float64) float64 {
	max := xs[0]
	
	for _, v := range xs {
        if (v > max) {
            max = v
        }
	}

	return max
}