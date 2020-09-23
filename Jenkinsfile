pipeline {
  environment {
    registryAPI = "redshoore/mukgo-api"
    registryDB = "redshoore/mukgo-db"
    registryLog = "redshoore/mukgo-log"

    imageAPI = ""
    imageDB = ""
    imageLog = ""

    registryCredential = "redshore_docker_hub"
  }

  agent any

  // If anything fails, the whole Pipeline stops.
  stages {
    stage('Build Image') {
      steps {
        script {
          imageAPI = docker.build registryAPI + ":$BUILD_NUMBER" 
          imageDB = docker.build registryDB + ":$BUILD_NUMBER" 
          imageLog = docker.build registryLog + ":$BUILD_NUMBER" 
        }
      }
    }

    stage('Deploy Image') {
      steps {
        script {
          docker.withRegistry("", registryCredential) {
            imageAPI.push()
            imageDB.push()
            imageLog.push()
          }
        }
      }
    }

    stage('Clean Up') {
      steps {
        sh "docker rmi $registryAPI:$BUILD_NUMBER"
        sh "docker rmi $registryDB:$BUILD_NUMBER"
        sh "docker rmi $registryLog:$BUILD_NUMBER"
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