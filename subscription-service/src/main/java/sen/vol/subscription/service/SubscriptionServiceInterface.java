package sen.vol.subscription.service;

import org.springframework.http.ResponseEntity;

import java.io.IOException;

public interface SubscriptionServiceInterface {

    ResponseEntity<String> saveEmail(String email);

    ResponseEntity<String> createResponse() throws IOException;
}
