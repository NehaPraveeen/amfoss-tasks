## Task 06: Pirate King's Scheduler

## --> The engineering issues identified:
The main task was to build a backend simulation in Golang that calculates how long different pirate crews have to wait in line to unload cargo at an island dock. Only one ship can use the dock at a time, so if other ships show up later, they get stuck waiting out in the water. 

I needed to calculate the total waiting time, turnaround time, averages, and print a text timeline chart directly onto the terminal screen. The task asked for three algorithms: FCFS, SJF, and Round Robin.

## --> The approach used to investigate and tackle them:
1. **Using Python habits:** Since I only know basic Python from school, stepping into Golang was like learning a whole new language. I focused on what I knew best—using loops and tracking numbers moving along a timeline.
2. **Creating the data template:** I used a `PirateCrew struct` to make a clear template that keeps track of each ship's name, arrival hour, and unloading values.
3. **Writing the FCFS math loop:** I set up a loop that goes through the crews list one by one using a running `currentTime` clock tracker. If a ship hasn't sailed up yet, the clock winds forward to its arrival time. As they unload, the clock increases by their burst time, and the code subtracts the hours to find exactly how long they sat idle in the water.

## --> Fixes Applied & Honest Review:
I successfully wrote and executed the code for the **First Come First Serve (FCFS)** logic, and it runs perfectly in my terminal. However, because everything in Go is so new to me and I only have a basic school programming background, I couldn't fully understand how to implement the complex logic variations for **Shortest Job First (SJF)** and **Round Robin (RR)** in time for the deadline. Instead of copy-pasting code I didn't understand, I decided to submit what I fully built and mastered using what I know.

## --> Rust, Git and Linux concepts learned:
1. **CPU Scheduling:** Learning how computer systems handle a queue of tasks when multiple programs are waiting.
2. **First Come First Serve (FCFS):** A straightforward sorting rule where whoever arrives first gets processed completely first.
3. **Golang Basics:** Learning how to import modules like `fmt` for text printing, setting up strict data structures, and formatting data tables using tabs (`\t`).
4. **Go Run Command:** Using the terminal tool to compile and execute a backend simulation script instantly on the fly.

## --> Resources used:
- amFOSS GitBook tasks guide overview.
- Golang documentation for basic loops and string formatting structures.
 
