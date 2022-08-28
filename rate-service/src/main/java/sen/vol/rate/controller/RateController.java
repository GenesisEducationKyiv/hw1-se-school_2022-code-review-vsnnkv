package sen.vol.rate.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import sen.vol.rate.service.HTTPResponseDTO;
import sen.vol.rate.service.RateService;

@RestController
public class RateController {

    @Autowired
    private final RateService rateService;

    public RateController(RateService rateService) {
        this.rateService = rateService;
    }

    @GetMapping()
    public ResponseEntity<String> getRate(){
        HTTPResponseDTO<String> rateResponse = rateService.getRateBtsToUah();
        return ResponseEntity.status(rateResponse.getCode()).body(rateResponse.getMessage());
    }
}
