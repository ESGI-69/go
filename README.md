# Soutenance GO

## Installation

### Builder l'image docker

```bash
docker-compose build .
```

### Lancer le docker compose

```bash
docker-compose up
```

Cela va lancer le serveur sur le port 3000 et générer les fichier de documentation API.

Ça permet de ne pas push les fichier de documentation API sur le repo pour ne pas avoir de conflit. 

## Utilisation

Rendez-vous sur l'adresse http://localhost:3000

## API Documentation

Reendez-vous sur l'adresse http://localhost:3000/swagger/index.html
