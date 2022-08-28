package sen.vol.subscription.service;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
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


    File emails;

    @Autowired
    public SubscriptionService(RabbitTemplate rabbitTemplate, RateServiceClient rateServiceClient) {
        this.rabbitTemplate = rabbitTemplate;
        this.rateServiceClient = rateServiceClient;
        checkEmailsFile();
    }

    public HTTPResponseDTO<String> saveEmail(String email) {
        try {
            if (lookIfEmailInTheList(email)) {
                return new HTTPResponseDTO<>("E-mail  вже є в базі данних", 409);
            }
            saveEmailToFile(email);
            return new HTTPResponseDTO<>("E-mail додано");
        } catch (Exception exception) {
            exception.printStackTrace();
            return new HTTPResponseDTO<>("Помилка сервера", 500);
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

    public HTTPResponseDTO<String> createResponse() throws IOException {
        try {
            List<String> emailsList = getEmails();

            String response = rateServiceClient.getRateBtsToUah();

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

            return new HTTPResponseDTO<>("E-mailʼи відправлено");
        } catch (Exception exception){
            exception.printStackTrace();
            return new HTTPResponseDTO<>("Помилка сервера", 500);
        }
    }

}
