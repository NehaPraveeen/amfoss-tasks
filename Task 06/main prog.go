package main
import "fmt"
typr PirateCrew struct{
  ID string
  ArrivalTime int
  BurstTime int
  CompletionTime int
  TurnaroundTime int
  WaitingTime int
}
funcn main() {
  fmt.Println("--------Pirate King's Sheduler Simulator--------")
  fmt.Print("Timeline Chart: ")
  crews:=[]PirateCrew{
    {ID: "Luffy-Crew", ArrivalTime: 0, BurstTime: 5},
    {ID: "Zoro-Crew", ArrivalTime: 1, BurstTime: 3},
    {ID: "Sanji-Crew", ArrivalTime: 2, BurstTime: 4},
  }
  fmt.Println("ALGORITHM: FIRST COME FIRST SERVE(FCFS)")
  fmt.Println("Timeline Chart: ")
  currentTime:= 0
  TotalWait:= 0
  totalTurnaround:= 0
  for i:=0; i<len(crews); i++ {
    if currentTime<crews[i].ArrivalTime {
      currentTime=cews[i].ArrivalTime
    }
    startTime:= currentTime
    currentTime+=crews[i].BurstTime

    crew[i].CompletionTime=currentTime
    crews[i].TurnaroundTime=crews[i].CompletionTime-crews[i].ArrivalTime
    crews[i].WaitingTime=crews[i].TurnaroundTime-crews[i]BurstTime
    TotalWait+=crews[i].WaitingTime
    totalTurnaroundTime+=crews[i].TurnaroundTime

    fmt.Printf("[%s: %d-%d]",crews[i].ID,startTime, currentTime)
  }
  fmt.Println("\nCrew ID\t\tArrival\t\tBurst\t\tTurnaround")
  for i:= 0;i<len(crews);i++{
    fmt.Printf("%s\t%d\t%d\t%d\t%d\n"),crews[i].ID, crews[i].ArrivalTIme, crews[i].BurstTime, crews[i].WaitingTime, crews[i].TurnaroundTime)
  }
  numCrews:=float64(len(crews))
  fmt.Printf("\nAverage Waiting Time: %.2f hours\n",float64(totalWait)/numCrews)
  fmt.Printf("Average TurnaroundTime Time: %.2f hours\n",float64(totalTurnaround)/numCrews)
}
















    
    
