package sen.vol.subscription.model;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
public class RateResponseDTO {

    private Integer price;

    private String mail;
}
