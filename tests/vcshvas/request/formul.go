package request

var (
	k     float64 = 10e9
	Ietal         = []float64{10e-10, 10e-9, 10e-8, 10e-7, 10e-6, 10e-5, 10e-4, 10e-3}
	Fetal         = []float64{10e-1, 10e-0, 10e1, 10e2, 10e3, 10e4, 10e5}
)

func Fcalc(I float64) float64 {
	return I * k
}

func Imeas(F float64) float64 {
	return F / k
}

func SigF(Fc float64, Fi float64) float64 {
	return (Fc - Fi) / Fi * 100
}

func SigI(Ic float64, Iet float64) float64 {
	return (Ic - Iet) / Iet * 100
}
