## The engineering issues identified: 
The main issue i found in the whiskey-peak directory was the issue of Configuration drift. This happens when files with different settings are in the same project and hence gets confused and doesn't align together to decide how the system should run.
The main blueprint file application.toml specified that the system must run as _legacy mode = true_. however runtime.toml was forcing the parameter to false.

## The approach used to investigate them:
1. I started by running _cargo check_ to see if the files were broken. Everything ran fine right away, which proved the main code was healthy.
2. after that i ran the command _ls -la_ config to check what was inside the settings folder and noticed a second file sitting there named _runtime.toml_.
3. I used the _cat_ command to read both files on my screen side-by-side. By looking through them, I spotted that they had different legacy_mode with one being true and other having false.

