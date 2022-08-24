# Gojira
A CLI application that integrates with Jira.
After annoyance by having to leave Vim / IDE / Editor to check something in jira,
I created a little application that brings jira issues straight into the terminal.

## Installation

There are multiple ways of installing this software.

#### Homebrew (Mac OS & Linux)
 ```bash
 > brew tap sbvalois/sbvalois
 > brew install gojira
 ```
#### Download the binary
- Find the newest release under [releases](https://github.com/sbvalois/gojira/releases)
- Download the correct version for your system.
- Add the binary to your path for convenience. 

#### Build directly from source
This requires you to have go installed on your computer.
Simply run the command to install it in go's `bin` directory
```bash
> go install
```
 

## Usage

### Setup
Run the Gojira setup wizard
```bash
> gojira setup
```

### Issues assigned to you
Gets all issues with status `todo` / `open`
```bash
> gojira issues open
```
Gets all issues with status `in progress`
```bash
> gojira issues inprogress
```

### Info on a issue
Prints information of a single issue
```bash
> gojira issue {{ISSUE-123}}
```

### Change status on issue
Sets an issue to `in progress`
```bash
> gojira issue start {{ISSUE-123}}
```
Sets an issue to `review`
```bash 
> gojira issue review {{ISSUE-123}}
```
Sets an issue to `done`
```bash 
> gojira issue done {{ISSUE-123}}
```
Sets an issue back to `open`
```bash
> gojira issue open {{ISSUE-123}}
```

### In browser
Opens a specific jira issue in your default browser
```bash
> gojira issue browser {{ISSUE-123}}
```
