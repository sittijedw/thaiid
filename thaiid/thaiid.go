package thaiid

func IsValid(id string) bool {
	if len(id) == 13 {
		multiplyAddResult := 0
		for i := range id {
			if id[i] >= 48 && id[i] <= 57 {
				if i != 12 {
					multiplyAddResult += (int(id[i]) - 48) * (13 - i)
				}
			} else {
				return false
			}
		}
		return (11-multiplyAddResult%11)%10 == int(id[12])-48
	} else {
		return false
	}
}
