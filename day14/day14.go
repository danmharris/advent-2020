package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
    "regexp"
)

func ApplyMaskToVal(val int, mask string) int {
    padded := []byte(fmt.Sprintf("%036b", val))
    for i, v := range mask {
        if v == '1' {
            padded[i] = '1'
        }
        if v == '0' {
            padded[i] = '0'
        }
    }

    newVal, _ := strconv.ParseInt(string(padded), 2, 0)
    return int(newVal)
}

func applyMaskToAddr(addr []byte, mask string, idx int) []int{
    if idx >= len(addr) {
        newAddrSingle, _ := strconv.ParseInt(string(addr), 2, 0)
        return []int{int(newAddrSingle)}
    }
    if mask[idx] == '1' {
        addr[idx] = '1'
        return applyMaskToAddr(addr, mask, idx+1)
    }
    if mask[idx] == '0' {
        return applyMaskToAddr(addr, mask, idx+1)
    }
    if mask[idx] == 'X' {
        addr[idx] = '0'
        newAddr := applyMaskToAddr(addr, mask, idx+1)
        addr[idx] = '1'
        newAddr = append(newAddr, applyMaskToAddr(addr, mask, idx+1)...)
        return newAddr
    }

    return []int{}
}

func ApplyMaskToAddr(addr int, mask string) []int {
    addrBytes := []byte(fmt.Sprintf("%036b", addr))

    return applyMaskToAddr(addrBytes, mask, 0)
}

func main() {
    sc := bufio.NewScanner(os.Stdin)

    var mask string
    mem := make(map[int]int)
    mem2 := make(map[int]int)

    maskRe := regexp.MustCompile(`^mask = (\w+)$`)
    memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

    for sc.Scan() {
        line := sc.Text()

        if m := maskRe.FindStringSubmatch(line); m != nil {
            mask = m[1]
        } else if m := memRe.FindStringSubmatch(line); m != nil {
            addr, _ := strconv.Atoi(m[1])
            val, _ := strconv.Atoi(m[2])

            // Part 1
            mem[addr] = ApplyMaskToVal(val, mask)

            // Part 2
            addrs := ApplyMaskToAddr(addr, mask)
            for _, a := range addrs {
                mem2[a] = val
            }
        } else {
            panic("Unknown line")
        }
    }

    var total int
    var total2 int
    for _, v := range mem {
        total += v
    }
    for _, v := range mem2 {
        total2 += v
    }

    fmt.Printf("%d\n", total)
    fmt.Printf("%d\n", total2)
}
