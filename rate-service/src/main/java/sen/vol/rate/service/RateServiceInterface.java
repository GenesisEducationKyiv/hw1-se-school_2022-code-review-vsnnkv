package sen.vol.rate.service;

import org.springframework.http.ResponseEntity;

public interface RateServiceInterface {
    ResponseEntity<Integer> getRate();
}
