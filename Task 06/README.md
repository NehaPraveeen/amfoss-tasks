## Task 06: Pirate King's Scheduler

I treated this like a basic queue math loop using my Python background. I created a strict `struct` template to hold each crew's information. Then, I wrote a `for` loop with a running clock tracker (`currentTime`) to move time forward as each crew unloads. The code automatically subtracts their arrival hour from their finish hour to find the waiting times and prints a simple text timeline chart on the screen.


<img width="855" height="205" alt="Screenshot From 2026-08-30 17-19-26" src="https://github.com/user-attachments/assets/eaea65b6-e8a4-4378-a4aa-5d0b820d0bb2" />




I wrote and ran the code for the **First Come First Serve (FCFS)** logic, and it works perfectly in my terminal. However, because Go is completely new to me and I only have a basic school programming background, I couldn't fully understand how to write the code loops for **Shortest Job First (SJF)** and **Round Robin (RR)** before the deadline. 


- **CPU Scheduling:** How a computer manages a waiting line of different tasks.
- **FCFS Algorithm:** A simple rule where the first process to arrive gets handled first.
- **Golang Basics:** How to import `fmt` to print text, make loops, and use structured templates instead of loose Python dictionaries.
 
