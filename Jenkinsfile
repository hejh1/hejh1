pipeline {
    agent {
        kubernetes {
            defaultContainer 'jnlp'
            yaml '''
apiVersion: v1
kind: Pod
metadata:
labels:
  component: ci
spec:
  # Use service account that can deploy to all namespaces
  serviceAccountName: jenkins-wi-jenkins
  containers:
  - name: docker
    image: docker:latest
    command:
    - cat
    tty: true
    volumeMounts:
    - mountPath: /var/run/docker.sock
      name: docker-sock
  volumes:
    - name: docker-sock
      hostPath:
        path: /var/run/docker.sock
'''
        }
    }
    options {
        timeout(time: 1, unit: 'HOURS')
    }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build') {
            steps {
                echo 'Building 1..'
                script {
                    docker.withRegistry('https://us-central1-docker.pkg.dev/','glowing-sprite-347007') {
                        container('docker') {
                            testImage = docker.build("glowing-sprite-347007/quickstart-docker-repo/firebase:${env.BUILD_TAG}", "./docker/")
                            testImage.push()
                            // sh """
                            // docker build -t youpi/firebase:${env.BUILD_TAG} ./docker/
                            // """
                        }
                    }
                }
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
