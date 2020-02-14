def main():
	M = 121
	RR = 30
	CC = 4
	ROFFSET = 0
	C = 0
	JPRIME = False
	N = 0
	MULT = [0 for x in range(10)]
	J = 1
	K = 1
	ORD = 2
	SQUARE = 9
	P = [0 for x in range(122)]
	P[1] = 2

	while K < M:
		while True:
			J += 2
			if J == SQUARE:
				ORD = ORD + 1
				SQUARE = P[ORD] * P[ORD]
				MULT[ORD - 1] = J

			N = 2
			JPRIME = True
			while N < ORD and JPRIME:
				while MULT[N] < J:
					MULT[N] += P[N] + P[N]
				if MULT[N] == J:
					JPRIME = False
				N += 1
			if JPRIME:
				break
		K += 1
		P[K] = J

	PNBR = 1
	POFFSET = 1
	while POFFSET <= M:
		print("----------------------------------------------")
		print("**** The First " + str(M) + " Prime numbers # Page " + str(PNBR) + " **** ")
		print("----------------------------------------------")
		ROFFSET = POFFSET
		while ROFFSET < POFFSET + RR:
			C = 0
			while C < CC:
				if ROFFSET+C*RR <= M:
					print("{:11d}".format(P[ROFFSET+C*RR]), end="")
				C += 1
			print("")
			ROFFSET += 1
		print("")
		PNBR += 1
		POFFSET += RR * CC


main()
