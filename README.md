# Command Hive

Command Hive is a small CLI application whose main purpose is to allow the user to save commands in a array. The user can also delete a command, list all commands, search a command.

Below you can check all the commands available so far.

## List of commands

| Command | Description            |
| ------- | ---------------------- |
| ./ch -a | Adds a new command     |
| ./ch -l | List all commands      |
| ./ch -d | Delete a command by ID |
| ./ch -s | Search by id           |
| ./ch -h | Help                   |
| ./ch -v | Shows the version      |

## Screenshots

### Add commands : 
![image](https://github.com/rishabdhar12/commands-hive/blob/master/screenshots/add_commands.png)

### List commands : 
![image](https://github.com/rishabdhar12/commands-hive/blob/master/screenshots/list_commands.png)

## Improvements

- Copy to clipboard feature.
- Add ability to sync commands with gdrive.
- Add search by category, platform.
- Master password to encrypt the file.

## Notes

- As of now it is only available for Unix systems.

## Current Version

- 1.0.0

## How to build the application

- go build -o ch main.go
