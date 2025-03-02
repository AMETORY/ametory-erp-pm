import { MemberModel } from "./member";

export interface TeamModel {
    id: string; // BaseModel field
    name: string;
    members: MemberModel[];
}

