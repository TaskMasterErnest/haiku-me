pipeline {

  agent any

  stages {

    stage("unit testing") {
      steps {
        echo 'UNIT TEST EXECUTION STARTED'
        sh 'make unit-tests'
      }
    }

    stage("run application") {
      steps {
        echo 'RUNNING AN APPLICATION'
        sh 'make run'
      }
    }

  }

}