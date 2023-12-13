Configuração do Servidor Reverse Proxy com Nginx

1. Instalar o Nginx:

sudo apt update
sudo apt install nginx

2. Configurar o Nginx como Reverse Proxy:

Edite o arquivo de configuração do Nginx:

sudo nano /etc/nginx/sites-available/default

Adicione a configuração do servidor reverse proxy:

server {
    listen 80;

    server_name SEU_DOMINIO_OU_IP;

    location / {
        proxy_pass http://localhost:SUA_PORTA_LOCAL;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}

Substitua SEU_DOMINIO_OU_IP pelo domínio personalizado ou IP do seu roteador e SUA_PORTA_LOCAL pela porta da sua aplicação.

Salve e reinicie o Nginx:

sudo service nginx restart

3. Atualizando a Aplicação para Ouvir Todas as Interfaces:

Certifique-se de que a aplicação ouve em todas as interfaces. Exemplo em Python com Flask:

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=SUA_PORTA_LOCAL)

4. Automatizando a Atualização do IP (Opcional):

Automatize a atualização do IP usando DuckDNS. Crie um script e agende a execução:

#!/bin/bash

TOKEN="SEU_TOKEN_DO_DUCKDNS"
DOMAIN="SEU_DOMINIO.duckdns.org"

CURRENT_IP=$(curl -s https://ifconfig.me)
echo "Meu IP atual: $CURRENT_IP"

Atualizar o IP no DuckDNS:

curl "https://www.duckdns.org/update?domains=$DOMAIN&token=$TOKEN&ip=$CURRENT_IP"

Agende a execução periódica do script usando cron:

0 * * * * /caminho/para/seu/script.sh

Agora, sua aplicação deve ser acessível fora da sua rede doméstica através do domínio ou IP configurado no Nginx. Certifique-se de ajustar todos os placeholders para suas configurações específicas.
