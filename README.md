# Soutenance GO

## Installation

### Builder l'image docker

```bash
docker build .
```

### Lancer le docker compose

```bash
docker-compose up -d
```

### Lancer le go en dehors du docker

Pour ça il faut commenter le service web dans le docker-compose.yml

Une fois ça fait il faut lancer le go avec la commande suivante (il faut bien s'assurer que le docker compose est bien lancé) :

```bash
PORT=3000 DB_HOST=localhost DB_PORT=3306 DB_USER=test DB_PASSWORD=test DB_NAME=go-project go run src/main.go
```
