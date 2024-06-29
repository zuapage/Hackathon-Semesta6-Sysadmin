pipeline {
    agent any
    environment {
        SEMESTA_APP1 = "poetryndream/semesta-app1:${env.BUILD_ID}"
        SEMESTA_APP2 = "poetryndream/semesta-app2:${env.BUILD_ID}"
    }

    stages {
        stage('Clone Repo from Github') {
            steps {
                git branch: 'master', credentialsId: 'jenkins-master', url: 'git@github.com:zuapage/Hackathon-Semesta6-Sysadmin.git'
            }
        }
        
        stage('Remove Latest Build') {
            steps {
                sh '''
                docker images poetryndream/semesta-app* --format "{{.Repository}}:{{.Tag}}" | xargs -r docker rmi
                '''
            }
        }

        stage('Build App') {
            parallel {
                stage('build semesta-app1') {
                    steps {
                        script {
                            docker.build(SEMESTA_APP1, "-f docker/app1/Dockerfile .")
                        }
                    }
                }

                stage('build semesta-app2') {
                    steps {
                        script {
                            docker.build(SEMESTA_APP2, "-f docker/app2/Dockerfile .")
                        }
                    }
                }
            }
        }

        stage('Push Docker Image') {
            parallel {
                stage('push image semesta-app1') {
                    steps {
                        withCredentials([usernamePassword(credentialsId: 'docker-hub', 
                                usernameVariable: 'USERNAME', 
                                passwordVariable: 'PASSWORD')]) {
                            sh('docker login -u ${USERNAME} -p ${PASSWORD}')
                            sh('docker push ${SEMESTA_APP1}')
                        }
                    }
                }
                stage('push image semesta-app2') {
                    steps {
                        withCredentials([usernamePassword(credentialsId: 'docker-hub', 
                                usernameVariable: 'USERNAME', 
                                passwordVariable: 'PASSWORD')]) {
                            sh('docker login -u ${USERNAME} -p ${PASSWORD}')
                            sh('docker push ${SEMESTA_APP2}')
                        }
                    }
                }
            }
        }


        stage('Deploy Application to Production') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker-hub', 
                        usernameVariable: 'USERNAME', 
                        passwordVariable: 'PASSWORD')]) {
                    sh ('docker login -u $USERNAME -p ${PASSWORD}')
                    sh """
                    sed -i 's|image: poetryndream/semesta-app1.*|image: ${SEMESTA_APP1}|' docker/compose_apps/docker-compose.yml
                    sed -i 's|image: poetryndream/semesta-app2.*|image: ${SEMESTA_APP2}|' docker/compose_apps/docker-compose.yml
                    cat docker/compose_apps/docker-compose.yml   
                    """
                    withCredentials([sshUserPrivateKey(credentialsId: 'jenkins-master', keyFileVariable: 'JK_AUTH')]) {
                       sh '''                  
                        scp -i ${JK_AUTH} -r -P 1566 docker/compose_apps bunnies@192.168.56.121:/usr/src/compose
                        ssh -i ${JK_AUTH} -p 1566 bunnies@192.168.56.121 "cd /usr/src/compose/compose_apps/ && docker compose down"
                        ssh -i ${JK_AUTH} -p 1566 bunnies@192.168.56.121 "cd /usr/src/compose/compose_apps/ && docker compose up -d"
                       '''
                    }

                }
            }
        }
 
    }
}
                
