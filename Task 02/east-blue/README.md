## The engineering issues identified:
1. The main issue I found in the east-blue directory was a warning about __Mixed Protocol Versions__. This happens when a project reads a list of data entries, but the entries are written in different version languages that don't align together.
2. The first station (station-alpha) was using an old language called v1.The second station (station-beta) was using a newer language called v2.

## The approach used to investigate them:
1. First I went into east-blue. After which i ran the command _cat logs/restore.log_ where I noticed a warning sign saying _"legacy station snapshot contained mixed protocol versions"_ which made me realize that the files are mixed up.
2. Since it mentioned leagacy-stations and I remembered a similar file name in the directory where archives was there, i inspected that using command _cat legacy-stations.yml_ and found out that one had proto "v1" and other had proto "v2".

## The fixes applied:
1. using the command _nano legacy-stations.yml_ i changed the proto "v1" to proto "v2", to save it i used ctrl+o and pressed enter. that was followed by ctrl+X to exit that screen.
2. After that process i used command _cat legacy-stations.yml_ though which i confirmed that both prototype was same and was changed to "v2".
3. To double confirm i used command _cd archives/east-blue_ followed by _cargo check_ and got the comment _Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.05s_.

<img width="1512" height="979" alt="Screenshot From 2026-08-29 20-40-05" src="https://github.com/user-attachments/assets/a5c85e52-653d-4e9f-9eb3-d1bb72ee79de" />

## The Rust, Git, and Linux concepts involved:
1. Data Protocols: The version language rules that computers use to talk to each other so they can understand incoming data.
2. YAML Format: A simple text layout format used by developers to make clean, organized lists using text indentation.
3. Log Journals: Special error diary files created by programs to help engineers trace down silent system warnings.
4. Cargo Check: A fast compiler test that scans code files to verify grammar and layouts without building a heavy application.
5. Linux Terminal Commands: Text-based instructions used to move around and control a computer without using a mouse.

    1. cd    : change folders
    2. cat   : read text contents
    3. clear : clean up the terminal screen
    4. ls -la: list all files and folders in a directory, including hidden ones

7. Nano Editor: A basic, lightweight text editor built directly into the terminal window used to edit files using keyboard.

## Any assumptions made while restoring historical behavior:
I assumed that protocol v2 was correct and changed protocol v1 to v2.







