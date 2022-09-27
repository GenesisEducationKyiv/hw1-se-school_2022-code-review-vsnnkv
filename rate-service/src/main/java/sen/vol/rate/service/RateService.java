package sen.vol.rate.service;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import sen.vol.rate.model.RateResponseDTO;

@Service
public class RateService implements RateServiceInterface {

    private final RestTemplate restTemplate;
    private final String coinGeckoSimpleApiUrl;

    public RateService(RestTemplateBuilder restTemplateBuilder,
                       @Value("${external.api.coingecko.simple.price}") String coinGeckoSimpleApiUrl){
        this.restTemplate = restTemplateBuilder.build();
        this.coinGeckoSimpleApiUrl = coinGeckoSimpleApiUrl;
    }

    public ResponseEntity<Integer> getRate(){
        try {
            RateResponseDTO responseEntity = this.restTemplate.getForObject(
                    coinGeckoSimpleApiUrl, RateResponseDTO.class);
            if (responseEntity == null){
                return  new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
            }
            return  ResponseEntity.ok(responseEntity.getBitcoin().getUah());

        } catch (Exception exception){
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}
