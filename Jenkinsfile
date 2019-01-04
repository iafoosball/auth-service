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
                sh "./delete-kong.sh"
                sh "docker-compose down --rmi='all'"
                sh "docker-compose build --pull"
            }
        }
        stage('deploy') {
            steps {
                catchError {
                    sh "docker-compose up -d --force-recreate"
                }
                sh "./create-kong.sh"
            }
        }
    }
    post {
       always {
            sh "docker system prune -f"
        }
        
       failure {
            sh "docker rm -f iafoosball_auth-service iafoosball_redis"
        }
    }
}
