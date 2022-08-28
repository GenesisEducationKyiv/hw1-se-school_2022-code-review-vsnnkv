package sen.vol.subscription.service;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class HTTPResponseDTO<T> {

    private T message;

    private Integer code = 200;

    public HTTPResponseDTO(T message) {
        this.message = message;
    }
}
