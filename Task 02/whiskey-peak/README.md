## The engineering issues identified: 
The main issue i found in the whiskey-peak directory was the issue of Configuration drift. This happens when files with different settings are in the same project and hence gets confused and doesn't align together to decide how the system should run.
The main blueprint file application.toml specified that the system must run as _legacy mode = true_. however runtime.toml was forcing the parameter to false.

## The approach used to investigate them:
1. I started by running _cargo check_ to see if the files were broken. Everything ran fine right away, which proved the main code was healthy.
2. after that i ran the command _ls -la_ config to check what was inside the settings folder and noticed a second file sitting there named _runtime.toml_.
3. I used the _cat_ command to read both files on my screen side-by-side. By looking through them, I spotted that they had different legacy_mode with one being true and other having false.

## The fixes applied:
I used the command _nano runtime.toml_ to open the file and change _legacy_node = false_ to _legacy_node = true_ so it matches the main settings. To save my work, I pressed _Ctrl+O_ and hit _Enter_, followed by _Ctrl+X_ to exit the screen. After that, I ran _cat runtime.toml_ to make sure it saved properly, and used _cargo check_ to double-check that everything compiles perfectly.

## he Rust, Git, and Linux concepts involved: 
1. Cargo Check: A fast compiler test that scans your code files to verify if the code layout and grammar are valid without wasting time fully building an application.
2. Linux Terminal Commands: Text-based instructions used to control a computer without a mouse.

    1. cd    : change folders
    2. cat   : read text
    3. ls -la: list files
3. Nano Editor: A basic, lightweight text editor built directly into the terminal window used to edit files using your keyboard.

## Any assumptions made while restoring historical behavior:
I assumed that legacy_node = true was correct, so I changed false to true inside the runtime file to match it.

