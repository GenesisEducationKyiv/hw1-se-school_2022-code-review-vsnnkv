package sen.vol.subscription.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import sen.vol.subscription.model.RateResponseDTO;
import sen.vol.subscription.rest.RateServiceClient;

import java.io.*;
import java.util.List;

@Service
public class SubscriptionService implements SubscriptionServiceInterface {

    private static final String TOPIC_EXCHANGE_RATE = "js.rate.notify.exchange";

    private static final String ROUTING_KEY_RATE = "js.key.rate";
    private final RabbitTemplate rabbitTemplate;

    private final RateServiceClient rateServiceClient;

    private final FileServiceInterface fileServiceInterface;


    @Autowired
    public SubscriptionService(RabbitTemplate rabbitTemplate, RateServiceClient rateServiceClient,
                               FileServiceInterface fileServiceInterface) {
        this.rabbitTemplate = rabbitTemplate;
        this.rateServiceClient = rateServiceClient;
        this.fileServiceInterface = fileServiceInterface;
    }

    public ResponseEntity<String> saveEmail(String email){
        fileServiceInterface.checkEmailsFile();
        try{
            if (fileServiceInterface.lookIfEmailInTheList(email)) {
                return ResponseEntity.status(409).body("E-mail  вже є в базі данних");
            }
            fileServiceInterface.saveEmailToFile(email);
            return ResponseEntity.ok("E-mail додано");
        } catch (Exception exception){
            exception.printStackTrace();
            return ResponseEntity.status(500).body("Помилка сервера");
        }
    }

    public ResponseEntity<String> createResponse() throws IOException {
        try {
            List<String> emailsList = fileServiceInterface.getEmails();

            Integer response = rateServiceClient.getRateBtsToUah();

            for (String email : emailsList) {
                RateResponseDTO rateResponseDTO = new RateResponseDTO(response, email);

                ObjectMapper objectMapper = new ObjectMapper();
                try {
                    rabbitTemplate.convertAndSend(TOPIC_EXCHANGE_RATE, ROUTING_KEY_RATE,
                            objectMapper.writeValueAsString(rateResponseDTO));
                } catch (JsonProcessingException e) {
                    e.printStackTrace();
                }
            }

            return ResponseEntity.ok("E-mailʼи відправлено");
        } catch (Exception exception){
            exception.printStackTrace();
            return ResponseEntity.status(500).body("Помилка сервера");
        }
    }

}
