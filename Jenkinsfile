pipeline {
    agent {
        kubernetes {
            yamlFile 'docker_pod.yaml'
        }
    }
    environment {
        DOCKER_IMAGE = 'nmmalmquist/beckett-ridge-family-medicine-website'
        DOCKER_HUB_USERNAME = "nmmalmquist"
        DOCKER_HUB_ACCESS_TOKEN = credentials('docker-hub-token-nmmalmquist')
        SSH_REMOTE = 'nmalmquist@99.45.53.198'
        SSH_PORT = '1234'
    }
    stages {
        stage('Build Docker Image') {
            steps {
                container('docker'){
                    echo "---> Starting to build image! ${env.DOCKER_IMAGE}"
                    sh "docker build . -t ${env.DOCKER_IMAGE} --no-cache"
                    sh "docker login -u ${env.DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_ACCESS_TOKEN}"
                    sh "docker push ${env.DOCKER_IMAGE}"
                    echo "---> Finished building image and pushed to Dockerhub"
                }
            }
        }
        stage('Rollout Restart Kubernetes With New Image') {
            steps {
                script {
                    withCredentials([sshUserPrivateKey(credentialsId: "ssh-key-for-raspberry-pi", keyFileVariable: 'keyfile')]) {
                        // StrictHostKeyChecking not make sure no prompting for fingerprint
                        sh "ssh -o \"StrictHostKeyChecking no\" -v -i ${keyfile} ${env.SSH_REMOTE} -p ${env.SSH_PORT} 'kubectl rollout restart deployment/diatte-website-release-helm-chart-malmquist-fe -n diatte-website'"
                    }
                }
            }
        }
        
    }
}