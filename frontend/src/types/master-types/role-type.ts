
export interface UserType {
    email: string;
    password: string;
}

export interface RoleType {
    Id: number;
    Name: string;
    Permissions: [] | null;
    DateCreated: string
}