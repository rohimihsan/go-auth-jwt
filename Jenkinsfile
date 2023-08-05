pipeline {
    agent any
    
    // environment {
    //     // Set Go workspace path (optional)
    //     GOPATH = "PATH='/usr/local/go/bin:$PATH"
    // }
    
    stages {        
        stage('Test') {
            steps {
                sh "go version"
            }
        }

        stage('Build') {
            steps {
                // Set the Go workspace
                // dir("go/src/github.com/rohimihsan/go-auth-jwt") {
                    // Build the Golang project
                    sh "go build -o go-auth-jwt"
                // }
            }
        }
        
        stage('Deploy') {
            steps {
                // Here you can add steps to deploy your Golang application (e.g., copying to a server)
                // For example, you can use SCP or SSH to copy the binary to a remote server
                sh "go run main.go"
            }
        }
    }
}