package sen.vol.subscription.service;

import org.springframework.http.ResponseEntity;

import java.io.IOException;

public interface SubscriptionEmailService {

    ResponseEntity<String> saveEmail(String email);

    ResponseEntity<String> createResponse() throws IOException;
}
