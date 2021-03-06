package romanNumerals

import "fmt"


func initTranslation() map[string]int {
	var translation = make(map[string]int)
	translation["I"] = 1
	translation["V"] = 5
	translation["X"] = 10
	translation["L"] = 50
	translation["C"] = 100
	translation["D"] = 500
	translation["M"] = 1000
	return translation
}

func toRoman(x int) string {

	switch {
	case x >= 1000:
		return "M" + toRoman(x-1000)
	case x >= 900:
		return "CM" + toRoman(x-900)
	case x >= 500:
		return "D" + toRoman(x-500)
	case x >= 400:
		return "CD" + toRoman(x-400)
	case x >= 100:
		return "C" + toRoman(x-100)
	case x >= 90:
		return "XC" + toRoman(x-90)
	case x >= 50:
		return "L" + toRoman(x-50)
	case x >= 40:
		return "XL" + toRoman(x-40)
	case x >= 10:
		return "X" + toRoman(x-10)
	case x >= 9:
		return "IX" + toRoman(x-9)
	case x >= 5:
		return "V" + toRoman(x-5)
	case x >= 4:
		return "IV" + toRoman(x-4)
	case x >= 1:
		return "I" + toRoman(x-1)
	}
	return ""
}

func toArabic(r string, debug bool) int { 
	translation := initTranslation()
	evalMin := 0
	evalBuffer := ""
	accumulator := 0
	for i, c := range r {

		if debug {
			fmt.Println("DEBUGGING: %s", r)
			fmt.Println("==================")
			fmt.Println("c is %s", string(c))
			fmt.Println("evalBuffer is %s", evalBuffer)
		}

		if len(evalBuffer) > 0 && translation[string(c)] > translation[evalBuffer[len(evalBuffer)-1:]] {
			accumulator = accumulator + translation[string(c)] - eval(evalBuffer)
			evalMin = i+1
			
			if debug {
				fmt.Println("c is greater than the last char in evalBuffer.")
				fmt.Println("eval(evalBuffer) was ", eval(evalBuffer))
				fmt.Println("accumulator is now ", accumulator)
			}



		}
		//accumulator = accumulator + translation[string(c)]
		evalBuffer = r[evalMin:i+1]

		if(debug){
			fmt.Println("evalBuffer now contains ", evalBuffer)
		}

	}
	accumulator += eval(evalBuffer)
	return accumulator
}

func eval(buf string) int {
	translation := initTranslation()
	accumulator := 0
	for _, c := range buf {
		accumulator += translation[string(c)]
	}
	return accumulator
}


func ToArabic(r string) int {
	return toArabic(r, false)
}

func ToArabicDebug(r string, debug bool) int {
	return toArabic(r, debug)
}