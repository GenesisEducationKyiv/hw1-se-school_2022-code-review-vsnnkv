package sen.vol.notification.config;

import org.springframework.amqp.core.*;
import org.springframework.amqp.rabbit.annotation.EnableRabbit;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
@EnableRabbit
public class RabbitMQConfig {
    public static final String QUEUE_RATE = "js.rate.notify";
    private static final String TOPIC_EXCHANGE_RATE = "js.rate.notify.exchange";
    private static final String ROUTING_KEY_RATE = "js.key.rate";

    @Autowired
    private AmqpAdmin amqpAdmin;

    @Bean
    public TopicExchange rateExchange(){
        return new TopicExchange(TOPIC_EXCHANGE_RATE);
    }

    @Bean
    public Queue queueRate(){
        return new Queue(QUEUE_RATE);
    }

    @Bean
    public Binding rateBinding(){
        return BindingBuilder.bind(queueRate())
                .to(rateExchange())
                .with(ROUTING_KEY_RATE);
    }
}
