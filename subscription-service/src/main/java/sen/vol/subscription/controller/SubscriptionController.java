package sen.vol.subscription.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import sen.vol.subscription.controller.dto.EmailRequestDTO;
import sen.vol.subscription.service.HTTPResponseDTO;
import sen.vol.subscription.service.SubscriptionService;

import java.io.IOException;

@RestController
public class SubscriptionController {

    @Autowired
    private final SubscriptionService subscriptionService;

    public SubscriptionController(SubscriptionService subscriptionService) {
        this.subscriptionService = subscriptionService;
    }

    @PostMapping("/subscribe")
    public ResponseEntity<String> subscribeEmail(@RequestBody EmailRequestDTO emailRequestDTO){

        HTTPResponseDTO<String> httpResponseDTO = subscriptionService.saveEmail(emailRequestDTO.getEmail());
        return ResponseEntity.status(httpResponseDTO.getCode()).body(httpResponseDTO.getMessage());
    }

    @GetMapping("/sendEmails")
    public ResponseEntity<String> sendEmails() throws IOException {
        HTTPResponseDTO<String> response = subscriptionService.createResponse();
        return ResponseEntity.status(response.getCode()).body(response.getMessage());
    }

}
