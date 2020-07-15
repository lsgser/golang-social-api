package main

import(
	"log"
	"os"
	"github.com/joho/godotenv"
	"net/http"
	R "./routes"
)

func init() {
	godotenv.Load()
}

func main(){
	log.Println("Running on port :",os.Getenv("PORT"))
	router := R.NewRouter()
	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"),router))	
}
