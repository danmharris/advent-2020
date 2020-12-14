package main

import (
    "bufio"
    "os"
    "fmt"
    "strings"
    "strconv"
)

func FindEarliest(time int, busIds []int) (int, int) {
    for true {
        time++
        for _, id := range busIds {
            if id < 0 {
                continue
            }
            if time % id == 0 {
                return time, id
            }
        }
    }

    return -1, -1
}

func main() {
    sc := bufio.NewScanner(os.Stdin)

    sc.Scan()
    time, _ := strconv.Atoi(sc.Text())
    sc.Scan()
    busIdsFull := strings.Split(sc.Text(), ",")

    var busIds []int
    for _, i := range busIdsFull {
        if i != "x" {
            id, _ := strconv.Atoi(i)
            busIds = append(busIds, id)
        } else {
            busIds = append(busIds, -1)
        }
    }

    earliest, id := FindEarliest(time, busIds)
    fmt.Printf("Earliest: %d\tID: %d\tProduct: %d\n", earliest, id, (earliest-time)*id)
}
