package main

import (
	"net/http"

	"github.com/pillsomi/gymshark/app/controllers"
	"github.com/pillsomi/gymshark/app/handlers"
	"github.com/pillsomi/gymshark/app/storage"
)

func main() {

	store := storage.New()
	controller := controllers.New(store)
	calculateHandler := handlers.CalculateNumberOfBoxes(controller)
	http.HandleFunc("/calculate", calculateHandler)

	getPackageSizesHandler := handlers.GetPackageSizes(controller)
	http.HandleFunc("/package/sizes", getPackageSizesHandler)

	updatePackageSizesHandler := handlers.UpdatePackageSizes(controller)
	http.HandleFunc("/package/sizes/update", updatePackageSizesHandler)

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.ListenAndServe(":8080", nil)
}
