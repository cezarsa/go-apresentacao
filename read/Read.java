import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Read {
    public static void main(String[] args) throws IOException {
        String content = new String(Files.readAllBytes(Paths.get("../arquivo.txt")));
        System.out.println(content);
    }
}