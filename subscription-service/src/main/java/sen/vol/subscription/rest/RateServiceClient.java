package sen.vol.subscription.rest;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@FeignClient(name = "rate-service")
public interface RateServiceClient {

    @GetMapping(value = "/api/rate")
    Integer getRateBtsToUah();
}
