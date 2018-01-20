package classical

import "../utilities"

func Encode (ptext string, key int) string {
	ciphertext := make([] rune, len(ptext))

	for i, v := range ptext {
		ciphertext[i] = utilities.GetLetter(v, key)
	}
	return string(ciphertext)
}


func Decode (ctext string, key int) string {
	plaintext := make([] rune, len(ctext))

	for i, v := range ctext {
		plaintext[i] = utilities.GetLetter(v, -key)
	}
	return string(plaintext)
}
