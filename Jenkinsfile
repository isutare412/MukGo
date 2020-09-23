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
          imageAPI = docker.build(registryAPI + ":$BUILD_NUMBER", "-f ${WORKSPACE}/api_server.Dockefile .")
          imageDB = docker.build(registryDB + ":$BUILD_NUMBER", "-f ${WORKSPACE}/db_server.Dockefile .")
          imageLog = docker.build(registryLog + ":$BUILD_NUMBER", "-f ${WORKSPACE}/log_server.Dockefile .")
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