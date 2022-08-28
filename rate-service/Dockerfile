FROM openjdk:8-jdk-alpine
EXPOSE 8081
COPY build/libs/rate-service-1.0-SNAPSHOT.jar .
RUN  apk add --no-cash bash
COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh
ENTRYPOINT ["/wait-for.sh", "gateway:8989", "--timeout=120", "--", "java", "-jar", "rate-service-1.0-SNAPSHOT.jar"]
