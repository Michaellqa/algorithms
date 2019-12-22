package nychaos

import (
	"fmt"
)

// if arr[i] > i + 2 => fail

// dif = 0
// dif = a[i] - i
// 1 -> 1 apply -1 count++
// next a[i] = i - 1 OK
// next a[i] = i - 1 OK
//
// dif == 2 count+=2
// next a[i] = i - 1 OK
//   dif==1 case
// next a[i] = i - 2 => count+=2 && dif--
//

// Arrays: New Year Chaos
func minimumBribes(q []int32) {
	var (
		dif            int32 = 0
		count          int32 = 0
		corruptedUnits int32 = 0
	)
	for i, v := range q {
		i := int32(i)
		if i-v < -2 {
			fmt.Println("Too chaotic")
			return
		}
		if i-v-dif == 0 {
			if corruptedUnits != 0 {
				corruptedUnits--
			}
			continue
		}
		if i-v-dif == 1 {
			if corruptedUnits == 2 {
				count += 2
			} else {
				count++
			}
			dif++
		}
		if i-v-dif == 2 {

			dif++
		}

		// sufferer  i-v-dif > 0
		safe := i - v - dif
		dif -= safe
	}
}
