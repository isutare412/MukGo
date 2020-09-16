pipeline {
    // Lets Jenkins use Docker for us later.
    agent any

    // If anything fails, the whole Pipeline stops.
    stages {
        stage('Build Server') {
            // Use golang image
            agent { docker { image 'golang:1.15.2' } }

            steps {
                sh 'make api'
            }
        }

        stage('Test Server') {
            // Use golang image
            agent { docker { image 'golang:1.15.2' } }

            steps {
                sh 'go test ./server/...'
            }
        }
    }

    post {
        always {
            // Clean up our workspace.
            deleteDir()
        }
    }
} 