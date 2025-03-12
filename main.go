package main

import (
	"errors"
	"fmt"
	"time"
)

// ===================== INTERFACE TRANSPORT =====================
// Interface commune pour tous les moyens de transport
type TransportMethod interface {
	DeliverPackage(destination string) (string, error) // Livrer un colis
	GetStatus() string                                // Récupérer l'état du transport
}

// ===================== STRUCTURES DES TRANSPORTS =====================
// 🚛 Camion
type Truck struct {
	ID       string
	Capacity int
}

func (t Truck) DeliverPackage(destination string) (string, error) {
	time.Sleep(3 * time.Second) // Simulation d'une livraison lente
	return fmt.Sprintf("Truck %s (capacité %d) a livré le colis à %s", t.ID, t.Capacity, destination), nil
}

func (t Truck) GetStatus() string {
	return "Camion prêt"
}

// 🚁 Drone
type Drone struct {
	ID      string
	Battery int
}

func (d Drone) DeliverPackage(destination string) (string, error) {
	// Vérifie si la batterie est suffisante
	if d.Battery < 20 {
		return "", errors.New("Drone à court de batterie, livraison annulée")
	}
	time.Sleep(1 * time.Second) // Livraison rapide
	return fmt.Sprintf("Drone %s a livré le colis à %s", d.ID, destination), nil
}

func (d Drone) GetStatus() string {
	return "Drone prêt"
}

// 🚢 Bateau
type Boat struct {
	ID      string
	Weather string
}

func (b Boat) DeliverPackage(destination string) (string, error) {
	// Vérifie si la météo permet la navigation
	if b.Weather == "Storm" {
		return "", errors.New("Tempête détectée, livraison annulée")
	}
	time.Sleep(5 * time.Second) // Livraison plus lente
	return fmt.Sprintf("Boat %s a livré le colis à %s", b.ID, destination), nil
}

func (b Boat) GetStatus() string {
	return "Bateau prêt"
}

// ===================== FABRIQUE DE TRANSPORTS =====================
// Crée un moyen de transport en fonction d'un type donné (truck, drone, boat)
func GetTransportMethod(method string) (TransportMethod, error) {
	switch method {
	case "truck":
		return Truck{ID: "T123", Capacity: 10}, nil
	case "drone":
		return Drone{ID: "D456", Battery: 100}, nil
	case "boat":
		return Boat{ID: "B789", Weather: "Clear"}, nil
	default:
		return nil, errors.New("méthode de transport inconnue")
	}
}

// ===================== FONCTION TRACKING DE LIVRAISON =====================
// Fonction qui suit une livraison et envoie le résultat dans un channel
func TrackDelivery(transport TransportMethod, destination string, ch chan string) {
	status, err := transport.DeliverPackage(destination)
	if err != nil {
		ch <- fmt.Sprintf("Échec de la livraison : %v", err) // Envoie l'erreur au channel
		return
	}
	ch <- status // Envoie le succès au channel
}

// ===================== PROGRAMME PRINCIPAL =====================
func main() {
	fmt.Println("Système de Gestion de Livraison")

	// Création des transports via la fabrique avec gestion des erreurs
	truck, err1 := GetTransportMethod("truck")
	drone, err2 := GetTransportMethod("drone")
	boat, err3 := GetTransportMethod("boat")

	// Vérification des erreurs
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Erreur lors de la création des transports :", err1, err2, err3)
		return
	}

	// Création d'un channel avec une capacité définie pour éviter le blocage
	ch := make(chan string, 3)

	// Lancement des livraisons en parallèle avec des goroutines
	go TrackDelivery(truck, "New York", ch)
	go TrackDelivery(drone, "Los Angeles", ch)
	go TrackDelivery(boat, "Paris", ch)

	// Boucle pour récupérer et afficher les résultats des 3 livraisons
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch) // Récupère et affiche un message du channel
	}

	fmt.Println("Toutes les livraisons ont été suivies.")
}
