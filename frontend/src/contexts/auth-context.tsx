import { createContext, ParentProps } from "solid-js";
import { createStore } from "solid-js/store";
import { AuthContextType, AuthStoreType } from "~/types/auth-types/store";
import { LoginWithNimService } from "~/services/auth/login";
import { LoginWithNimCredentialsType } from "~/types/auth-types/credentials";



export const AuthContext = createContext<AuthContextType>();


export function AuthProvider(props: ParentProps) {
    const [user, setUser] = createStore<AuthStoreType>({
        user: null,
    })
    
    const login = async (credentials: LoginWithNimCredentialsType) => {
        const response = await LoginWithNimService(credentials)

        if (response.status == 401) {
            // TODO: Refetch to update access token
            console.log("Unauthenticated")
            return false
        }

        const data = await response.json()
        
        setUser({
            user: data.data,
            token: data.access_token
        })
        return true
        
    }

    const contextPayload: AuthContextType = {
        auth: user,
        login: login
    }

    return (
        <AuthContext.Provider value={contextPayload}>
            {props.children}
        </AuthContext.Provider>
    )   
}
