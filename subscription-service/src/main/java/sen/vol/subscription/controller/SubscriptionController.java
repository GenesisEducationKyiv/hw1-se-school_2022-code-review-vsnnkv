package sen.vol.subscription.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import sen.vol.subscription.model.EmailRequestDTO;
import sen.vol.subscription.service.SubscriptionServiceInterface;

import java.io.IOException;

@RestController
public class SubscriptionController {

    private final SubscriptionServiceInterface subscriptionServiceInterface;

    public SubscriptionController(SubscriptionServiceInterface subscriptionServiceInterface) {
        this.subscriptionServiceInterface = subscriptionServiceInterface;
    }

    @PostMapping("/api/subscribe")
    public ResponseEntity<String> subscribeEmail(@RequestBody EmailRequestDTO emailRequestDTO) {

        return subscriptionServiceInterface.saveEmail(emailRequestDTO.getEmail());
    }

    @GetMapping("/api/sendEmails")
    public ResponseEntity<String> sendEmails() throws IOException {
        return subscriptionServiceInterface.createResponse();
    }

}
