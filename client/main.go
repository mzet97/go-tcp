// package main

// import (
//     "bufio"
//     "fmt"
//     "log"
//     "math/big"
//     "net"
//     "os"
//     "time"
// )

// func main() {
//     var ip, port string
//     fmt.Print("Digite o IP do servidor: ")
//     fmt.Scanln(&ip)
//     fmt.Print("Digite a porta do servidor: ")
//     fmt.Scanln(&port)

//     conn, err := net.Dial("tcp", ip+":"+port)
//     if err != nil {
//         log.Fatalf("Erro ao conectar: %v\n", err)
//     }
//     defer conn.Close()
//     log.Println("Conectado ao servidor.")

//     stdin := bufio.NewReader(os.Stdin)
//     server := bufio.NewReader(conn)

//     for {
//         fmt.Print("Você: ")
//         text, err := stdin.ReadString('\n')
//         if err != nil {
//             log.Fatalf("Erro ao ler stdin: %v\n", err)
//         }

//         startNano := time.Now().UnixNano()

//         if _, err := conn.Write([]byte(text)); err != nil {
//             log.Fatalf("Erro ao enviar: %v\n", err)
//         }

//         resp, err := server.ReadString('\n')
//         if err != nil {
//             log.Fatalf("Erro ao ler resposta: %v\n", err)
//         }

//         endNano := time.Now().UnixNano()
//         diffNano := endNano - startNano

//         frac := new(big.Rat).SetFrac(
//             big.NewInt(diffNano),
//             big.NewInt(1000000000),
//         )

//         diffStr := frac.FloatString(20)

//         fmt.Printf("Tempo de resposta: %s segundos\n", diffStr)
//         fmt.Printf("Servidor: %s", resp)
//     }
// }

package main

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "math/big"
    "net"
    "time"
)

func main() {
    addr := "127.0.0.1:9000"
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        log.Fatalf("Erro ao conectar: %v\n", err)
    }
    defer conn.Close()
    fmt.Printf("Conectado ao servidor em %s\n\n", addr)

    const maxSize = 1 << 30
    for size := 2; size <= maxSize; size *= 2 {
        payload := bytes.Repeat([]byte{'A'}, size)

        start := time.Now()

        if _, err := conn.Write(payload); err != nil {
            log.Fatalf("Erro ao enviar payload: %v\n", err)
        }

        resp := make([]byte, size)
        if _, err := io.ReadFull(conn, resp); err != nil {
            log.Fatalf("Erro ao ler resposta: %v\n", err)
        }

        elapsedNs := time.Since(start).Nanoseconds()

        rat := new(big.Rat).SetFrac(
            big.NewInt(elapsedNs),
            big.NewInt(1_000_000_000),
        )
        diffStr := rat.FloatString(6)
        
        fmt.Printf("Tamanho: %10d bytes → Tempo de resposta: %s s\n", size, diffStr)
    }
}
