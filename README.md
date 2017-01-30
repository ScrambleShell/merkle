## [Merkle](https://en.wikipedia.org/wiki/Merkle_tree)

[![GoDoc](https://godoc.org/github.com/chasestarr/merkle?status.svg)](https://godoc.org/github.com/chasestarr/merkle)

```go
import (
  "fmt"

  "github.com/chasestarr/merkle"
)

func main() {
  input := []string{"abcd", "efgh", "ijkl"}
	hash := merkle.Hash(input)
	fmt.Println(hash) // "A8uyUMnjjSXD0Tb0Tm-oyrJ3tVnif0ry4l0YbTSxvos="
}
