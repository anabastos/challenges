/*
Meia Noite = 12:00:00AM / 00:00:00 
Meio Dia = 12:00:00PM / 12:00:00

string: 07:05:45PM
output: 19:05:45


string contains PM 
	hora = string[:2]
	if hora == 12
		hora = 12:00
	hora = hora + 12
	return hora

string contains AM
	hora = string[:2]
	if hora == 12
		hora = 00
	return hora
	

*/
package main
  
import (
  	// "fmt"
	"strings"
  	// "io/ioutil"
  )
  
func MidnightFormat(code string) string {
	index := strings.Index(code, "AM")
    if  index == -1 {
    	return "Not is morning"
	}

	hour := code[0: 2]
	if hour != "12" {
		return "Not is midnight"
	}
	hourFormated := code[0: 8]
	finalHour := "00" + hourFormated[2:8]
 	return finalHour	
}
func To24Hours(code string) string {
	// x := "chars@arefun"

    // i := strings.Index(x, "@")

    // fmt.Println("Index: ", i)
	
 	return ""
 }
// func main() {
//   	fmt.Println("Hello QR Code")
 
//  	qrcode := GenerateQRCode("555-2368")
//  	fmt.Println(qrcode)
// }