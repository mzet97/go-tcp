package main

import (
    "bufio"
    "io"
    "log"
    "net"
    "strings"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "9000"
    CONN_TYPE = "tcp"
)


func handleConnection(conn net.Conn) {
    defer conn.Close()
    log.Printf("Conexão recebida de %s\n", conn.RemoteAddr())

    reader := bufio.NewReader(conn)
    for {
        msg, err := reader.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                log.Printf("Erro ao ler: %v\n", err)
            }
            return
        }

        trimmed := strings.TrimSpace(msg)
        log.Printf("Cliente disse: %q\n", trimmed)
        resposta := trimmed + "\n"
        
        switch strings.ToLower(trimmed) {
        case "ping":
            resposta = "pong\n"
        case "hello":
            resposta = "Olá!\n"
        default:
            resposta = "Você disse: " + trimmed + "\n"
        }

        if _, err := conn.Write([]byte(resposta)); err != nil {
            log.Printf("Erro ao escrever: %v\n", err)
            return
        }
    }
}

func main() {
    listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        log.Fatalf("Não foi possível iniciar o listener: %v\n", err)
    }
    defer listener.Close()
    log.Println("Servidor TCP rodando na porta 9000...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Erro ao aceitar conexão: %v\n", err)
            continue
        }
        go handleConnection(conn)
    }
}
