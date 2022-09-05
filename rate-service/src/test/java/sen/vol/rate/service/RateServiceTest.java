package sen.vol.rate.service;

import org.junit.Assert;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.*;
import org.mockito.junit.MockitoJUnitRunner;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.client.RestTemplate;


@RunWith(MockitoJUnitRunner.class)
public class RateServiceTest {

    @Mock
    private RestTemplate restTemplate;


    @InjectMocks
    @Spy
    private RateService rateService = new RateService(new RestTemplateBuilder(),
            "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah");

    @Test
    public void RateServiceTest() throws Exception {

        ResponseEntity<Integer> responseEntity = rateService.getRate();
        Assert.assertEquals(200, responseEntity.getStatusCodeValue());
    }

    @Test
    public void RateServiceTest_nullResponse(){
        rateService = new RateService(new RestTemplateBuilder(),
                "https://api.coicko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah");
        ResponseEntity<Integer> responseEntity = rateService.getRate();
        Assert.assertEquals(500, responseEntity.getStatusCodeValue());
    }

    @Test(expected = Exception.class)
    public void RateServiceTest_Exception(){
        rateService = new RateService(null,
                null);
        rateService.getRate();


    }
}
