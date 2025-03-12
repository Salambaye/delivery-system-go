package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Interface pour les moyens de transport
type TransportMethod interface {
	DeliverPackage(destination string) (string, error) // Livrer un colis
	GetStatus() string                                // État du transport
}

// Implémentation du camion (Truck)
type Truck struct {
	ID       string
	Capacity int
}

func (t Truck) DeliverPackage(destination string) (string, error) {
	time.Sleep(3 * time.Second) // Simule une livraison lente
	return fmt.Sprintf("Truck %s (capacité %d) a livré le colis à %s", t.ID, t.Capacity, destination), nil
}

func (t Truck) GetStatus() string {
	return "Camion prêt"
}

// Implémentation du drone (Drone)
type Drone struct {
	ID      string
	Battery int
}

func (d Drone) DeliverPackage(destination string) (string, error) {
	if d.Battery < 20 {
		return "", errors.New("Drone à court de batterie, livraison annulée")
	}
	time.Sleep(1 * time.Second) // Livraison rapide
	return fmt.Sprintf("Drone %s a livré le colis à %s", d.ID, destination), nil
}

func (d Drone) GetStatus() string {
	return "Drone prêt"
}

// Implémentation du bateau (Boat)
type Boat struct {
	ID      string
	Weather string
}

func (b Boat) DeliverPackage(destination string) (string, error) {
	if b.Weather == "Storm" {
		return "", errors.New("Tempête détectée, livraison annulée")
	}
	time.Sleep(5 * time.Second) // Livraison plus lente
	return fmt.Sprintf("Boat %s a livré le colis à %s", b.ID, destination), nil
}

func (b Boat) GetStatus() string {
	return "Bateau prêt"
}

// Fonction principale
func main() {
	fmt.Println("Système de Gestion de Livraison")

	// Liste des moyens de transport
	transports := []TransportMethod{
		Truck{ID: "A8U5", Capacity: 5},
		Drone{ID: "1234N", Battery: 15},
		Boat{ID: "6TD4G", Weather: "Clear"},
	}

	// Destinations des livraisons
	destinations := []string{"Marseille", "Belgique", "Allemagne"}

	// WaitGroup pour gérer l'attente des goroutines
	var wg sync.WaitGroup

	// Channels pour stocker les résultats et erreurs
	results := make(chan string, len(transports))
	errorsChan := make(chan string, len(transports))

	// Lancement des livraisons en parallèle
	for i, transport := range transports {
		wg.Add(1)

		go func(t TransportMethod, dest string) {
			defer wg.Done()

			// Exécuter la livraison
			result, err := t.DeliverPackage(dest)
			if err != nil {
				errorsChan <- fmt.Sprintf("Erreur : %s", err)
			} else {
				results <- result
			}
		}(transport, destinations[i]) // Passage des variables par copie
	}

	// Goroutine pour fermer les channels après exécution
	go func() {
		wg.Wait()
		close(results)
		close(errorsChan)
	}()

	// Lecture des résultats
	for res := range results {
		fmt.Println(res)
	}
	for err := range errorsChan {
		fmt.Println(err)
	}

	fmt.Println("Toutes les livraisons ont été traitées.")
}
