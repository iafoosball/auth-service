pipeline {
    agent any
    environment {
        COMPOSE_PROJECT_NAME = "${env.JOB_NAME}-${env.BUILD_ID}"
        COMPOSE_FILE = "docker-compose.yml"
        GOOGLE_OAUTH2_CLIENT_ID = credentials('google-oauth2-client-id')
        GOOGLE_OAUTH2_CLIENT_SECRET = credentials('google-oauth2-client-secret')
    }
    stages {
        stage('build') {
            steps {
                sh "docker-compose build --pull"
            }
        }
        stage('deploy') {
            steps {
                sh "docker-compose up -d"
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
