# GoLang API

## Setup

### Clonar Repositorio

```sh
$ git clone https://github.com/srgrcp/dvp-api.git
```

### Docker

Iniciar servicios con `docker-compose`:

```sh
$ docker-compose up -d
```

La migración de la base de datos es automatica y se generarán 3 usuarios de prueba.

Una vez iniciados los servicios se puede acceder a la API a través de `http://localhost:3700`

## Probar con Postman

Dentro del repositorio podemos encontrar un archivo de colección de Postman llamado `DVP-API.postman_collection.json`

## Probar con otra herramienta

* **Crear ticket**: `POST /api/ticket/`
##### Body
```json
{
    "status": "abierto",
    "userId": 2
}
```

* **Actualizar ticket**: `PUT /api/ticket/{id}`
##### Body
```json
{
    "status": "cerrado",
    "userId": 1
}
```

* **Eliminar ticket**: `DELETE /api/ticket/{id}`

* **Obtener ticket**: `GET /api/ticket/{id}`

* **Obtener tickets**: `GET /api/ticket/q?page=0&sort=-id` `sort` negativo ordena de mayor a menor.

* **Obtener tickets con filtro**: `GET /api/ticket/q?page=0&sort=-id&filters=["status","cerrado"]`
