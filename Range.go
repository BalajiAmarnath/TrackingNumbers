package main

import "fmt"


type Range struct{
	 lo int 
	 hi int
}
func  Contains(r Range,x int) bool{
   return r.lo<=x && x<=r.hi
}

func ContainsOpen(r Range,x int) bool{
   return r.lo<x && x<r.hi
}

func  setLo(r *Range,x int){
   r.lo = x
}

func  setHi(r *Range,x int){
   r.hi = x
}

func inBetween(r Range,anotherRange Range,x int) bool{
   newRange :=Range{r.hi,anotherRange.lo}
   return ContainsOpen(newRange,x)
}

func copy(r Range)Range{
 copy := Range{r.lo,r.hi}
 return copy 
}
 


func main(){
 
   r1 := Range{10,100}
   fmt.Println("Contains :" , Contains(r1,10))
   fmt.Println("ContainsOpen :" , Contains(r1,10))
   setLo(&r1,50)
   fmt.Println(" setLoCalled Lo = :" ,r1.lo)   
   fmt.Println("Copy:" , copy(r1).lo ," ",copy(r1).hi)
}