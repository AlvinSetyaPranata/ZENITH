import { RoleType } from "../master-types/role-type";

export interface LoginResponseType {
    access_token: string;
    data: {
        id: number;
        role_id: RoleType
    };
    messege: string;
}