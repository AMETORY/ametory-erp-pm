import { UserModel } from "./user";

export interface CompanyModel {
  id?: string;
  name?: string;
  logo?: string;
  cover?: string;
  legal_entity?: string;
  email?: string;
  phone?: string;
  fax?: string;
  address?: string;
  contact_person?: string;
  contact_person_position?: string;
  tax_payer_number?: string;
  user_id?: string;
  user?: UserModel;
  status?: string;
  setting?: CompanySetting 
}

export interface CompanySetting {
  company_id?: string
  gemini_api_key?: string
  whatsapp_web_host?: string
  whatsapp_web_mock_number?: string
  whatsapp_web_is_mocked?: boolean
}