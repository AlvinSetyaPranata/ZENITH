import { UserType } from "../master-types/role-type";
import { LoginWithNimCredentialsType } from "./credentials";



export interface AuthStoreType {
    user: UserType | null;
    token?: string;
}


export interface AuthContextType {
    auth: AuthStoreType;
    login: (credentials: LoginWithNimCredentialsType) => Promise<boolean>;
}

