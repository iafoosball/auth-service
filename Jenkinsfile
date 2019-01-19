pipeline {
    agent any
    environment {
        GOOGLE_OAUTH2_CLIENT_ID = credentials('google-oauth2-client-id')
        GOOGLE_OAUTH2_CLIENT_SECRET = credentials('google-oauth2-client-secret')
    }
    stages {
        stage('Prepare stag env') {
            steps {
                sh "./delete-kong.sh"
                sh "docker rm -f auth-service auth-redis"
            }
        }
        stage('Build') {
            steps {
                sh "docker-compose build"
            }
        }
        stage('Deploy') {
            steps {
                sh "docker-compose -p auth-stag up -d --force-recreate"
                sh "./create-kong.sh"
            }
        }
    }
    post {
       always {
            sh "docker-compose down --rmi='all'"
            sh "docker system prune -f"
        }
    }
}
