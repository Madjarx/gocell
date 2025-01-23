package cli

/*
    ## Command executor

    Executes the CLI command 
    Returns the output
    
*/

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
)

type CommandExecutor struct{}

func NewCommandExecutor() *CommandExecutor {
    return &CommandExecutor{}
}

func (ce *CommandExecutor) Execute(commandPath string, args []string) (string, error) {
    cmd := exec.Command(commandPath, args...)
    
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    
    if err := cmd.Run(); err != nil {
        return "", fmt.Errorf(
            "command failed: %s\nError: %v\nStderr: %s", 
            strings.Join(cmd.Args, " "), 
            err, 
            stderr.String(),
        )
    }
    
    return strings.TrimSpace(stdout.String()), nil
}