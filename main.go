package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const WebhookDiscord = "https://discord.com/api/webhooks/1484608662931771563/gL_GQeXBxqeqrLcxx79Qt1cHhpvmh5f9BiXOSQXFr-uQgPYhnAAeVzpLlEAqY27zLg-F"
const SiteAlvo = "https://www.youtube.com"

func enviarDiscord(texto string) {
	corpo := map[string]string{
		"content":  texto,
		"username": "Vigilante de Sites",
	}

	jsonDados, _ := json.Marshal(corpo)
	http.Post(WebhookDiscord, "application/json", bytes.NewBuffer(jsonDados))
}

func checarSite() {
	resposta, erro := http.Get(SiteAlvo)

	if erro != nil || resposta.StatusCode >= 400 {
		enviarDiscord(fmt.Sprintf("🚨 **ALERTA DE QUEDA!** O site %s caiu ou está com problemas!", SiteAlvo))
		fmt.Println("❌ Falha detectada e enviada ao Discord.")
	} else {
		enviarDiscord(fmt.Sprintf("✅ **OK!** O site %s está online e respondendo.", SiteAlvo))
		fmt.Println("✅ Status OK enviado ao Discord.")
	}
}

func main() {
	fmt.Println("🛡️ Monitor iniciado! Pressione Ctrl+C para parar.")

	// 1. Checa a primeira vez imediatamente
	checarSite()

	// 2. Configura o relógio para bater de 1 em 1 minuto
	relogio := time.NewTicker(1 * time.Minute)

	// 3. Loop infinito
	for range relogio.C {
		checarSite()
	}
}
