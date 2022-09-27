package sen.vol.subscription.service;

import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.List;

public interface FileServiceInterface {

    void saveEmailToFile(String email) throws IOException;

    void checkEmailsFile();

    Boolean lookIfEmailInTheList(String email) throws FileNotFoundException;

    List<String> getEmails() throws IOException;
}
