package mytest

func Mytest(x float64) float64{
    z := 0.0
    for i:=0;i<10000;i++{
	z -=(z*z - x)/(2*x)
    }
    return z
}

