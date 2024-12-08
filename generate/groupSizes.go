package generate

import (
	"fmt"
	"strings"
)

type sizeCheck struct {
	name        string
	bytesNeeded uint
}

func groupConditionSizes(sc []sizeCheck, conditions []string) (output string) {
	if len(sc) >= 1 {
		mm := map[uint][]string{}

		for _, check := range sc {
			_, ok := mm[check.bytesNeeded]
			if !ok {
				mm[check.bytesNeeded] = []string{check.name}
				continue
			}

			mm[check.bytesNeeded] = append(mm[check.bytesNeeded], check.name)
		}

		grouped := make([]string, len(mm))
		var c int

		for sizeOf, indexList := range mm {
			vars := additionJoin(indexList, sizeOf)

			switch sizeOf {
			case 0:
				panic("invalid size 0 for sizeCheck :(")
			case 1:
				grouped[c] = vars
			default:
				if len(indexList) == 1 {
					grouped[c] = fmt.Sprintf("%d*%s", sizeOf, vars)
				} else {
					grouped[c] = fmt.Sprintf("%d*(%s)", sizeOf, vars)
				}
			}
			c++
		}
		output = strings.Join(grouped, "+")
	}

	bools := strings.Join(conditions, "+")
	if bools != "" {
		if output != "" {
			return output + "+" + bools
		} else {
			return bools
		}
	}
	return output
}

func additionJoin(list []string, sizeOf uint) string {
	if len(list) == 1 {
		return list[0]
	}

	if sizeOf == 1 {
		return strings.Join(list, "+")
	}
	return fmt.Sprintf("(%s)", strings.Join(list, "+"))
}
