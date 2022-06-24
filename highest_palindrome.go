package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'highestValuePalindrome' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER n
 *  3. INTEGER k
 */

func highestValuePalindrome(s string, n int32, k int32) string {
    var count int32 = 0
    var boolean []bool = make([]bool, n)
    var nn int  = int(n)
    for i  := 0; int32(i) < n/2; i++ {
            if s[i]==s[nn-i-1] {
                boolean[i] = false
            }else{
                boolean[i] = true
                count++
                if(s[nn-i-1]>s[i]){
                    s =s[:i]+string(s[nn-i-1])+s[i+1:]
                }else{
                    s =s[:nn-i-1]+string(s[i])+s[nn-i:]
                }
            }
    }
    if count>k {
        return "-1"
    }else { 
     count = k-count
     fmt.Print(count)
     var j =0 
     for ((count>0))&& (j!=nn-1){
        if(string(s[j])!="9"){
        if(boolean[j]==true){
            count--
            if(count<0){
                count+=1
                break
            }
        }else{
        count--
        count--
        if(count<0){
                count+=2
                break
            }
        }
        s = s[:j]+"9"+s[j+1:]
        s = s[:nn-j-1]+"9"+s[nn-j:]
        j++
        }else{
            j++
        }
     }
     if (n%2!=0)&&(count%2!=0)&&(count>0) {
         s = s[:nn/2]+"9"+s[(nn/2)+1:]
         count--
     }
     return s
    }

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    s := readLine(reader)

    result := highestValuePalindrome(s, n, k)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
