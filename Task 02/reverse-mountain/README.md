## --> The engineering issues identified:
The main issue I found in the reverse-mountain directory was that __an important folder path was completely missing__. This happens when the project is looking for a specific folder, but that folder isn't there, so the system gets confused and prints out warning flags.The main settings file _application.toml_ expected a folder called "assets", but looking at the directory, that folder didn't exist anywhere.

## --> The approach used to investigate them:
I started by running the _cargo check_ command to see if the files were broken. The code ran successfully without any errors which proved that there were no bugs with the main code. .Then I checked the application's diary notes by running the cat logs/runtime.log command and saw a warning line saying that an asset directory path didn't exist.Finally I checked inside the settings folder by running the _ls -la config/_ command. By looking through the list, I confirmed that the required folder path wasn't actually there.

## --> Fixes Applied :
Since the code compiled fine but the physical folder was missing, I used the command _mkdir config/assets_ to create the empty folder path right inside the terminal. This matches what the settings file wanted so it runs cleanly without throwing warning flags anymore.

## --> Rust, Git and Linux concepts involved:
1. Directory Layouts: Understanding that programs sometimes need physical folders to exist in the system to run cleanly without warnings.
2. Log Journals: Special error diary files created by programs to help engineers trace down silent system warnings.
3. Cargo and Rust: Tools used to manage and build projects written in the Rust programming language.
4. Cargo Check:A fast compiler test that scans your code files to verify if the code layout and grammar are valid without wasting time fully building an application.
5. Linux Terminal Commands: Text-based instructions used to control a computer without a mouse.
         1.  cd    : change folders
         2.  ls    : list files. 
         3.  cat   : read text
         4.  mkdir : make a brand-new empty folder
         5.  Nano Editor: A basic, lightweight text editor built directly into the terminal window used to edit files using your keyboard.
        
## --> assumptions made while restoring historical behavior:
I assumed that since the main file wanted an assets folder to be there, creating it was the right way to match the settings.
