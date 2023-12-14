# Imagem Go com a versão 1.18.10
FROM golang:1.18.10-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o arquivo go.mod e go.sum para o diretório de trabalho no contêiner
COPY go.mod .
COPY go.sum .

# Baixa as dependências do Go
RUN go mod download

# Copia o código-fonte para o diretório de trabalho no contêiner
COPY . .

# Compila o código-fonte
RUN go build -o campus_connect cmd/server/main.go

# Exponha a porta em que a aplicação vai rodar
EXPOSE 18181

# Comando para executar a aplicação quando o contêiner iniciar
CMD ["./campus_connect", "-addr", "0.0.0.0:18181"]
