package luhn

// IsValid checks whether the given number contains correct checksum digit at the last position
func IsValid(num string) bool {
	if len(num) < 2 {
		return false
	}

	return Checksum(num[0:len(num)-1]) == num[len(num)-1:len(num)]
}

// Checksum returns the given string checksum digit as a single character string
func Checksum(num string) string {
	buf := []byte(num)
	var sum uint

	reverse(buf)

	for i, d := range buf {
		d = d - '0' // convert characters to its numerical values

		if i&0x1 == 0 {
			d <<= 1 // double the value of every second digit

			if d > 9 {
				d -= 9 // sum the digits of the numbers e.g. 11->2, 14->5, 18->9 etc.
			}
		}
		sum += uint(d)
	}

	return string((10-sum%10)%10 + '0') // 10 - last digit and convert number to a digit character
}

func reverse(num []byte) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}
