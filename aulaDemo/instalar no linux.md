

Pasta no Linux
  /etc/systemd/system

criar arquivo:
  nomedoapp.service

jogar o app em uma pasta em 
  /opt/nomedasuaempresa
  chmod +x nomedoapp

1. Recarregar configurações do systemd:
   sudo systemctl daemon-reload

2. Habilitar o serviço (para iniciar automaticamente no boot):
   sudo systemctl enable nomedoapp.service

3. Iniciar o serviço:
   sudo systemctl start nomedoapp.service

4. Verificar status:
   sudo systemctl status nomedoapp.service

5. Para parar:
   sudo systemctl stop nomedoapp.service

6. Para desabilitar (não iniciar no boot):
   sudo systemctl disable nomedoapp.service

7. Ver logs:
   sudo journalctl -u nomedoapp.service -f