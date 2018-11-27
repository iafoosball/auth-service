pipeline {
    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    stages {
        stage('build') {
            steps {
                sh "docker-compose build --pull"
            }
        }
        stage('deploy') {
            steps {
                sh "docker-compose up"
                sh "./create-kong.sh"
            }
        }
    }
    post {
       always {
            sh "docker-compose down -v --rmi 'all'"
            sh "./delete-kong.sh"
        }
    }
}
