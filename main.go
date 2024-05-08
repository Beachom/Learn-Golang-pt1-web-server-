package main
import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); != nil {
		fmt.Fprintf(w, "ParseFrom() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST reuqest successful")
	name := r.FormValue("name")
	address := r,FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

funchelloHandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/hello" {
	http.Error(w, "404 not found", http.SatusNotFound)
	return
}
if r.Method != "GET"{
	http.Error(w, "method is not supported", http.SatusNotFound)
	return
}
fmt.FPrintf(w, "hello!")
}

func main(){
	fileServer := http.FileServer(https.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServer(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}