package main

func algoFindLyric(param int) string {
	if param < 1 {
		return result
	}else {
		result += lyric[param] + " and "
		algoFindLyric(param -1)
	}

	return "This is " + result[:len(result)-5]
}
