package validation

import("github.com/badoux/checkmail"
"unicode"
"fmt"
)

func Email(Email string)(int){
	err := checkmail.ValidateFormat(Email)
   if err == nil{
	fmt.Println("validation",Email)
	return 1
   }else{
	fmt.Println("validation",Email)
     	return 0
   }
}
func Password(password string)(int){
	var c1=0
	var c2=0
	var c3=0
         for _,char:=range password{
			if unicode.IsDigit(char){
				   c1++
			}else if(unicode.IsLetter(char)){
				c2++
			}else if(unicode.IsPunct(char) || unicode.IsSymbol(char)){
				c3++
			}
		 }
		 if c1>0&& c2>0 && c3>0{
			return 1
		 }else{
			return 0
		 }
}