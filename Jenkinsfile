pipeline {
    agent any
    environment {
        GOOGLE_OAUTH2_CLIENT_ID = credentials('google-oauth2-client-id')
        GOOGLE_OAUTH2_CLIENT_SECRET = credentials('google-oauth2-client-secret')
    }
    stages {
        stage('Deploy to stag') {
            steps {
                catchError {
                    sh "docker-compose down"
                    sh "docker-compose -p auth-stag up -d --build --force-recreate"
                    sh "./create-kong.sh"
                }
            }
        }
    }
    post {
        always {
            sh "docker system prune -f"
        }
        failure {
            sh "./delete-kong.sh"
            sh "docker rm -f auth-service auth-redis"
        }
    }
}
