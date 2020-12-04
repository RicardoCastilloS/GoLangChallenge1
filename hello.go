//Challenge part 1
package main 
import ("fmt"
		"strings"
		"bytes"
		"unicode"
)
func IsLetter(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}
func IsDigit(s string) bool {
    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }
    return true
}
func main(){
	var b bytes.Buffer   
	var counter = 0
	s := strings.Split("Aspiration....com","")
    for i:= 0; i< len(s); i ++ {
		if(IsDigit(s[i]) ||IsLetter(s[i])){
		if (counter+1)%3==0&&counter!=0{
			b.WriteString(strings.ToUpper(s[i]))
		} else {
			b.WriteString(strings.ToLower(s[i]))
		}
		counter ++
	} else{
		b.WriteString(s[i])
	}
	}
	fmt.Print( b.String())
}
