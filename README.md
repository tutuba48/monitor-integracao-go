# 🛡️ Monitor de Integrações Concorrente

Um sistema de Health Check de alta performance desenvolvido em Go (Golang). O script utiliza Goroutines para testar a disponibilidade e o tempo de resposta de múltiplas APIs simultaneamente. Se um site cair, ele dispara um alerta automático no Discord.

## 🚀 Tecnologias Utilizadas
- **Go (Golang)**
- **Goroutines & Channels:** Para execução paralela e assíncrona.
- **Net/HTTP:** Para consumo das APIs e envio de alertas via JSON.
- **Time/Ticker:** Agendamento nativo do ciclo de checagem.

## ⚙️ Como executar
1. Certifique-se de ter o Go instalado.
2. Adicione os links e o Webhook do Discord no `main.go`.
3. Execute: `go run main.go`
