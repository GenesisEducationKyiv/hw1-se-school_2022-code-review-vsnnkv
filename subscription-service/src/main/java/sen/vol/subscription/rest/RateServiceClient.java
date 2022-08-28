package sen.vol.subscription.rest;

import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import sen.vol.subscription.service.HTTPResponseDTO;

@FeignClient(name = "rate-service")
public interface RateServiceClient {

    @RequestMapping(value = "/rate/", method = RequestMethod.GET)
    String getRateBtsToUah();
}
