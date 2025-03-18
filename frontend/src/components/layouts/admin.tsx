import { useContext, useEffect, useState, type FC } from "react";
import {
  HiArrowSmRight,
  HiChartPie,
  HiInbox,
  HiShoppingBag,
  HiTable,
  HiUser,
} from "react-icons/hi";
import Topnav from "../topnav";
import { CollapsedContext } from "../../contexts/CollapsedContext";
import Sidebar from "../sidebar";
import { getProfile } from "../../services/api/authApi";
import {
  ActiveCompanyContext,
  CompaniesContext,
  CompanyIDContext,
} from "../../contexts/CompanyContext";
import useWebSocket, { ReadyState } from "react-use-websocket";
import { asyncStorage } from "../../utils/async_storage";
import {
  LOCAL_STORAGE_COMPANY_ID,
  LOCAL_STORAGE_TOKEN,
} from "../../utils/constants";
import { WebsocketContext } from "../../contexts/WebsocketContext";
import { MemberContext, ProfileContext } from "../../contexts/ProfileContext";
import { Toaster } from "react-hot-toast";
import Loading from "../Loading";
import { LoadingContext } from "../../contexts/LoadingContext";
import { getSetting } from "../../services/api/commonApi";

interface AdminLayoutProps {
  children: React.ReactNode;
}

const AdminLayout: FC<AdminLayoutProps> = ({ children }) => {
  const { activeCompany, setActiveCompany } = useContext(ActiveCompanyContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const { member, setMember } = useContext(MemberContext);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { loading, setLoading } = useContext(LoadingContext);
  const [socketUrl, setSocketUrl] = useState(``);
  const { collapsed, setCollapsed } = useContext(CollapsedContext);
  const { companyID, setCompanyID } = useContext(CompanyIDContext);
  const { companies, setCompanies } = useContext(CompaniesContext);
  const [token, setToken] = useState("");
  const { sendMessage, sendJsonMessage, lastMessage, readyState } =
    useWebSocket(socketUrl, {
      onMessage(event) {
        // console.log("Received message:", event.data);
        setWsMsg(JSON.parse(event.data));
      },
      onOpen() {
        console.log("Connected to the web socket");
        setWsConnected(true);
      },
      onClose() {
        console.log("Disconnected from the web socket");
        setWsConnected(false);
      },
      queryParams: {
        token: token,
      },
    });

  const connectionStatus = {
    [ReadyState.CONNECTING]: "Connecting",
    [ReadyState.OPEN]: "Open",
    [ReadyState.CLOSING]: "Closing",
    [ReadyState.CLOSED]: "Closed",
    [ReadyState.UNINSTANTIATED]: "Uninstantiated",
  }[readyState];

  useEffect(() => {
    getProfile().then((res: any) => {
      setProfile(res.user);
      setCompanies(res.user.companies);
      setMember(res.member);
    });
    getSetting()
      .then((val: any) => setActiveCompany(val.data))
      .catch((err) => {});
    asyncStorage.getItem(LOCAL_STORAGE_TOKEN).then((token) => {
      setToken(token);
      asyncStorage.getItem(LOCAL_STORAGE_COMPANY_ID).then((id) => {
        if (!id) return;
        let url = `${process.env.REACT_APP_BASE_WS_URL}/api/v1/ws/${id}`;
        setSocketUrl(url);
      });
    });
  }, []);

  const renderSelectCompany = () => {
    return (
      <div className="flex flex-row items-center justify-center h-full   w-full">
        <div className="bg-white p-4 rounded-md shadow-md">
          <h2 className="text-lg font-bold">Select Company First</h2>
          <p>You need to select company first to access this page</p>
        </div>
      </div>
    );
  };
  return (
    <div className="w-screen h-screen  flex flex-col">
      {loading && <Loading />}
      <Toaster position="bottom-left" reverseOrder={false} />
      <Topnav />
      <div className="flex flex-row flex-1">
        {companyID && (
          <aside
            style={{
              width: collapsed ? 65 : 300,
              top: 65,
              height: "calc(100% - 65px)",
            }}
            className=" bg-red-50 h-full fixed left-0 "
          >
            <Sidebar />
          </aside>
        )}
        {companyID ? (
          <div
            style={{
              width: collapsed ? "calc(100% - 65px)" : "calc(100% - 300px)",
              height: "calc(100% - 65px)",
              left: collapsed ? 65 : 300,
              top: 65,
            }}
            className=" fixed  "
          >
            {children}
          </div>
        ) : (
          renderSelectCompany()
        )}
      </div>
    </div>
  );
};
export default AdminLayout;
