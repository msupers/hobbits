pipeline{
		agent { docker 'registry.cmri.cn/zhcsep/dind-nodejs-maven:20190916' }
		stages {
			stage('Scm') {
                steps{
				    withCredentials([usernamePassword(credentialsId: 'git-token', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
				        sh "rm -rf {{.GitProject}};git clone http://zhangmiaoyjy:$PASSWORD@dev.cmri.cn/gitlab/{{.GitGroup}}/{{.GitProject}}.git"
				       }
				}
			}
			stage('sonar check') {
                steps{
                    dir("{{.GitProject}}"){
                    sh "mvn clean ; mvn compile"
                    sh 'mvn sonar:sonar \
                          -Dsonar.host.url=http://10.2.40.44:9000 \
                          -Dsonar.login=4de60cbde64da7376422c765afa1b99b3db75bfc'
                    }
                }
              }
        }
}
</script>
<sandbox>true</sandbox>
</definition>
<triggers/>
<disabled>false</disabled>
</flow-definition>