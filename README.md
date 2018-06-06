MyGoStarter
===========

Template básico de uma aplicação em Http em GO

## Pré-requisitos
```
sudo apt-get update
sudo apt-get -y upgrade
sudo curl -O https://storage.googleapis.com/golang/go1.9.1.linux-amd64.tar.gz
sudo tar -xvf go1.9.1.linux-amd64.tar.gz
sudo mv go /usr/local
mkdir -p ~/go
mkdir -p ~/word
export GOROOT=$HOME/go
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/work

```
Copie esse projeto para seu diretório $GOPATH/src/github.com/dvriesman


## Rodar para testes
```
go run cmd/mygostarter/*.go
```

## Compilar

Copie esse projeto para seu diretório $GOPATH/src/github.com/dvriesman

```
./build.sh
```

## Gerar a docker
```
docker build --no-cache -t  [user/app] .

```

## Publicar
```
docker login
docker push  [user/app]
```

