## --> The engineering issues identified:
The main issue i found in the alabasta directory was the issue of __Configuration drift__. This happens when files with different settings are in the same project and hence gets confused and doesn't align together to decide how the system should run.
The main file _application.toml_ could handle Port 9010 but the old file _override.toml_ was forced to use Port 9011 instead.

## --> The approach used to investigate them: 
1. I started by running the _cargo check_ command to check the overall project. the code run successfully without any errors which proved that there were no errors with the syntax or source code.
2. Then i looked into the projects folders by using _ls -la config/_ command.
3. Finally used cat commands to read through the files settings. by looking into _application.toml_ and _override.toml_ side-by-side i plotted the conflicting ports and rectified it by assuming that Port 9010 was correct. 

## --> Fixes Applied :
I used the _nano config/override.toml_ command to open up the forced settings file directly in my terminal. I erased the incorrect port number 9011 and changed it to 9010 so that it matches. I verified it by using cat to read the file text again.

## --> Rust, Git and Linux concepts involved: 
1.**Ports:** Virtual door numbers on a computer that let different apps send and receive internet messages without getting their data mixed up.

2. **Configuration Drift:** A problem that happens when different settings files within the same application mismatch, conflict, and cause the system to break.

3.**Cargo and Rust:** Tools used to manage and build projects written in the Rust programming language. 

4.**Cargo Check:** A fast compiler test that scans your code files to verify if the code layout and grammar are valid without wasting time fully building an application.

5.**Linux Terminal Commands:** Text-based instructions used to control a computer without a mouse .

        1. _cd_  :change folders
        2. _ls_  : list files 
        3. _cat_ : read text
        
6.**Nano Editor:** A basic, lightweight text editor built directly into the terminal window used to edit files using your keyboard.

## --> assumptions made while restoring historical behavior: 
I assumed that port 9010 is correct and changed port 9011 to 9010.

