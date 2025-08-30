export type User = {
    id: number;
    name: string;
    password: string;
    profile_picture: string;
    email: string;
    role: "user" | "admin";
};
