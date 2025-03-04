import { useEffect, useState, type FC } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "../pages/Home";
import ProjectDetail from "../pages/ProjectDetail";
import { LoadingContext } from "../contexts/LoadingContext";
import PrivateRoute from "./PrivateRoute";
import PublicRoute from "./PublicRoute";
import { CollapsedContext } from "../contexts/CollapsedContext";
import { CompanyModel } from "../models/company";
import { CompaniesContext, CompanyIDContext } from "../contexts/CompanyContext";
import { asyncStorage } from "../utils/async_storage";
import {
  LOCAL_STORAGE_COMPANY_ID,
  LOCAL_STORAGE_TOKEN,
} from "../utils/constants";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { UserModel } from "../models/user";
import { ProfileContext } from "../contexts/ProfileContext";
import AcceptInvitation from "../pages/AcceptInvitation";

interface AppRoutesProps {
  token?: string | null;
}

const AppRoutes: FC<AppRoutesProps> = ({ token }) => {
  const [loading, setLoading] = useState(false);
  const [isWsConnected, setWsConnected] = useState(false);
  const [wsMsg, setWsMsg] = useState<string | null>(null);
  const [collapsed, setCollapsed] = useState(false);
  const [companyID, setCompanyID] = useState<string | null>(null);
  const [companies, setCompanies] = useState<CompanyModel[] | null>(null);
  const [profile, setProfile] = useState<UserModel | null>(null);

  useEffect(() => {
    // console.log("token", token);
    asyncStorage.getItem(LOCAL_STORAGE_COMPANY_ID).then((id) => {
      setCompanyID(id);
    });
  }, []);
  return (
    <LoadingContext.Provider value={{ loading, setLoading }}>
      <CollapsedContext.Provider value={{ collapsed, setCollapsed }}>
        <ProfileContext.Provider value={{ profile, setProfile }}>
          <CompaniesContext.Provider value={{ companies, setCompanies }}>
            <CompanyIDContext.Provider
              value={{
                companyID,
                setCompanyID: (val) => {
                  asyncStorage.setItem(LOCAL_STORAGE_COMPANY_ID, val);
                  setCompanyID(val);
                },
              }}
            >
              <WebsocketContext.Provider
                value={{ isWsConnected, setWsConnected, wsMsg, setWsMsg }}
              >
                <BrowserRouter>
               
                  {token && <PrivateRoute />}
                  {!token && <PublicRoute />}
                </BrowserRouter>
              </WebsocketContext.Provider>
            </CompanyIDContext.Provider>
          </CompaniesContext.Provider>
        </ProfileContext.Provider>
      </CollapsedContext.Provider>
    </LoadingContext.Provider>
  );
};
export default AppRoutes;
