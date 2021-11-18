FROM jenkins/jenkins
ENV JAVA_OPTS -Djenkins.install.runSetupWizard=false
ENV CASC_JENKINS_CONFIG /var/jenkins_home/casc.yaml
COPY plugins.txt /usr/share/jenkins/ref/plugins.txt
COPY jobs/ /var/jenkins_home/jobs/
COPY scriptApproval.xml /var/jenkins_home/scriptApproval.xml
USER root
RUN chown -R jenkins:jenkins /var/jenkins_home/scriptApproval.xml
RUN chown -R jenkins:jenkins /var/jenkins_home/jobs
USER jenkins
RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/ref/plugins.txt
COPY casc.yaml /var/jenkins_home/casc.yaml
