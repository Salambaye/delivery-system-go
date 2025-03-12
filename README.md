# Gestion de Livraison

Description

Ce projet est un système de gestion des livraisons implémenté en Go. Il permet de simuler des livraisons en utilisant différentes méthodes de transport :

Camion (Truck) : Livraison fiable mais lente, capable de transporter plusieurs colis.

Drone : Livraison rapide, limitée aux courtes distances, avec un risque de batterie faible.

Bateau (Boat) : Grande capacité de transport, mais dépend des conditions météorologiques.

Les livraisons sont exécutées en parallèle grâce aux goroutines et aux channels, permettant ainsi de suivre les statuts des livraisons en temps réel.

## Fonctionnalités

Interface commune TransportMethod pour toutes les méthodes de transport.

Implémentation de trois types de transport : Truck, Drone, Boat.

Une fabrique GetTransportMethod pour instancier dynamiquement une méthode de transport.

Un système de suivi des livraisons avec gestion des erreurs.

Utilisation de goroutines et de channels pour simuler les livraisons en parallèle.

Installation

Assurez-vous d'avoir Go installé sur votre machine.

Clonez ce dépôt :

git clone https://github.com/Salambaye/delivery-system-go.git

Accédez au dossier du projet :

cd delivery-system-go

Exécution

Compilez et exécutez le programme avec la commande suivante :

go run main.go

### Explication du Code

Interface TransportMethod : définit les méthodes DeliverPackage(destination string) et GetStatus().

Implémentations des transports : Chaque struct (Truck, Drone, Boat) implémente cette interface avec des comportements spécifiques.

Fabrique GetTransportMethod : Permet de créer dynamiquement des objets en fonction d'un type donné.

Goroutines et Channels : Les livraisons sont exécutées en parallèle et les statuts sont envoyés via un channel.

Exemple de Sortie

Truck T123 delivered package to New York
Drone D456 delivered package to Los Angeles
Delivery failed: Boat delayed due to bad weather

Améliorations Possibles

Ajout d'autres moyens de transport (train, avion, vélo).

Implémentation d'un système de recharge automatique pour les drones.

Gestion avancée des conditions météorologiques pour les bateaux.

Interface utilisateur pour visualiser le suivi des livraisons en temps réel.

#### Auteurs :
- Rostom MOUADDEB
- Salamata Nourou MBAYE
- Maurice NAHOUNME
- Celaire Idriss OKA
- Khadim Mbacké FALL

Momo

Licence

Ce projet est sous licence MIT.
