package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func tellTime(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		currentTime := time.Now()
		fmt.Fprintf(w, "%0dh%d", currentTime.Hour(), currentTime.Minute())
	default:
		fmt.Print("Erreur")
	}
}
func dice(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dice" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(1000)
		fmt.Fprintf(w, "résultat du d1000 : %03d\n", x)
	default:
		fmt.Print("Erreur")
	}
}
func dices(w http.ResponseWriter, r *http.Request) {
	var hasType string
	hasType = r.URL.Query().Get("type")
	var typeValue int
	fmt.Sscan(hasType, &typeValue)
	if r.URL.Path == "/dices" {
		switch r.Method {
		case http.MethodGet:
			diceArray := [8]int{2, 4, 6, 8, 10, 12, 20, 100}
			rand.Seed(time.Now().UnixNano())
			for i := 0; i < 15; i++ {
				var y = rand.Intn(8)
				x := rand.Intn(diceArray[y])
				if y == 7 {
					fmt.Fprintf(w, "%03d ", x)
				} else if y == 4 || y == 5 || y == 6 {
					fmt.Fprintf(w, "%02d ", x)
				} else {
					fmt.Fprintf(w, "%d ", x)
				}
			}
		case len(hasType) != 0 && (typeValue == 2 || typeValue == 4 || typeValue == 6 || typeValue == 8 || typeValue == 10 || typeValue == 12 || typeValue == 20 || typeValue == 100):
			//jsp comment faire pour si ya des parametres

		default:
			fmt.Print("Erreur")
		}
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

}
func main() {
	http.HandleFunc("/", tellTime)
	http.HandleFunc("/dice", dice)
	http.HandleFunc("/dices", dices)

	http.ListenAndServe(":4567", nil)

}
