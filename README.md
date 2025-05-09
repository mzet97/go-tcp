# go-tcp

**go-tcp** é um cliente de benchmark em Go para medir a latência de ida e volta (RTT) em conexões TCP. Ele envia payloads de tamanhos em potências de 2 (de 2 bytes até 1 GiB) a um servidor de echo TCP e calcula o tempo de resposta com alta precisão.

## Recursos

* Conexão TCP simples e configurável (IP e porta via prompt)
* Envio de payloads cujo tamanho dobra a cada iteração (2, 4, 8 … até 1 GiB)
* Medição de RTT em nanosegundos convertida para segundos
* Formatação de resultados com precisão de 6 casas decimais
* Exibição clara de tempo de resposta para cada tamanho de payload

## Pré-requisitos

* Go 1.21 ou superior instalado ([Download](https://go.dev/dl/))

## Instalação

1. Clone este repositório:

   ```bash
   git clone https://github.com/mzet97/go-tcp.git
   cd go-tcp
   ```
2. Baixe as dependências:

   ```bash
   go mod tidy
   ```
3. Compile o cliente:

   ```bash
   go build -o bin/go-tcp
   ```

> **Observação:** Ajuste o nome do output ou diretório `bin/` conforme sua preferência.

## Uso

1. Execute o cliente:

   ```bash
   ./bin/go-tcp
   ```
2. Informe o **IP** e a **porta** do servidor TCP de echo quando solicitado:

   ```text
   Digite o IP do servidor: 127.0.0.1
   Digite a porta do servidor: 9000
   ```
3. O cliente fará automaticamente envios de payloads de 2, 4, 8 … bytes até 1 GiB e exibirá:

   ```text
   Tamanho:         2 bytes → Tempo de resposta: 0.000523 s
   Tamanho:         4 bytes → Tempo de resposta: 0.000612 s
   …
   Tamanho: 1073741824 bytes → Tempo de resposta: 8.593567 s
   ```

## Exemplo de servidor de echo

Para melhores resultados de benchmark, utilize um servidor que simplesmente retorne os mesmos bytes recebidos. Um exemplo em Go:

```go
listener, _ := net.Listen("tcp", ":9000")
for {
    conn, _ := listener.Accept()
    go func(c net.Conn) {
        defer c.Close()
        buf := make([]byte, 64*1024)
        for {
            n, err := c.Read(buf)
            if err != nil { break }
            c.Write(buf[:n])
        }
    }(conn)
}
```

## Contribuição

Contribuições são bem-vindas! Siga estes passos:

1. Fork este repositório
2. Crie uma branch para sua feature (`git checkout -b feature/nome-da-feature`)
3. Faça commit das suas alterações (`git commit -m "Add nova feature"`)
4. Faça push para a branch (`git push origin feature/nome-da-feature`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a [MIT License](./LICENSE). Veja o arquivo LICENSE para mais detalhes.
