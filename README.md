# ciber-canyas

Stupid project for counting canyes, just for fun. CIFP Francesc de Borja Moll internal meme.

## Development set up

### Requirements

+ docker && docker-compose installed. [Reference](https://docs.docker.com/engine/install/ubuntu/)

### Commands

**Start dev compose**:

```bash
mv .env.example .env
mv dev.docker-compose.yml docker-compose.yml 
docker-compose up -d && docker-compose logs -f
```

**When new go packages or npm modules are needed/added to package.json or go.mod**:

```bash
docker-compose down (frontend || backend) && docker-compose up -d (frontend || backend)
```

### Architecture

![img/canyes-diagram.png](img/canyes-diagram.png)

### Technologies

**Frontend**:

+ Vite
+ Svelte [Reference](https://svelte.dev/docs)

**Backend**:

+ Go [Reference](https://go.dev/doc/)
+ Gin [Reference](https://gin-gonic.com/docs/)

**DDBB**:

+ Postgres
+ Go implementation [Reference](https://blog.logrocket.com/building-simple-app-go-postgresql/)
