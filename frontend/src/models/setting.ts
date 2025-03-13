export interface SettingModel {
    id: string;
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
    status: string;
    bank_account: string | null;
    bank_code: string | null;
    beneficiary_name: string | null;
}
