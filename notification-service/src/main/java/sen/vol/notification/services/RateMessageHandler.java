package sen.vol.notification.services;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Service;
import sen.vol.notification.config.RabbitMQConfig;

@Service
//@RequiredArgsConstructor
public class RateMessageHandler {
    private final JavaMailSender javaMailSender;
    private final Logger LOGGER = LoggerFactory.getLogger(RateMessageHandler.class);

    @Autowired
    public RateMessageHandler(JavaMailSender javaMailSender) {
        this.javaMailSender = javaMailSender;
    }

    @RabbitListener(queues = RabbitMQConfig.QUEUE_RATE)
    public void receive(Message message) throws JsonProcessingException {
        System.out.println(message);
        byte[] body = message.getBody();
        String jsonBody = new String(body);
        ObjectMapper objectMapper = new ObjectMapper();
        RateResponseDTO rateResponseDTO = objectMapper.readValue(jsonBody, RateResponseDTO.class);
        LOGGER.info(rateResponseDTO.toString());

        SimpleMailMessage mailMessage = new SimpleMailMessage();
        mailMessage.setTo(rateResponseDTO.getMail());
        mailMessage.setFrom("mailsender222@ukr.net");

        mailMessage.setSubject("Курс BTC до UAH");
        mailMessage.setText("курс BTC до UAH відповідно до данних сайту coingecko.com складає: "
                + rateResponseDTO.getPrice());

        try {
            javaMailSender.send(mailMessage);
        } catch (Exception exception) {
            LOGGER.error(rateResponseDTO.toString());
        }
    }
}
