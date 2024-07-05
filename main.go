package main

import (
	"github.com/Simnek/web-go/initalize"
)

func main() {

	initalize.FastInitialize()

	//listenAddr := flag.String("listenaddr", ":8080", "Address for http listening")
	//flag.Parse()
	//
	//http.HandleFunc("/user", api.HandleGetUser)
	//http.HandleFunc("/customer", api.HandleGetCustomer)
	//http.HandleFunc("/user/create", api.HandlePostUserWithCORS)
	//
	//http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/about", aboutHandler)
	//http.HandleFunc("/contact", contactHandler)
	//
	//http.ListenAndServe(*listenAddr, nil)

	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("Unable to load .env")
	//}
	//
	//args := os.Args
	//if len(args) < 1 {
	//	log.Fatalln("Include the command to run. Commands available: initialize")
	//}
	//arg := strings.ToLower(args[1])
	//switch arg {
	//case "initialize":
	//	initalize.Initialize()
	//}

	//bench.FastFor()

}
