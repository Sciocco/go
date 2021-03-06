package arm

var Anames = []string{
	"XXX",
	"CALL",
	"CHECKNIL",
	"DATA",
	"DUFFCOPY",
	"DUFFZERO",
	"END",
	"FUNCDATA",
	"GLOBL",
	"JMP",
	"NOP",
	"PCDATA",
	"RET",
	"TEXT",
	"TYPE",
	"UNDEF",
	"USEFIELD",
	"VARDEF",
	"VARKILL",
	"AND",
	"EOR",
	"SUB",
	"RSB",
	"ADD",
	"ADC",
	"SBC",
	"RSC",
	"TST",
	"TEQ",
	"CMP",
	"CMN",
	"ORR",
	"BIC",
	"MVN",
	"BEQ",
	"BNE",
	"BCS",
	"BHS",
	"BCC",
	"BLO",
	"BMI",
	"BPL",
	"BVS",
	"BVC",
	"BHI",
	"BLS",
	"BGE",
	"BLT",
	"BGT",
	"BLE",
	"MOVWD",
	"MOVWF",
	"MOVDW",
	"MOVFW",
	"MOVFD",
	"MOVDF",
	"MOVF",
	"MOVD",
	"CMPF",
	"CMPD",
	"ADDF",
	"ADDD",
	"SUBF",
	"SUBD",
	"MULF",
	"MULD",
	"DIVF",
	"DIVD",
	"SQRTF",
	"SQRTD",
	"ABSF",
	"ABSD",
	"SRL",
	"SRA",
	"SLL",
	"MULU",
	"DIVU",
	"MUL",
	"DIV",
	"MOD",
	"MODU",
	"MOVB",
	"MOVBS",
	"MOVBU",
	"MOVH",
	"MOVHS",
	"MOVHU",
	"MOVW",
	"MOVM",
	"SWPBU",
	"SWPW",
	"RFE",
	"SWI",
	"MULA",
	"WORD",
	"BCASE",
	"CASE",
	"MULL",
	"MULAL",
	"MULLU",
	"MULALU",
	"BX",
	"BXRET",
	"DWORD",
	"LDREX",
	"STREX",
	"LDREXD",
	"STREXD",
	"PLD",
	"CLZ",
	"MULWT",
	"MULWB",
	"MULAWT",
	"MULAWB",
	"DATABUNDLE",
	"DATABUNDLEEND",
	"MRC",
	"LAST",
}

var cnames5 = []string{
	"NONE",
	"REG",
	"REGREG",
	"REGREG2",
	"SHIFT",
	"FREG",
	"PSR",
	"FCR",
	"RCON",
	"NCON",
	"SCON",
	"LCON",
	"LCONADDR",
	"ZFCON",
	"SFCON",
	"LFCON",
	"RACON",
	"LACON",
	"SBRA",
	"LBRA",
	"HAUTO",
	"FAUTO",
	"HFAUTO",
	"SAUTO",
	"LAUTO",
	"HOREG",
	"FOREG",
	"HFOREG",
	"SOREG",
	"ROREG",
	"SROREG",
	"LOREG",
	"PC",
	"SP",
	"HREG",
	"ADDR",
	"TEXTSIZE",
	"GOK",
	"NCLASS",
	"SCOND = (1<<4)-1",
	"SBIT = 1<<4",
	"PBIT = 1<<5",
	"WBIT = 1<<6",
	"FBIT = 1<<7",
	"UBIT = 1<<7",
	"SCOND_XOR = 14",
	"SCOND_EQ = 0 ^ C_SCOND_XOR",
	"SCOND_NE = 1 ^ C_SCOND_XOR",
	"SCOND_HS = 2 ^ C_SCOND_XOR",
	"SCOND_LO = 3 ^ C_SCOND_XOR",
	"SCOND_MI = 4 ^ C_SCOND_XOR",
	"SCOND_PL = 5 ^ C_SCOND_XOR",
	"SCOND_VS = 6 ^ C_SCOND_XOR",
	"SCOND_VC = 7 ^ C_SCOND_XOR",
	"SCOND_HI = 8 ^ C_SCOND_XOR",
	"SCOND_LS = 9 ^ C_SCOND_XOR",
	"SCOND_GE = 10 ^ C_SCOND_XOR",
	"SCOND_LT = 11 ^ C_SCOND_XOR",
	"SCOND_GT = 12 ^ C_SCOND_XOR",
	"SCOND_LE = 13 ^ C_SCOND_XOR",
	"SCOND_NONE = 14 ^ C_SCOND_XOR",
	"SCOND_NV = 15 ^ C_SCOND_XOR",
}
