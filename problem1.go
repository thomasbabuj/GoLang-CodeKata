/*
*  Problem 1
*  If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these
*  multiples is 23.
*
*  Find the sum of all the multiples of 3 or 5 below 1000.
*
*  Reference links : https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
 */



package main

import (
    "fmt"
    "log"
    "time"
)

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name ,elapsed)
}

func sumMulitples(number int ) (answer string ) {

   sumOfMultiples := 0

   defer timeTrack(time.Now(), "Sum of all the multiples")

   for i:=1; i <= number; i++ {
      if  i % 3 == 0 || i % 5 == 0 && i != number {
        sumOfMultiples += i
      }
   }

   return fmt.Sprintf("Sum of all the multiples of 3 or 5 below %d : %d \n", number, sumOfMultiples)

}

func main() {
   numberRange := 1000000000

   output := sumMulitples (numberRange)

   fmt.Printf("%s \n", output)

}
