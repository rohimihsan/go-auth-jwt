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
                echo "run"
                sh "export JENKINS_NODE_COOKIE=dontKillMe; nohup ./go-auth-jwt > go-auth-jwt.log 2>&1 &"
            }
        }
    }

    post {
        always {
            // Wait for a few seconds to ensure the Go binary is started properly
            sleep time: 10, unit: 'SECONDS'

            // Optionally, you can include further steps here to validate the Go binary's behavior or perform other tasks.
        }
        success {
            echo 'Build and deployment successful!'
        }
        failure {
            echo 'Build or deployment failed!'
        }
    }
}
