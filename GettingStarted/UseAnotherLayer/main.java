package GettingStarted.UseAnotherLayer;

public class main {
    UserRepository userRepository = new UserRepositoryImpl();
    UserInteractor userInteractor = new UserInteractorImpl(userRepository);

    UserDTO userDTO = userInteractor.getUser(1);
    List<UserDTO> users = userInteractor.getAllUsers();

    UserDTO newUser = new UserDTO("John", "john@example.com");
    userInteractor.createUser(newUser);

    UserDTO userToUpdate = userInteractor.getUser(1);
    userToUpdate.setName("John Doe");
    userToUpdate.setEmail("johndoe@example.com");
    userInteractor.updateUser(userToUpdate);

    userInteractor.deleteUser(1);

}
