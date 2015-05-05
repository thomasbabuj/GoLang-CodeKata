/*
*  Problem 1
*  If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these
*  multiples is 23.
*
*  Find the sum of all the multiples of 3 or 5 below 1000.
*
*  Reference links : https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
*
*  To get a more efficient solution you could also calculate the sum of the numbers less
*  than 1000 that are divisible by 3, plus the sum of the numbers less than 1000 that are divisible
*  by 5. But as you have summed numbers divisible by 15 twice you would have to subtract the
*  sum of the numbers divisible by 15.
 */

package main

import (
    "fmt"
    "log"
    "time"
)

const target uint64 = 1000

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name ,elapsed)
}

func sumDivibleBy(number uint64) (sum uint64) {
    defer timeTrack(time.Now(), "Sum of all the multiples")
    p := target / number
    sum = number * (p * (p + 1)) / 2
    return
}

func main() {
    output := sumDivibleBy(3) + sumDivibleBy(5) - sumDivibleBy(15)

    fmt.Printf("Sum of all the multiples of 3 or 5 below %d : %d \n ", target, output )
}
