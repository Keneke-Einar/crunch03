package utils

func SaveFootprints(oldMap [][]rune, newMap [][]rune) [][]rune {
	for i := range oldMap {
		for j := range oldMap[i] {
			if oldMap[i][j] == '#' && newMap[i][j] == '.' {
				newMap[i][j] = 'o'
			}
			if oldMap[i][j] == 'o' && newMap[i][j] == '.' {
				newMap[i][j] = 'o'
			}
		}
	}
	return newMap
}
