package sen.vol.rate.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import sen.vol.rate.service.RateService;

@RestController
public class RateController {

    private final RateService rateService;

    public RateController(RateService rateService) {
        this.rateService = rateService;
    }


    @GetMapping("/api/rate")
    public ResponseEntity<Integer> getRate(){
        return rateService.getRate();
    }
}
