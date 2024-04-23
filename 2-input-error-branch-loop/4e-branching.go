package main

/*
sistem untuk pengecekan suhu tubuh
jika demam, maka tidak boleh masuk mall
jika tidak demam, check apakah pakai masker atau tidak.
*/
func main() {
	var suhu float32 = 38
	var isMasker bool = true
	if suhu >= 36 && suhu < 38 {
		// ...
		// check pakai masker atau tidak
	} else {
		//tidak boleh masuk
	}

	if suhu <= 36 && suhu >= 38 {
		// tidak boleh masuk
	}
	// jika boleh masuk maka check masker
}
