package sen.vol.subscription.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import sen.vol.subscription.model.EmailRequestDTO;
import sen.vol.subscription.service.SubscriptionService;

import java.io.IOException;

@RestController
public class SubscriptionController {

    @Autowired
    private final SubscriptionService subscriptionService;

    public SubscriptionController(SubscriptionService subscriptionService) {
        this.subscriptionService = subscriptionService;
    }

//    @PostMapping("/api/subscribe")
//    public ResponseEntity<String> subscribeEmail(@RequestBody EmailRequestDTO emailRequestDTO){
//
//        HTTPResponseDTO<String> httpResponseDTO = subscriptionService.saveEmail(emailRequestDTO.getEmail());
//        return ResponseEntity.status(httpResponseDTO.getCode()).body(httpResponseDTO.getMessage());
//    }
        @PostMapping("/api/subscribe")
    public ResponseEntity<String> subscribeEmail(@RequestBody EmailRequestDTO emailRequestDTO){

        return subscriptionService.saveEmail(emailRequestDTO.getEmail());
    }


//    @GetMapping("/api/sendEmails")
//    public ResponseEntity<String> sendEmails() throws IOException {
//        HTTPResponseDTO<String> response = subscriptionService.createResponse();
//        return ResponseEntity.status(response.getCode()).body(response.getMessage());
//    }
   @GetMapping("/api/sendEmails")
    public ResponseEntity<String> sendEmails() throws IOException {
//        HTTPResponseDTO<String> response = subscriptionService.createResponse();
//        return ResponseEntity.status(response.getCode()).body(response.getMessage());
        return subscriptionService.createResponse();
    }

}
