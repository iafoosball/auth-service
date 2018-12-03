pipeline {
    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
    }
    stages {
        stage('build') {
            steps {
                sh "./delete-kong.sh"
                sh "docker rm -f iafoosball_auth-service iafoosball_redis"
                sh "docker-compose down --rmi='all'"
                sh "docker-compose build --pull"
            }
        }
        stage('deploy') {
            steps {
                sh "docker-compose up -d --force-recreate"
                sh "./create-kong.sh"
            }
        }
    }
    post {
       always {
            sh "docker system prune -f"
        }
    }
}
