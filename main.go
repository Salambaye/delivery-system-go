package main

import (
	"errors" // Pour gérer les erreurs
	"fmt"    // Pour afficher du texte dans la console
	"time"   // Pour simuler le temps de livraison
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
	ID       string // Identifiant du camion
	Capacity int    // Capacité en nombre de colis
}

// Méthode de livraison pour un camion
func (t Truck) DeliverPackage(destination string) (string, error) {
	time.Sleep(3 * time.Second) // Simule un temps de transport de 3 secondes
	return fmt.Sprintf("Truck %s (capacité %d) a livré le colis à %s", t.ID, t.Capacity, destination), nil
}

// Retourne l'état du camion
func (t Truck) GetStatus() string {
	return "Camion prêt"
}

// 🚁 Drone
type Drone struct {
	ID      string // Identifiant du drone
	Battery int    // Niveau de batterie (en %)
}

// Méthode de livraison pour un drone
func (d Drone) DeliverPackage(destination string) (string, error) {
	// Vérifie si la batterie est suffisante
	if d.Battery < 20 {
		return "", errors.New("Drone à court de batterie, livraison annulée")
	}
	time.Sleep(1 * time.Second) // Livraison rapide en 1 seconde
	return fmt.Sprintf("Drone %s a livré le colis à %s", d.ID, destination), nil
}

// Retourne l'état du drone
func (d Drone) GetStatus() string {
	return "Drone prêt"
}

// 🚢 Bateau
type Boat struct {
	ID      string // Identifiant du bateau
	Weather string // Météo actuelle ("Clear" ou "Storm")
}

// Méthode de livraison pour un bateau
func (b Boat) DeliverPackage(destination string) (string, error) {
	// Vérifie si la météo permet la navigation
	if b.Weather == "Storm" {
		return "", errors.New("Tempête détectée, livraison annulée")
	}
	time.Sleep(5 * time.Second) // Livraison plus lente en 5 secondes
	return fmt.Sprintf("Boat %s a livré le colis à %s", b.ID, destination), nil
}

// Retourne l'état du bateau
func (b Boat) GetStatus() string {
	return "Bateau prêt"
}

// ===================== FONCTION TRACKING DE LIVRAISON =====================
// Fonction qui suit une livraison et envoie le résultat dans un channel
func TrackDelivery(transport TransportMethod, destination string, ch chan string) {
	// Tente d'effectuer la livraison
	status, err := transport.DeliverPackage(destination)
	if err != nil {
		ch <- fmt.Sprintf("Échec de la livraison : %v", err) // Envoie l'erreur au channel
		return
	}
	ch <- status // Envoie le succès au channel
}

// ===================== PROGRAMME PRINCIPAL =====================
func main() {
	fmt.Println("Système de suivi des livraisons")

	// Création des moyens de transport
	truck := Truck{ID: "A8U5", Capacity: 5}
	drone := Drone{ID: "1234N", Battery: 15}  // Batterie faible (échec attendu)
	boat := Boat{ID: "6TD4G", Weather: "Clear"} // Météo favorable

	// Création d'un channel pour recevoir les résultats des livraisons
	ch := make(chan string, 3) // Channel de type string, avec une capacité de 3

	// Lancer les livraisons en parallèle avec des goroutines
	go TrackDelivery(truck, "New York", ch)
	go TrackDelivery(drone, "Los Angeles", ch)
	go TrackDelivery(boat, "Paris", ch)

	// Boucle pour récupérer et afficher les résultats des 3 livraisons
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch) // Récupère et affiche un message du channel
	}

	fmt.Println("Toutes les livraisons ont été suivies.")
}
