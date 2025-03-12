package main

import (
	"fmt"
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
	time.Sleep(3 * time.Second)
	return fmt.Sprintf("Truck %s avec une capacité de %d a livré le colis à %s", t.ID, t.Capacity, destination), nil
}

func (t Truck) GetStatus() string {
	return "Camion prêt"
}

// Fonction principale pour tester
func main() {
	fmt.Println("Système de Gestion de Livraison ")

	truck := Truck{ID: "A8U5", Capacity: 5}

	transports := []TransportMethod{truck}
	destinations := []string{"Belgique"}

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
