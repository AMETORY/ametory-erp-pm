import { createContext } from "react";
import { CompanyModel } from "../models/company";

export const CompanyIDContext = createContext<{
  companyID: string | null;
  setCompanyID: (companyID: string | null) => void;
}>({
  companyID: null,
  setCompanyID: () => {},
});

export const CompaniesContext = createContext<{
  companies: CompanyModel[] | null;
  setCompanies: (companies: CompanyModel[] | null) => void;
}>({
  companies: null,
  setCompanies: () => {},
});
