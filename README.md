# About Command Vault

Command Hive is a small CLI application whose main purpose is to allow the user to save commands in a array. The user can also delete a command, list all commands, search a command.

Below you can check all the commands available so far.

## How to use

| Command | Description            |
| ------- | ---------------------- |
| ./cv -a | Adds a new command     |
| ./cv -l | List all commands      |
| ./cv -d | Delete a command by ID |
| ./cv -s | Search by id           |
| ./cv -h | Help                   |
| ./cv -v | Shows the version      |

## Screenshots

<!-- ![image](https://user-images.githubusercontent.com/27534241/183315162-e8027a6c-e7f8-43b0-bffb-5c51d53b0d8e.png) -->

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
