package main

import (
    "bufio"
    "fmt"
    "log"
    "math/big"
    "net"
    "os"
    "time"
)

func main() {
    var ip, port string
    fmt.Print("Digite o IP do servidor: ")
    fmt.Scanln(&ip)
    fmt.Print("Digite a porta do servidor: ")
    fmt.Scanln(&port)

    conn, err := net.Dial("tcp", ip+":"+port)
    if err != nil {
        log.Fatalf("Erro ao conectar: %v\n", err)
    }
    defer conn.Close()
    log.Println("Conectado ao servidor.")

    stdin := bufio.NewReader(os.Stdin)
    server := bufio.NewReader(conn)

    for {
        fmt.Print("Você: ")
        text, err := stdin.ReadString('\n')
        if err != nil {
            log.Fatalf("Erro ao ler stdin: %v\n", err)
        }

        startNano := time.Now().UnixNano()

        if _, err := conn.Write([]byte(text)); err != nil {
            log.Fatalf("Erro ao enviar: %v\n", err)
        }

        resp, err := server.ReadString('\n')
        if err != nil {
            log.Fatalf("Erro ao ler resposta: %v\n", err)
        }

        endNano := time.Now().UnixNano()
        diffNano := endNano - startNano

        frac := new(big.Rat).SetFrac(
            big.NewInt(diffNano),
            big.NewInt(1000000000),
        )

        diffStr := frac.FloatString(20)

        fmt.Printf("Tempo de resposta: %s segundos\n", diffStr)
        fmt.Printf("Servidor: %s", resp)
    }
}
