pipeline {
    agent any

    environment {
        PATH = "${PATH}:/usr/local/go/bin"
    }

    stages {
        // stage('Checkout') {
        //     steps {
        //         // Checkout the code from your Git repository
        //         checkout scm
        //     }
        // }
        stage('Test') {
            steps {
                // Set the Go workspace
                // dir("go/src/github.com/your-username/your-golang-project") {
                    // Run Golang tests
                    sh "go version"
                // }
            }
        }

        // stage('Init') {
        //     steps {
        //         // Set the Go workspace
        //         // dir("go/src/github.com/your-username/your-golang-project") {
        //             // Clean the workspace and build the Golang project
                
        //         catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
        //         // Your build steps here
        //             sh "go mod init go-auth-jwt"
        //             sh "go mod tidy"   
        //         }                
        //         // }
        //     }
        // }

        stage('Build') {
            steps {
                // Set the Go workspace
                // dir("go/src/github.com/your-username/your-golang-project") {
                    // Clean the workspace and build the Golang project
                     catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                        sh 'go mod init go-auth-jwt'
                        sh "go mod tidy"   
                    }

                    sh "go clean"
                    sh "go build -o go-auth-jwt"
                // }
            }
        }

        stage('Deploy') {
            steps {
                // Here you can add steps to deploy your Golang application (e.g., copying to a server)
                // For example, you can use SCP or SSH to copy the binary to a remote server
                // sh "scp myapp user@your-server:/path/to/deploy/"
                // Your test steps here
                sh './go-auth-jwt'

                // Find and terminate the Go binary process
                sh 'pkill -f go-auth-jwt'

                sh "./go-auth-jwt &"
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
