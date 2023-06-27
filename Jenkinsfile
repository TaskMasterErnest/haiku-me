pipeline {

  agent any
  tools {
    go 'Go'
  }

  stages {

    stage("unit testing") {
      steps {
        echo 'UNIT TEST EXECUTION STARTED'
        sh 'make unit-tests'
      }
    }

    // stage("run application") {
    //   steps {
    //     echo 'RUNNING AN APPLICATION'
    //     sh 'make run'
    //   }
    // }

    stage("build database and app images") {
      steps {
        parallel (
          "Build database image": {
            withDockerRegistry(credentialsId: 'docker-creds', url: "") {
              sh 'make build-db'
            }
          },
          "Build app image": {
            withDockerRegistry(credentialsId: 'docker-creds', url: "") {
              sh 'make build-app'
            }
          }
        )
      }
    }

    // stage("run the application") {
    //   steps {
    //     withDockerRegistry(credentialsId: 'docker-creds', url: "") {
    //       sh 'make compose'
    //     }
    //   }
    // }

    stage("Kubernetes - Deploy database") {
      withKubeConfig(credentialsId: 'kube-config', restrictKubeConfigAccess: false, serverUrl: '') {
        sh "bash kubernetes-yaml/deploy-db.sh"
      }
    }

    stage("Kubernetes - Deploy application") {
      withKubeConfig(credentialsId: 'kube-config', restrictKubeConfigAccess: false, serverUrl: '') {
        sh "bash kubernetes-yaml/deploy-app.sh"
      }
    }

  }

}