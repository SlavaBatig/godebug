package debug

import (
	"fmt"
	"time"
	"strings"
	"encoding/hex"
	"regexp"
	"strconv"
	"log"
	"os"
)

//Debug wrapper
func Debug(namespace string) (func(messages ...string), func(messages ...string)) {
	format := "2006-01-02 15:04:05.000"

	hexed := hex.EncodeToString([]byte(namespace))
	code := stringToColor(hexed)

	start := "\u001b[1m\u001b[38;5;" + strconv.Itoa(int(code)) + "m"
	end := "\u001b[0m"

	umillisec := time.Now().UnixNano() / int64(time.Millisecond)

	return func(messages ...string) {
		newTime := time.Now().UnixNano() / int64(time.Millisecond)
		difference := newTime - umillisec
		umillisec = newTime
		
		fmt.Println(start+namespace+" ["+time.Now().Format(format)+"]"+end+" "+strings.Join(messages, " ")+start+" +"+strconv.Itoa(int(difference))+"ms"+end)
	},
	func (messages ...string) {
		newTime := time.Now().UnixNano() / int64(time.Millisecond)
		difference := newTime - umillisec
		umillisec = newTime

		log.New(os.Stderr, "", 0).Println(start+namespace+":error ["+time.Now().Format(format)+"]"+end+" "+strings.Join(messages, " ")+start+" +"+strconv.Itoa(int(difference))+"ms"+end)
	}
}

func stringToColor(str string) int64 {
	re := regexp.MustCompile("[0-9]+")
	res := re.FindAllString(str, -1)
	
	sum := int64(0)
	
	for _, char := range res {
		parsed, err := strconv.ParseInt(char, 10, 64)
		if err != nil {
		}
		sum = int64(sum) + parsed
	}

	return sum % 256
}