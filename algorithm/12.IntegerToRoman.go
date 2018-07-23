package algorithm

func IntToRoman(num int) string {

	if num < 1 || 3999 < num {
		return ""
	}

	one := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	ten := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hud := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thu := []string{"", "M", "MM", "MMM"}

	res := thu[num/1000] + hud[(num/100)%10] + ten[(num/10)%10] + one[num%10]

	return res
}
