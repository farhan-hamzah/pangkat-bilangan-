package main
import "fmt"

func pangkat(base, eks int) int {
    if eks == 0 {
        return 1
    }
    if eks > 0 {
        return base * pangkat(base, eks-1)
    }
    return 1
}

func main() {
    var d1, d2 int
    fmt.Scan(&d1, &d2)
    fmt.Println(pangkat(d1, d2))
}
