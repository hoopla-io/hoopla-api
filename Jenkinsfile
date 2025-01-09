pipeline {
    agent any

    environment {
        GITHUB_CREDENTIALS_ID = 'a1a2d14a-f3a9-4a0f-87fc-a4b466453e7f'
        SERVER_USERNAME = credentials('5f60409a-e740-4e64-af81-10e5c804acb5')
        SERVER_PASSWORD = credentials('ea53a4b6-5073-4bd9-8f81-4026f54b5b23')
        SERVER_IP = credentials('9515f99b-360f-41d4-b7fa-0276e2a16c9c')
    }

    stages {
        stage('Deploy') {
            steps {
                script {
                    sshagent([GITHUB_CREDENTIALS_ID]) {
                        sh """
                            sshpass -p "${SERVER_PASSWORD}" ssh -A -tt -o StrictHostKeyChecking=no ${SERVER_USERNAME}@${SERVER_IP} "
                                echo 'Connected to server';
                                cd /opt/qahvazor-api/www;
                                git pull origin main;
                                docker compose up -d --build;
                            "
                        """
                    }
                }
            }
        }
    }

    post {
        always {
            echo "Cleaning up..."
            cleanWs()
        }
        success {
            echo "Pipeline succeeded!"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}