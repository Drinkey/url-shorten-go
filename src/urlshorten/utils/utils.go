package utils

import (
	"os"
	"fmt"
    "crypto/sha256"
)

func FileExist(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}


func Sha256Sum(input string) string {
    shaSum := sha256.Sum256([]byte(input))
    return fmt.Sprintf("%x", shaSum)
}
