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
                catchError(buildResult: 'FAILURE', stageResult: 'SUCCESS') {
                    sh 'make compose-up'
                }
            }
        }

        stage('Publish HTML Report') {
            steps {
                echo 'Checking if k6 HTML report is available...'
                sh 'ls docker/dashboard.html'

                echo 'Publishing k6 dashboard...'
                publishHTML([
                    allowMissing: false,
                    alwaysLinkToLastBuild: true,
                    keepAll: true,
                    reportDir: 'docker',
                    reportFiles: 'dashboard.html',
                    reportName: 'k6 Report'
                ])
            }
        }
    }

    post {
        always {
            echo 'Pipeline completed. Stopping containers...'
            sh 'make compose-down'
            sh 'docker ps'
        }

        success {
            echo 'Pushing images to Docker Hub...'
            sh 'make compose-push'
        }

        failure {
            echo 'Performance tests FAILED!'
            sh 'docker compose logs -f'
        }
    }
}