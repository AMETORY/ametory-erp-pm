import { CompanyModel } from "./company";

export interface RapidApiPluginModel {
  id: string;
  name: string;
  key: string;
  url: string;
  is_active: boolean;
  rapid_api_endpoints: RapidApiEndpointModel[];
}

export interface RapidApiEndpointModel {
  id: string;
  title: string;
  key: string;
  method: string;
  params: string;
  url: string;
  rapid_api_plugin_id: string;
  ParamData: null;
}


export interface CompanyRapidApiPluginModel {
  company_id: string;
  company: CompanyModel;
  rapid_api_plugin_id: string;
  rapid_api_plugin: RapidApiPluginModel;
  rapid_api_key: string;
  rapid_api_host: string;
}
