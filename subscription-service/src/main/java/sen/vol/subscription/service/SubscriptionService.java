package sen.vol.subscription.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import sen.vol.subscription.model.RateResponceDTO;
import sen.vol.subscription.rest.RateServiceClient;

import java.io.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

@Service
public class SubscriptionService {

    private static final String TOPIC_EXCHANGE_RATE = "js.rate.notify.exchange";

    private static final String ROUTING_KEY_RATE = "js.key.rate";
    private final RabbitTemplate rabbitTemplate;

    private final RateServiceClient rateServiceClient;


   private File emails;

    @Autowired
    public SubscriptionService(RabbitTemplate rabbitTemplate, RateServiceClient rateServiceClient) {
        this.rabbitTemplate = rabbitTemplate;
        this.rateServiceClient = rateServiceClient;
        checkEmailsFile();
    }

    public ResponseEntity<String> saveEmail(String email){
        try{
            if (lookIfEmailInTheList(email)) {
                return ResponseEntity.status(409).body("E-mail  вже є в базі данних");
            }
            saveEmailToFile(email);
            return ResponseEntity.ok("E-mail додано");
        } catch (Exception exception){
            exception.printStackTrace();
            return ResponseEntity.status(500).body("Помилка сервера");
        }
    }

    private void saveEmailToFile(String email) throws IOException {
        PrintWriter writer = new PrintWriter(new BufferedWriter(new FileWriter(emails, true)));

        writer.println(email);
        writer.flush();
        writer.close();

    }

    private void checkEmailsFile(){
        emails = new File("emails.txt");
        try {
            if (!emails.exists()) {
                emails.createNewFile();
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    private Boolean lookIfEmailInTheList(String email) throws FileNotFoundException {
        Scanner scanner = new Scanner(emails);

        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            if (line.equals(email)) {
                scanner.close();
                return true;
            }
        }

        return false;
    }

    public List<String> getEmails() throws IOException {
        BufferedReader reader = new BufferedReader(new FileReader(emails));

        List<String> results = new ArrayList<>();
        String line = reader.readLine();
        while (line != null) {
            results.add(line);
            line = reader.readLine();
        }
        return results;
    }

    public ResponseEntity<String> createResponse() throws IOException {
        try {
            List<String> emailsList = getEmails();

            Integer response = rateServiceClient.getRateBtsToUah();

            for (String email : emailsList) {
                RateResponceDTO rateResponseDTO = new RateResponceDTO(response, email);

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
