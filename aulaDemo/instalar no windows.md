Instalação no Windows

Pasta para o app:
  C:\Program Files\NomeDaSuaEmpresa\NomeDoApp

Instalação como Serviço do Windows:

1. Compilar o executável:
   go build -o nomedoapp.exe main.go

2. Copiar o executável para a pasta do programa:
   copy nomedoapp.exe "C:\Program Files\NomeDaSuaEmpresa\NomeDoApp\"

3. Instalar como serviço (executar como Administrador):
   sc create GoDemoService binPath= "C:\curso\demo.exe" DisplayName= "Go Demo Service" start= auto

4. Iniciar o serviço:
   sc start GoDemoService

5. Parar o serviço:
   sc stop GoDemoService

6. Verificar status:
   sc query GoDemoService

7. Desinstalar o serviço:
   sc delete GoDemoService

8. Configurar para iniciar automaticamente:
   sc config GoDemoService start= auto

9. Ver logs do serviço:
   - Painel de Controle > Ferramentas Administrativas > Visualizador de Eventos
   - Windows Logs > System
   - Ou usar: wevtutil qe System /f:text /q:"*[System[Provider[@Name='GoDemoService']]]"

Alternativa usando PowerShell (como Administrador):
   Get-Service GoDemoService
   Start-Service GoDemoService
   Stop-Service GoDemoService
   Set-Service GoDemoService -StartupType Automatic
