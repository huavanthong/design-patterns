// Step 1: Create a abstract class
// ===> Đây tập trung những method common, mà class khác đều phải có
interface ImageReader {
    DecodedImage getDecodeImage();
}

// Step 2: Create an algorithm to decode image
// ===> Hiểu rằng, việc một class decode, ta phải hiểu được
// có bao nhiêu parameters, các methods common nào cần thiết.
// Có trong class này, để ta provide cho những Decoder còn lại
// handle và implement riêng cho nó.
class DecodedImage {
    private String image;

    // Constructor
    public DecodedImage(String image) {
        this.image = image;
    }

    @Override
    public String toString() {
        return image + ": is decoded";
    }
}

// Step 3: Create a class to implement all interfaces
class GifReader implements ImageReader {
    private DecodedImage decodedImage;

    // Constructor
    public GifReader(String image) {
        this.decodedImage = new DecodedImage(image);
    }

    // Overview method decode for this class
    @Override
    public DecodedImage getDecodeImage() {
        return decodedImage;
    }
}

// Step 4: With new decoder,
//         we create another class to implement all interfaces
class JpegReader implements ImageReader {
    private DecodedImage decodedImage;

    public JpegReader(String image) {
        decodedImage = new DecodedImage(image);
    }

    @Override
    public DecodedImage getDecodeImage() {
        return decodedImage;
    }
}

public class FactoryMethodDemo {
    public static void main(String[] args) {
        // Initialize value
        DecodedImage decodedImage;
        ImageReader reader = null;
        String image = args[0];

        // Get format image from user
        String format = image.substring(image.indexOf('.') + 1, (image.length()));
        
        // Select extension image to start decoding image
        if (format.equals("gif")) {
            reader = new GifReader(image);
        }
        if (format.equals("jpeg")) {
            reader = new JpegReader(image);
        }
        assert reader != null;

        // Handle image
        decodedImage = reader.getDecodeImage();
        System.out.println(decodedImage);
    }
}