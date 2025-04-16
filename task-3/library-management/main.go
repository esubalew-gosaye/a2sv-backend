package main

import (
	"library-management/controllers"
	"library-management/services"
)

func main() {
    // Create a new library instance
    library := services.NewLibrary()
    
    // Create a new controller with the library
    controller := controllers.NewLibraryController(library)
    
    // Run the application
    controller.Run()
}
