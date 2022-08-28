FROM openjdk:8-jdk-alpine
EXPOSE 8082
COPY build/libs/subscription-service-1.0-SNAPSHOT.jar .
RUN apk add --no-cash bash
COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh
ENTRYPOINT ["/wait-for.sh", "rate-service:8081", "--timeout=1800", "--", "java", "-jar", "subscription-service-1.0-SNAPSHOT.jar"]
