package sen.vol.rate.service;

import com.netflix.ribbon.proxy.annotation.Http;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

@Service
public class RateService {

    private final RestTemplate restTemplate;
    private final String coinGeckoSimpleApiUrl;

    public RateService(RestTemplateBuilder restTemplateBuilder,
                       @Value("${external.api.coingecko.simple.price}") String coinGeckoSimpleApiUrl) {
        this.restTemplate = restTemplateBuilder.build();
        this.coinGeckoSimpleApiUrl = coinGeckoSimpleApiUrl;
    }

    public HTTPResponseDTO<String> getRateBtsToUah(){

        try {
            RateResponseDTO responseData = this.restTemplate.getForObject(coinGeckoSimpleApiUrl, RateResponseDTO.class);
            if (responseData == null) {
                return new HTTPResponseDTO<>("Помилка сервера", 500);
            }
            return new HTTPResponseDTO<>(String.valueOf(responseData.getBitcoin().getUah()));
        } catch (Exception exception){
            return new HTTPResponseDTO<>("Помилка сервера", 500);
        }
    }
}
