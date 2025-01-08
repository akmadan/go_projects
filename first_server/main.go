package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	// Serve the static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/math-form/", http.StripPrefix("/math-form/", fs))

	// POST call
	http.HandleFunc("/add-form", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Cannot parse form", http.StatusBadRequest)
			return
		}

		// Get the numbers from the body
		num1Str := r.FormValue("num1")
		num2Str := r.FormValue("num2")

		// Convert the numbers to integers
		num1, err1 := strconv.Atoi(num1Str)
		num2, err2 := strconv.Atoi(num2Str)

		// Check if the conversion is successful
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Write the response
		result := num1 + num2

		// Return the result as JSON
		response := map[string]interface{}{
			"result": result,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Addition of query parameters
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		a := r.FormValue("a")
		b := r.FormValue("b")
		aInt, err1 := strconv.Atoi(a)
		bInt, err2 := strconv.Atoi(b)
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Addition of %d and %d is %d", aInt, bInt, aInt+bInt)
	})

	// Greet
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings")
	})

	// Root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	// Start the server
	fmt.Println("Server is running on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server is not running")
	}
}
