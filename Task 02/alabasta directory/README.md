## The engineering issues identified:
The main issue i found in the alabasta directory was the issue of __Configuration drift__. This happens when files with different settings are in the same project and hence gets confused and doesn't align together to decide how the system should run.
The main file _application.toml_ could handle Port 9010 but the old file _override.toml_ was forced to use Port 9011 instead.

## The approach used to investigate them: 
1. I started by running the _cargo check_ command to check the overall project. the code run successfully without any errors which proved that there were no errors with the syntax or source code.
2. Then i looked into the projects folders by using _ls -la config/_ command.
3. Finally used cat commands to read through the files settings. by looking into _application.toml_ and _override.toml_ side-by-side i plotted the conflicting ports and rectified it by assuming that Port 9010 was correct. I used the _nano config/override.toml_  command to open up the forced settings file directly in my terminal. I erased the incorrect port number 9011 and changed it to 9010 so it matches the rest of the application parameters. I verified it by using cat to read the file text again.
