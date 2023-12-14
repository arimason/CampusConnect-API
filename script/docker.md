``` bash
sudo apt update
sudo apt install docker.io
sudo usermod -aG docker $USER
sudo systemctl start docker
```

``` bash
docker network create minha-rede
```

``` bash
docker run --network=minha-rede --name meu-postgres \
    -e POSTGRES_PASSWORD=root \
    -e POSTGRES_USER=root \
    -e POSTGRES_DB=campus-connect-api \
    -p 5433:5433 \
    -d postgres:16.1 \
    -c 'port=5433'
```

``` bash
docker build -t campus_connect .
```

``` bash
docker run -d --network minha-rede -p 18181:18181 campus_connect
```

``` bash
docker exec -it <nome_conteiner_api> sh
```

``` sh
sh ./script/migration.sh 
```

``` bash
docker exec -it meu-postgres psql -U root -d campus-connect-api -p 5433
```

``` bash
\dt
```