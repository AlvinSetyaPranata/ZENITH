import { createContext, ParentProps } from "solid-js";
import { createStore } from "solid-js/store";
import { AuthContextType, AuthStoreType } from "~/types/auth-types/store";
import { LoginWithNimService, LogoutService } from "~/services/auth/login";
import { LoginWithNimCredentialsType } from "~/types/auth-types/credentials";
import { UserType } from "~/types/master-types/role-type";

export const AuthContext = createContext<AuthContextType>();

export function AuthProvider(props: ParentProps) {
  const [user, setUser] = createStore<AuthStoreType>({
    user: null,
    token: null,
  });

  const storeUserData = (user: UserType, token: string) => {
    // this function should be run before profile data is stored in memory
    // for persistent storage we use session storage with value that has been obfuscate before

    if (typeof window === "undefined") return


    if (user) {
      // TODO: Add obfuscate method to perform
      sessionStorage.setItem("user", JSON.stringify(user));
    }

    if (token) {
      // TODO: Add obfuscate method to perform
      sessionStorage.setItem("token", token);
    }

    setUser({
      user: user,
      token: token,
    });
  };

  const getUserData = () => {

    if (typeof window === "undefined") {
      console.error("Componenet didnt mount yet!")
      return null
    }

    if (user.user) return user;
    
    
    
    const userData = sessionStorage.getItem("user");
    const token = sessionStorage.getItem("token");
    
    // Profile data in session is missing causing
    if (!userData) {
      sessionStorage.clear()
      return null
    };
    
    // TODO: Add deobfuscate method to extract user profile to inserted in memory
    
    const parsedUserData = JSON.parse(userData)

    setUser({ user: parsedUserData, token: token });

    return user
  };

  const login = async (credentials: LoginWithNimCredentialsType) => {
    const {response, status} = await LoginWithNimService(credentials);

    if (!status) {
      return false;
    }

    const data = await response.json();

    storeUserData(data.data, data.access_token);
    return true;
  };


  const logout = async () => {
    const status =  await LogoutService(user.user?.id)

    if (window != undefined && status) {
      sessionStorage.clear()
    }

    return status
  }

  const contextPayload: AuthContextType = {
    auth: user,
    login: login,
    logout: logout,
    storeUserData: storeUserData,
    getUserData: getUserData,
  };

  return (
    <AuthContext.Provider value={contextPayload}>
      {props.children}
    </AuthContext.Provider>
  );
}
