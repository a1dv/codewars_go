package kata

var decoder = map[rune]int {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func Decode(roman string) int {
    current := decoder[rune(roman[0])]
    if(len(roman) == 0) {
      return 0
    }
    if(len(roman) == 1) {
      return current
    }
    next := decoder[rune(roman[1])]
    if (current < next) {
        if (len(roman) == 2) {
          return next - current
        } else {
          return next - current + Decode(roman[2:])
        }
    }
    return current + Decode (roman[1:])
}
