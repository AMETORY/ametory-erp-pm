import { CompanyModel } from "./company";
import { RoleModel } from "./role";
import { TeamModel } from "./team";
import { UserModel } from "./user";

export interface MemberModel {
    id?: string;
    company_id?: string;
    company?: CompanyModel;
    user_id: string;
    user: UserModel;
    role_id?: string;
    role?: RoleModel;
    team_id?: string;
    team?: TeamModel;
}

