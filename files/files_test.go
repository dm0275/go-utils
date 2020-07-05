package files

import (
    "github.com/dm0275/go-utils/errors"
    "io/ioutil"
    "os"
    "testing"
)

func TestReadFile(t *testing.T) {
    type args struct {
        filePath string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {name: "Validate_FileContents", args: args{filePath: "./sample.txt"}, want: "valid"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := ReadFile(tt.args.filePath); got != tt.want {
                t.Errorf("ReadFile() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestOverwriteFile(t *testing.T) {
    type args struct {
        filePath string
        content  string
        perms    os.FileMode
    }
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {name: "Validate_No_Errors", args: args{filePath: "./sample.txt", content: "Updated", perms: 0644}, wantErr: false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := OverwriteFile(tt.args.filePath, tt.args.content, tt.args.perms); (err != nil) != tt.wantErr {
                t.Errorf("OverwriteFile() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestOverwriteFile_FileContents(t *testing.T) {
    type args struct {
        filePath string
        content  string
        perms    os.FileMode
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {name: "Validate_FileContents", args: args{filePath: "./sample.txt", content: "Updated", perms: 0644}, want: "Updated"},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := OverwriteFile(tt.args.filePath, tt.args.content, tt.args.perms)
            text, fileErr := ioutil.ReadFile("./sample.txt")
            errors.CheckError(fileErr)
            if string(text) != tt.want {
                t.Errorf("OverwriteFile() error = %v, want %v", err, tt.want)
            }
        })
    }
}
