Replace the required fields (which are highlighted) in `main.go`, and you're done! The information will be sent and formatted automatically. I take no responsibility for misuse. 

To run a Go program compiled automatically every time the operating system is started, and in a hidden way (without displaying a window), you can create a startup service. This service can be configured to run in the background, without interaction with the user. Here are the general steps for setting this up:

- Create a startup script:
    Create a shell script that starts your compiled Go program. This script will run when the system starts. You can use the approaches mentioned earlier to hide the terminal window. For example, using nohup:
```
#!/bin/bash
nohup /path/to/your/program/go &
```
Save this script as start_program.sh.

Set the script as a service:
On systemd-based systems (such as many modern Linux distributions), you can create a service file to start your script. Create a file called my_program.service in /etc/systemd/system/ with the following contents:
```
[Unit]
Description=My Go program

[Service]
ExecStart=/path/to/start_program.sh
Type=simple
Restart=always

[Install]
WantedBy=multi-user.target
```
Replace /path/to/start_program.sh with the actual path of the script you created.

Activate the service:
After creating the service file, reload the systemd daemon to read the changes, start the service and activate it so that it starts automatically during boot:
```
sudo systemctl daemon-reload
sudo systemctl start my_program
sudo systemctl enable my_program

This will cause your Go program to start automatically every time the operating system starts, and it will run in a hidden manner.
```

To run a program automatically every time Windows starts and hide the program window, you can add a shortcut to the program in the user's startup folder and set it to run minimized. Here are the steps to do this:

1. Create a shortcut to the program:
    First, create a shortcut to the compiled Go program. You can do this manually or using File Explorer:

- Manually:
            Go to the location of the compiled program.
            Right-click on the program's executable file and select "Create shortcut".
            Copy the created shortcut to the desktop temporarily.

- Using File Explorer:
            Navigate to the location of the compiled program.
            Hold down the Shift key and right-click on the program's executable file.
            In the context menu that appears, select "Copy as path".

2. Open the user's startup folder:
3. Press Win + R to open the Run dialog box and type shell:startup. This will open the user's startup folder.

4. Paste the shortcut into the startup folder:
- Paste the shortcut you created earlier into the user's startup folder. This will cause the program to start automatically when Windows starts.
5. Configure the shortcut to run minimized:
- After pasting the shortcut into the startup folder, right-click on the shortcut and select "Properties". On the "Shortcut" tab, in the "Run" field, select "Minimized" from the drop-down menu.

6. Make sure that the shortcut points to the correctly compiled Go program.
