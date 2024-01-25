package types

type None struct{}

//
// T1
//

type T1[Typ1 any] struct {
	V1 Typ1
}

func NewT1[Typ1 any](v1 Typ1) T1[Typ1] {
	return T1[Typ1]{
		V1: v1,
	}
}

//
// T2
//

type T2[Typ1, Typ2 any] struct {
	V1 Typ1
	V2 Typ2
}

func NewT2[Typ1, Typ2 any](v1 Typ1, v2 Typ2) T2[Typ1, Typ2] {
	return T2[Typ1, Typ2]{
		V1: v1,
		V2: v2,
	}
}

//
// T3
//

type T3[Typ1, Typ2, Typ3 any] struct {
	V1 Typ1
	V2 Typ2
	V3 Typ3
}

func NewT3[Typ1, Typ2, Typ3 any](v1 Typ1, v2 Typ2, v3 Typ3) T3[Typ1, Typ2, Typ3] {
	return T3[Typ1, Typ2, Typ3]{
		V1: v1,
		V2: v2,
		V3: v3,
	}
}

//
// T4
//

type T4[Typ1, Typ2, Typ3, Typ4 any] struct {
	V1 Typ1
	V2 Typ2
	V3 Typ3
	V4 Typ4
}

func NewT4[Typ1, Typ2, Typ3, Typ4 any](v1 Typ1, v2 Typ2, v3 Typ3, v4 Typ4) T4[Typ1, Typ2, Typ3, Typ4] {
	return T4[Typ1, Typ2, Typ3, Typ4]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
	}
}

//
// T5
//

type T5[Typ1, Typ2, Typ3, Typ4, Typ5 any] struct {
	V1 Typ1
	V2 Typ2
	V3 Typ3
	V4 Typ4
	V5 Typ5
}

func NewT5[Typ1, Typ2, Typ3, Typ4, Typ5 any](
	v1 Typ1, v2 Typ2, v3 Typ3, v4 Typ4, v5 Typ5,
) T5[Typ1, Typ2, Typ3, Typ4, Typ5] {
	return T5[Typ1, Typ2, Typ3, Typ4, Typ5]{
		V1: v1,
		V2: v2,
		V3: v3,
		V4: v4,
		V5: v5,
	}
}
