/*
Author: Theodore Ikehara

Date: 2/1/2023

About: This is an sftp client that transfers your files on you local machine to an ssh server in real time
*/

package main

import(
	"fmt"
	"strings"
	"os"
	"net/http"
	"time"
	"github.com/pkg/sftp"
)

func main(){
	sshConfig := &ssh.ClientConfig{
		user: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", host+":"+port, sshConfig)
	if err != nil{
		return err
	}
	defer conn.Close()
	
	file, err := sftpClient.Create("Some path to remote file")
	if err != nil{
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello, SFTP!"))
	if err != nil{
		return err
	}
}
