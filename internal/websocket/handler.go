package websocket

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

type WebSocketHandler struct {
    clients   map[*websocket.Conn]bool
    Broadcast chan Message
    upgrader  websocket.Upgrader
}

func NewWebSocketHandler() *WebSocketHandler {
    return &WebSocketHandler{
        clients:   make(map[*websocket.Conn]bool),
        Broadcast: make(chan Message),
        upgrader: websocket.Upgrader{
            CheckOrigin: func(r *http.Request) bool {
                return true
            },
        },
    }
}

func (h *WebSocketHandler) HandleConnections(w http.ResponseWriter, r *http.Request) {
    // Log: Nueva conexión WebSocket
    log.Println("Nueva conexión WebSocket establecida")

    ws, err := h.upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatalf("Error al actualizar la conexión a WebSocket: %v", err)
    }
    defer ws.Close()

    // Registra al cliente
    h.clients[ws] = true
    log.Printf("Cliente registrado. Total de clientes conectados: %d", len(h.clients))

    for {
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            // Log: Error al leer el mensaje del cliente
            log.Printf("Error al leer el mensaje del cliente: %v", err)
            delete(h.clients, ws)
            log.Printf("Cliente desconectado. Total de clientes conectados: %d", len(h.clients))
            break
        }

        // Log: Mensaje recibido del cliente
        log.Printf("Mensaje recibido del cliente: %+v", msg)

        // Envía el mensaje al canal Broadcast
        h.Broadcast <- msg
        log.Println("Mensaje enviado al canal Broadcast")
    }
}

func (h *WebSocketHandler) HandleMessages() {
    for {
        // Espera un mensaje del canal Broadcast
        msg := <-h.Broadcast
        log.Printf("Mensaje recibido en el canal Broadcast: %+v", msg)

        // Envía el mensaje a todos los clientes conectados
        for client := range h.clients {
            err := client.WriteJSON(msg)
            if err != nil {
                // Log: Error al enviar el mensaje al cliente
                log.Printf("Error al enviar el mensaje al cliente: %v", err)
                client.Close()
                delete(h.clients, client)
                log.Printf("Cliente desconectado. Total de clientes conectados: %d", len(h.clients))
            } else {
                // Log: Mensaje enviado correctamente al cliente
                log.Printf("Mensaje enviado correctamente al cliente: %+v", msg)
            }
        }
    }
}