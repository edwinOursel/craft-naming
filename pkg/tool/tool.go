package tool

import (
	"fmt"
	"strconv"
)

func Process() {
	numberOfPrimesToFind := 121
	maxRowNumber := 30
	maxColumnNumber := 4
	var rowOffset int
	var column int
	var isPrime bool
	var primeMultipleIndex int
	primesMultiples := [10]int{}
	numberToTest := 1
	numberOfPrimesFound := 1
	limitSqrtIndex := 2
	limit := 9
	primes := [122]int{}
	primes[1] = 2
	for numberOfPrimesFound < numberOfPrimesToFind {
		for {
			numberToTest += 2
			if numberToTest == limit {
				limitSqrtIndex = limitSqrtIndex + 1
				limit = primes[limitSqrtIndex] * primes[limitSqrtIndex]
				primesMultiples[limitSqrtIndex- 1] = numberToTest
			}
			primeMultipleIndex = 2
			isPrime = true
			for primeMultipleIndex < limitSqrtIndex && isPrime {
				for primesMultiples[primeMultipleIndex] < numberToTest {
					primesMultiples[primeMultipleIndex] += primes[primeMultipleIndex] + primes[primeMultipleIndex]
				}
				if primesMultiples[primeMultipleIndex] == numberToTest {
					isPrime = false
				}
				primeMultipleIndex++
			}
			if isPrime {
				break
			}
		}
		numberOfPrimesFound++
		primes[numberOfPrimesFound] = numberToTest
	}

	pageIndex := 1
	pageOffset := 1
	for pageOffset <= numberOfPrimesToFind {
		fmt.Println("---------------------------------------------")
		fmt.Println("**** The First " + strconv.Itoa(numberOfPrimesToFind) + " Prime numbers # Page " + strconv.Itoa(pageIndex) + " **** ")
		fmt.Println("---------------------------------------------")
		for rowOffset = pageOffset; rowOffset < pageOffset + maxRowNumber; rowOffset++ {
			for column = 0; column < maxColumnNumber; column++ {
				if rowOffset + column * maxRowNumber <= numberOfPrimesToFind {
					fmt.Printf("%12v", primes[rowOffset + column * maxRowNumber])
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
		pageIndex += 1
		pageOffset += maxRowNumber * maxColumnNumber
	}
}
