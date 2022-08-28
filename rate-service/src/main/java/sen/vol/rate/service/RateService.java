package sen.vol.rate.service;

import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
public class RateService {

    private final RestTemplate restTemplate;


    public RateService(RestTemplateBuilder restTemplateBuilder) {
        this.restTemplate = restTemplateBuilder.build();
    }

    public HTTPResponseDTO<String> getRateBtsToUah(){

        try {
            String url = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah";
            RateResponseDTO responseData = this.restTemplate.getForObject(url, RateResponseDTO.class);
            if (responseData == null) {
                return new HTTPResponseDTO<>("Помилка сервера", 500);
            }
            return new HTTPResponseDTO<>(String.valueOf(responseData.getBitcoin().getUah()));
        } catch (Exception exception){
            return new HTTPResponseDTO<>("Помилка сервера", 500);
        }
    }
}
