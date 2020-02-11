package tool

import (
	"fmt"
	"strconv"
)

func Process() {
	M := 121
	RR := 30
	CC := 4
	var ROFFSET int
	var C int
	var JPRIME bool
	var N int
	MULT := [10]int{}
	J := 1
	K := 1
	ORD := 2
	SQUARE := 9
	P := [122]int{}
	P[1] = 2
	for K < M {
		for {
			J += 2
			if J == SQUARE {
				ORD = ORD + 1
				SQUARE = P[ORD] * P[ORD]
				MULT[ORD - 1] = J
			}
			N = 2
			JPRIME = true
			for N < ORD && JPRIME {
				for MULT[N] < J {
					MULT[N] += P[N] + P[N]
				}
				if MULT[N] == J {
					JPRIME = false
				}
				N++
			}
			if JPRIME {
				break
			}
		}
		K++
		P[K] = J
	}

	PNBR := 1
	POFFSET := 1
	for POFFSET <= M {
		fmt.Println("---------------------------------------------")
		fmt.Println("**** The First " + strconv.Itoa(M) + " Prime numbers # Page " + strconv.Itoa(PNBR) + " **** ")
		fmt.Println("---------------------------------------------")
		for ROFFSET = POFFSET; ROFFSET < POFFSET + RR; ROFFSET++ {
			for C = 0; C < CC; C++ {
				if ROFFSET+C*RR <= M {
					fmt.Printf("%12v", P[ROFFSET+C*RR])
				}
			}
			fmt.Println("")
		}
		fmt.Println("")
		PNBR += 1
		POFFSET += RR * CC
	}
}
