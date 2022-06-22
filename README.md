# Gojira
A CLI application that integrates with Jira.
After annoyance by having to leave Vim / IDE / Editor to check something in jira,
I created a little application that brings jira issues straight into the terminal.

## Usage

### Setup
Run the Gojira setup wizard
```
gojira setup
```

### Issues assigned to you
Gets all issues with status `todo` / `open`
```
gojira issues open
```
Gets all issues with status `in progress`
```
gojira issues inprogress
```

### Info on a issue
Prints information of a single issue
```
gojira issue {{ISSUE-123}}
```

### Change status on issue
Sets an issue to `in progress`
```
gojira issue start {{ISSUE-123}}
```
Sets an issue to `review`
```
gojira issue review {{ISSUE-123}}
```
Sets an issue to `done`
```
gojira issue done {{ISSUE-123}}
```
Sets an issue back to `open`
```
gojira issue open {{ISSUE-123}}
```

### In browser
Opens a specific jira issue in your default browser
```
gojira issue browser {{ISSUE-123}}
```
