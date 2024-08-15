package CardsEnum

type CardType uint8

type CardValue uint8

const (
	ONE CardValue = iota + 1
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	SOT = iota + 3
	HORSE
	KING
)

func CardValuesMap() map[string]CardValue {
	return map[string]CardValue{
		"ONE":   ONE,
		"TWO":   TWO,
		"THREE": THREE,
		"FOUR":  FOUR,
		"FIVE":  FIVE,
		"SIX":   SIX,
		"SEVEN": SEVEN,
		"SOT":   SOT,
		"HORSE": HORSE,
		"KING":  KING,
	}
}

func CardValuesMapIndex() map[CardValue]string {
	m := make(map[CardValue]string)
	for k, v := range CardValuesMap() {
		m[v] = k
	}
	return m
}

const (
	GOLD CardType = iota
	SPADE
	CUP
	BAST
)

func CardTypeMap() (m map[string]CardType) {
	m = make(map[string]CardType)
	m["GOLD"] = GOLD
	m["SPADE"] = SPADE
	m["CUP"] = CUP
	m["BAST"] = BAST
	return
}

func CardTypeMapIndex() map[CardType]string {
	m := make(map[CardType]string)
	for k, v := range CardTypeMap() {
		m[v] = k
	}
	return m
}
