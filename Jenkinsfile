pipeline {
    agent any

    environment {
        APP_ENV = 'staging'
        APP_PORT = "${env.ALGORITHMS_API_APP_PORT}"
        APP_BASE_URL = "${env.ALGORITHMS_API_APP_BASE_URL}"
        K6_PORT = "${env.ALGORITHMS_API_K6_PORT}"
        K6_BASE_URL = "${env.ALGORITHMS_API_K6_BASE_URL}"
        K6_OPTIONS_FILE = "${env.ALGORITHMS_API_K6_OPTIONS_FILE}"
    }

    stages {
        stage('Checkout Code') {
            steps {
                echo 'Pulling code from GitHub...'
                git branch: 'main',
                    url: 'https://github.com/vanessahoamea/algorithms-api',
                    credentialsId: 'github-token'
            }
        }

        stage('Verify Environment Variables') {
            steps {
                echo "APP_ENV: ${APP_ENV}"
                echo "APP_PORT: ${APP_PORT}"
                echo "APP_BASE_URL: ${APP_BASE_URL}"
                echo "K6_PORT: ${K6_PORT}"
                echo "K6_BASE_URL: ${K6_BASE_URL}"
                echo "K6_OPTIONS_FILE: ${K6_OPTIONS_FILE}"
            }
        }

        stage('Verify Docker') {
            steps {
                echo 'Checking if Docker is available...'
                sh 'docker --version'
            }
        }

        stage('Run Docker Compose') {
            steps {
                echo 'Building Docker images for Algorithms API server and performance tests...'
                sh 'make compose-up'
            }
        }
    }

    post {
        always {
            echo 'Pipeline completed. Stopping containers...'
            sh 'make compose-down'
            sh 'docker ps'
        }

        failure {
            echo 'Performance tests FAILED!'
            sh 'docker-compose logs -f'
        }
    }
}