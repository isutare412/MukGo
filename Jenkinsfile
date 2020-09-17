pipeline {
  // Lets Jenkins use Docker for us later.
  agent any

  // If anything fails, the whole Pipeline stops.
  stages {
    stage('Build Server') {
      // Use golang image
      agent {
        docker {
          image 'golang:1.15.2'
          args '-e XDG_CACHE_HOME=/tmp/.cache'
        }
      }

      steps {
          sh 'make server'
      }
    }

    stage('Test Server') {
      // Use golang image
      agent {
        docker {
          image 'golang:1.15.2'
          args '-e XDG_CACHE_HOME=/tmp/.cache'
        }
      }

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