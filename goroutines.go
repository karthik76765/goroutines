package main
import(
	"fmt"
	"time"

)
func f(from string){
	for i := 0; i < 3; i++{
		fmt.println(from,":",i)
	} 


}
func main() {
	f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.println(msg)

	}("going")
	time.sleep(time.second)
	fmt.Println("done")
}