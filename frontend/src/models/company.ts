import { UserModel } from "./user";

export interface CompanyModel {
  id?: string;
  name: string;
  logo: string;
  cover: string;
  legal_entity: string;
  email: string;
  phone: string;
  fax: string;
  address: string;
  contact_person: string;
  contact_person_position: string;
  tax_payer_number?: string;
  user_id?: string;
  user?: UserModel;
  status: string;
}
