import "net/http"

func main() {
	route()
}

// Handles our basic routes
// http://www.gorillatoolkit.org/pkg/mux
func route() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)

	// Just some basic other examples
	r.HandleFunc("/wallets/create", WalletHandler)
	r.HandleFunc("/wallets/{id:[0-9]+", WalletHandler)
		.Methods("GET")
}
