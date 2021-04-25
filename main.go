package main

import (
    "strconv"
    "github.com/sethvargo/go-password/password"
    "github.com/gin-gonic/gin"
)

func StringToInt(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        return 0
    }
    return i
}

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.GET("/password-generator", func(c *gin.Context){
        length := StringToInt(c.Query("length"))
        digitLength := StringToInt(c.Query("digit_length"))
        digitSymbol := StringToInt(c.Query("digit_symbol"))
        allowUpper := StringToInt(c.Query("allow_upper"))
        allowUpperBool := false
        disallowRepeat := StringToInt(c.Query("disallow_repeat"))
        disallowRepeatBool := false

        if allowUpper == 1 {
            allowUpperBool = true
        }

        if disallowRepeat == 1 {
            disallowRepeatBool = true
        }



        // Generate a password that is 64 characters long with 10 digits, 10 symbols,
        // allowing upper and lower case letters, disallowing repeat characters.
        res, err := password.Generate(length, digitLength, digitSymbol, allowUpperBool, disallowRepeatBool)
        if err != nil {
            c.JSON(200, gin.H{
                "success": false,
                "message": err.Error(),
            })
            return
        }
        c.JSON(200, gin.H{
            "success": true,
            "message": res,
        })
    })
    r.Run()
}
