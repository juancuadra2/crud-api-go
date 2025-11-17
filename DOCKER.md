# Instrucciones de Docker

## 🐳 Construcción y Ejecución con Docker

### Opción 1: Usar Docker Compose (Recomendado)

Esta opción levanta tanto la aplicación como MySQL automáticamente:

```bash
# Construir y levantar todos los servicios
docker-compose up -d

# Ver los logs
docker-compose logs -f

# Detener los servicios
docker-compose down

# Detener y eliminar volúmenes (limpia la base de datos)
docker-compose down -v
```

### Opción 2: Solo construir la imagen Docker

Si ya tienes MySQL ejecutándose localmente:

```bash
# Construir la imagen
docker build -t crud-api-go .

# Ejecutar el contenedor
docker run -d -p 8080:8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=3306 \
  -e DB_USER=jcuadrado \
  -e DB_PASS=jcuadrado \
  -e DB_NAME=go_lab \
  -e SERVER_PORT=8080 \
  --name crud-api-go \
  crud-api-go
```

## 📦 Variables de Entorno

Las siguientes variables de entorno pueden ser configuradas:

| Variable | Descripción | Valor por Defecto |
|----------|-------------|-------------------|
| DB_HOST | Host de la base de datos | localhost |
| DB_PORT | Puerto de MySQL | 3306 |
| DB_USER | Usuario de la base de datos | jcuadrado |
| DB_PASS | Contraseña de la base de datos | jcuadrado |
| DB_NAME | Nombre de la base de datos | go_lab |
| SERVER_PORT | Puerto del servidor API | 8080 |

## 🔍 Comandos Útiles

```bash
# Ver contenedores en ejecución
docker ps

# Ver logs de la API
docker logs crud-api-go -f

# Ver logs de MySQL
docker logs crud-api-mysql -f

# Acceder al contenedor de la API
docker exec -it crud-api-go sh

# Acceder a MySQL
docker exec -it crud-api-mysql mysql -u jcuadrado -p

# Reconstruir la imagen después de cambios en el código
docker-compose up -d --build
```

## 🧪 Probar la API

Una vez que los contenedores estén ejecutándose, puedes probar la API en:

```
http://localhost:8080
```

Por ejemplo, para listar usuarios:

```bash
curl http://localhost:8080/users
```

## 📝 Notas

- El Dockerfile usa una construcción multi-etapa para optimizar el tamaño de la imagen final
- La imagen final es muy ligera (~15MB) gracias al uso de Alpine Linux
- MySQL incluye un healthcheck para asegurar que la API espere hasta que la base de datos esté lista
- Los datos de MySQL se persisten en un volumen Docker llamado `mysql_data`

