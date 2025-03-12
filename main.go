package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// TransportMethod définit l'interface commune pour tous les moyens de transport.
type TransportMethod interface {
	// Méthode pour livrer un colis
	DeliverPackage(destination string) (string, error)
	// Méthode pour obtenir l'état du transport
	GetStatus() string
}

// Implémentation du camion (Truck)
type Truck struct {
	ID       string
	Capacity int
}

func (t Truck) DeliverPackage(destination string) (string, error) {
	// Simulation d'une livraison lente
	time.Sleep(time.Second * time.Duration(rand.Intn(5)+3))
	return fmt.Sprintf("Truck %s avec une capacité de %d a livré le colis à %s", t.ID, t.Capacity, destination), nil
}

func (t Truck) GetStatus() string {
	return fmt.Sprintf("Truck %s is operational with capacity %d", t.ID, t.Capacity)
}

// Implémentation du drone (Drone)
type Drone struct {
	ID      string
	Battery int
}

func (d Drone) DeliverPackage(destination string) (string, error) {
	// On vérifie la batterie
	if d.Battery < 20 {
		return "", errors.New("Drone à court de batterie, la livraison est alors annulée")
	}
	time.Sleep(1 * time.Second)
	return fmt.Sprintf(" Drone %s a livré le colis à %s", d.ID, destination), nil
}

func (d Drone) GetStatus() string {
	return fmt.Sprintf("Drone %s est operationel avec %d%% de batterie", d.ID, d.Battery)
}

// Implémentation du bateau (Boat)
type Boat struct {
	ID      string
	Weather string // ça représente la météo actuelle
}

func (b Boat) DeliverPackage(destination string) (string, error) {
	// Vérification des conditions météo
	if b.Weather == "Storm" {
		return "", errors.New("Tempête détectée donc la livraison est annulée")
	}

	time.Sleep(5 * time.Second)
	return fmt.Sprintf("Boat %s a livré le colis à %s", b.ID, destination), nil
}

func (b Boat) GetStatus() string {
	return fmt.Sprintf("Boat %s est operationel", b.ID)
}

// Fonction principale pour tester
func main() {
	fmt.Println("Système de Gestion de Livraison ")

	truck := Truck{ID: "A8U5", Capacity: 5}
	drone := Drone{ID: "1234N", Battery: 15} //En %
	boat := Boat{ID: "6TD4G", Weather: "Clear"}

	transports := []TransportMethod{truck, drone, boat}
	destinations := []string{"Marseille", "Belgique", "Allemagne"}

	//Tester les livraisons ainsi que les statuts des transports
	for i, transport := range transports {
		result, err := transport.DeliverPackage(destinations[i])
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Succès :", result)
		}
		fmt.Println("Statut du transport :", transport.GetStatus())
	}

}
