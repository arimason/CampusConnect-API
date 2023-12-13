**Passos para Acessar sua API no Raspberry Pi de Fora da Rede Local**

**1. Configurar um servidor SSH no Raspberry Pi:**
   - Certifique-se de que o SSH está instalado e configurado no Raspberry Pi.
     ```
     sudo apt-get install openssh-server
     ```

**2. Configurar um servidor reverso SSH no Raspberry Pi:**
   - Abra a configuração do SSH no Raspberry Pi:
     ```
     sudo nano /etc/ssh/sshd_config
     ```
   - Encontre a linha `GatewayPorts` e altere para `GatewayPorts yes`.
   - Reinicie o serviço SSH:
     ```
     sudo service ssh restart
     ```

**3. Encontrar o IP público do Raspberry Pi:**
   - Descubra o endereço IP público do seu Raspberry Pi usando:
     ```
     curl ifconfig.me
     ```

**4. Conectar-se ao Raspberry Pi remotamente:**
   - Use o seguinte comando SSH para estabelecer uma conexão reversa:
     ```
     ssh -R 2222:localhost:22 <user>@<raspberry_pi_ip>
     ```

**5. Configurar o redirecionamento de porta no roteador:**
   - Acesse o painel de administração do seu roteador.
   - Procure por uma opção como "Port Forwarding" ou "Virtual Server".
   - Adicione uma regra de redirecionamento de porta para redirecionar a porta desejada (por exemplo, 2222) para o IP interno do Raspberry Pi na porta correspondente (por exemplo, 22 para SSH).

**6. Instalar o `sshuttle` no servidor externo:**
   - No servidor externo, instale o `sshuttle`:
     ```
     sudo apt-get install sshuttle
     ```

**7. Iniciar a VPN usando `sshuttle`:**
   - Execute o seguinte comando no servidor externo:
     ```
     sudo sshuttle -r <user>@<raspberry_pi_ip> 0/0 -vv
     ```

**8. Acessar a API:**
   - Agora, você pode acessar a API do Raspberry Pi fazendo solicitações para o IP público do seu roteador na porta especificada (por exemplo, http://seu_ip_publico:2222).

Lembre-se de considerar a segurança e atualizar o IP público do Raspberry Pi conforme necessário.
