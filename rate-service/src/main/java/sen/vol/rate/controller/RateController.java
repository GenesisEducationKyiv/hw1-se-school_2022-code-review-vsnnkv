package sen.vol.rate.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import sen.vol.rate.service.RateServiceInterface;

@RestController
public class RateController {

    private final RateServiceInterface rateServiceInterface;

    public RateController(RateServiceInterface rateServiceInterface) {
        this.rateServiceInterface = rateServiceInterface;
    }


    @GetMapping("/api/rate")
    public ResponseEntity<Integer> getRate() {
        return rateServiceInterface.getRate();
    }
}
