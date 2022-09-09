package sen.vol.subscription.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import sen.vol.subscription.model.EmailRequestDTO;
import sen.vol.subscription.service.SubscriptionEmailService;
import sen.vol.subscription.service.SubscriptionService;

import java.io.IOException;

@RestController
public class SubscriptionController {

    @Autowired
    private final SubscriptionEmailService subscriptionEmailService;

    public SubscriptionController(SubscriptionEmailService subscriptionEmailService) {
        this.subscriptionEmailService = subscriptionEmailService;
    }

    @PostMapping("/api/subscribe")
    public ResponseEntity<String> subscribeEmail(@RequestBody EmailRequestDTO emailRequestDTO) {

        return subscriptionEmailService.saveEmail(emailRequestDTO.getEmail());
    }

    @GetMapping("/api/sendEmails")
    public ResponseEntity<String> sendEmails() throws IOException {
        return subscriptionEmailService.createResponse();
    }

}
