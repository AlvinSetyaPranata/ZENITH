

export interface RoleType {
    Id: number;
    Name: string;
    Permissions: [] | null;
    DateCreated: string
}
export interface UserType {
    id: number;
    email: string;
    role: RoleType
}