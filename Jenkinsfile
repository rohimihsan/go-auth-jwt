pipeline {
    agent any

    environment {
        GOPATH = "${WORKSPACE}/go"
    }

    stages {
        // stage('Checkout') {
        //     steps {
        //         // Checkout the code from your Git repository
        //         checkout scm
        //     }
        // }

        stage('Build') {
            steps {
                // Set the Go workspace
                // dir("go/src/github.com/your-username/your-golang-project") {
                    // Clean the workspace and build the Golang project
                    sh "go clean"
                    sh "go build -o myapp"
                // }
            }
        }

        stage('Test') {
            steps {
                // Set the Go workspace
                // dir("go/src/github.com/your-username/your-golang-project") {
                    // Run Golang tests
                    sh "go version"
                // }
            }
        }

        stage('Deploy') {
            steps {
                // Here you can add steps to deploy your Golang application (e.g., copying to a server)
                // For example, you can use SCP or SSH to copy the binary to a remote server
                // sh "scp myapp user@your-server:/path/to/deploy/"
                sh "go run main.go"
            }
        }
    }

    post {
        success {
            echo 'Build and deployment successful!'
        }
        failure {
            echo 'Build or deployment failed!'
        }
    }
}
