package sen.vol.subscription.service;

import org.springframework.stereotype.Service;

import java.io.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

@Service
public class FileService implements FileServiceInterface {
    private File emails;

    public void saveEmailToFile(String email) throws IOException {
        PrintWriter writer = new PrintWriter(new BufferedWriter(new FileWriter(emails, true)));

        writer.println(email);
        writer.flush();
        writer.close();

    }

    public void checkEmailsFile(){
        emails = new File("emails.txt");
        try {
            if (!emails.exists()) {
                emails.createNewFile();
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public Boolean lookIfEmailInTheList(String email) throws FileNotFoundException {
        Scanner scanner = new Scanner(emails);

        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            if (line.equals(email)) {
                scanner.close();
                return true;
            }
        }

        return false;
    }

    public List<String> getEmails() throws IOException {
        BufferedReader reader = new BufferedReader(new FileReader(emails));

        List<String> results = new ArrayList<>();
        String line = reader.readLine();
        while (line != null) {
            results.add(line);
            line = reader.readLine();
        }
        return results;
    }
}
