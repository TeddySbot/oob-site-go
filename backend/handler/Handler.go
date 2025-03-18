package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// Routes site général
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/reglement", serveReglement)
	r.HandleFunc("/boutique", serveBoutique)
	r.HandleFunc("/lore", serveLore)
	r.HandleFunc("/liens-utiles", serveLiensUtiles)

	// Routes pour le site Vampire
	r.HandleFunc("/vampire", serveVampire)
	r.HandleFunc("/vampire/reglement", serveReglementVampire)

	// Routes pour le site AD
	r.HandleFunc("/armedemoniaque", serveAD)
	r.HandleFunc("/armedemoniaque/reglement", serveReglementAD)

	// Servir les fichiers statiques (CSS, JS)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("frontend"))))

	return r
}

// Site Général

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/general/index.html")
}

func serveReglement(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/general/reglement.html")
}

func serveBoutique(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/general/boutique.html")
}

func serveLore(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/general/lore.html")
}

func serveLiensUtiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/general/lien.html")
}

// Site Vampire

func serveVampire(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/vampire/vampire.html")
}

func serveReglementVampire(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/vampire/reglement-vampire.html")
}

// Site AD

func serveAD(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/ad/ad.html")
}

func serveReglementAD(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/template/ad/reglement-ad.html")
}
