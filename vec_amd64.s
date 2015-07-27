TEXT ·mul(SB),4,$0-48
	MOVUPS v1+0(FP),  X0
	MOVUPS v2+16(FP), X1
	MULPS  X1, X0
	MOVUPS X0, toReturn+32(FP)
	RET

TEXT ·add(SB),4,$0-48
	MOVUPS v1+0(FP),  X0
	MOVUPS v2+16(FP), X1
	ADDPS  X1, X0
	MOVUPS X0, toReturn+32(FP)
	RET
