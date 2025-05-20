import { HR, Tooltip } from "flowbite-react";
import { useContext, useEffect, useState, type FC } from "react";
import { AiOutlineDashboard } from "react-icons/ai";
import {
  BsAsterisk,
  BsGear,
  BsInstagram,
  BsKanban,
  BsPeople,
  BsTag,
  BsTelegram,
  BsWhatsapp,
} from "react-icons/bs";
import { GoTasklist } from "react-icons/go";
import { HiOutlineChat } from "react-icons/hi";
import { HiOutlineInboxArrowDown } from "react-icons/hi2";
import { LuContact2, LuLink2, LuPowerOff } from "react-icons/lu";
import { SiGoogleforms } from "react-icons/si";
import { useNavigate } from "react-router-dom";
import { CollapsedContext } from "../contexts/CollapsedContext";
import {
  getInboxMessagesCount,
  getSentMessagesCount,
} from "../services/api/inboxApi";
import { asyncStorage } from "../utils/async_storage";
import {
  LOCAL_STORAGE_COMPANIES,
  LOCAL_STORAGE_COMPANY_ID,
  LOCAL_STORAGE_DEFAULT_CHANNEL,
  LOCAL_STORAGE_DEFAULT_INSTAGRAM_SESSION,
  LOCAL_STORAGE_DEFAULT_TELEGRAM_SESSION,
  LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION,
  LOCAL_STORAGE_TOKEN,
} from "../utils/constants";
import { MdOutlineAssistant } from "react-icons/md";
import { MemberContext, ProfileContext } from "../contexts/ProfileContext";
import { PiBroadcast } from "react-icons/pi";
import { IoPricetag } from "react-icons/io5";
import { RiShoppingBagLine } from "react-icons/ri";
import { TbTemplate } from "react-icons/tb";

interface SidebarProps {}

const Sidebar: FC<SidebarProps> = ({}) => {
  const { member } = useContext(MemberContext);
  const { collapsed } = useContext(CollapsedContext);
  const [mounted, setMounted] = useState(false);
  const [inboxUnreadCount, setInboxUnreadCount] = useState(0);
  const [sentUnreadCount, setSentUnreadCount] = useState(0);
  const [indexUnreadChat, setIndexUnreadChat] = useState(0);
  const [waUnreadChat, setWaUnreadChat] = useState(0);

  useEffect(() => {
    setMounted(true);
  }, []);

  const nav = useNavigate();

  useEffect(() => {
    if (mounted) {
      getInboxMessagesCount()
        .then((resp: any) => setInboxUnreadCount(resp.data))
        .catch(console.error);
      getSentMessagesCount()
        .then((resp: any) => setSentUnreadCount(resp.data))
        .catch(console.error);
    }
  }, [mounted]);

  const handleNavigation =
    (path: string) => (event: React.MouseEvent<HTMLAnchorElement>) => {
      event.preventDefault();
      nav(path);
    };

  const checkPermission = (permission: string) => {
    if (member?.role?.permission_names) {
      return member.role.permission_names.includes(permission);
    }
    return false;
  };
  return (
    <div className="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800 flex flex-col">
      <ul className="space-y-2 font-medium flex-1 h-[calc(100vh-100px)] overflow-y-auto">
        <li className="" style={{}}>
          <span
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
            onClick={handleNavigation("/")}
          >
            <Tooltip content="Dashboard" placement="bottom">
              <AiOutlineDashboard />
            </Tooltip>
            {!collapsed && <span className="ms-3">Dashboard</span>}
          </span>
        </li>
        <HR />
        <li className="text-xs text-gray-300" style={{}}>
          Feature
        </li>
        <li className="" style={{}}>
          <span
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
            onClick={handleNavigation("/task")}
          >
            <Tooltip content="Task">
              <GoTasklist />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Task</span>
            )}
          </span>
        </li>
        {checkPermission("project_management:project:read") && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={handleNavigation("/project")}
            >
              <Tooltip content="Project">
                <BsKanban />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">Project</span>
              )}
            </span>
          </li>
        )}
        <li className="" style={{}}>
          <span
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
            onClick={handleNavigation("/inbox")}
          >
            <Tooltip content="Inbox">
              <HiOutlineInboxArrowDown />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Inbox</span>
            )}
            {!collapsed && inboxUnreadCount + sentUnreadCount > 0 && (
              <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                {inboxUnreadCount + sentUnreadCount}
              </span>
            )}
          </span>
        </li>
        <li className="" style={{}}>
          <span
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
            onClick={async () => {
              let channelID = await asyncStorage.getItem(
                LOCAL_STORAGE_DEFAULT_CHANNEL
              );
              if (channelID) {
                nav(`/chat/${channelID}`);
              } else {
                nav(`/chat`);
              }
            }}
          >
            <Tooltip content="Chat">
              <HiOutlineChat />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Chat</span>
            )}
            {!collapsed && indexUnreadChat > 0 && (
              <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                {indexUnreadChat}
              </span>
            )}
          </span>
        </li>
        {checkPermission("inventory:product:read") && (
          <li className=" cursor-pointer" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={handleNavigation("/product")}
            >
              <Tooltip content="Product">
                <RiShoppingBagLine />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">Product</span>
              )}
            </span>
          </li>
        )}
        <HR />
        <li className="text-xs text-gray-300" style={{}}>
          Omni Channel
        </li>
        {checkPermission("customer_relationship:whatsapp:read") && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={async () => {
                let sessionID = await asyncStorage.getItem(
                  LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION
                );
                if (sessionID) {
                  nav(`/whatsapp/${sessionID}`);
                } else {
                  nav(`/whatsapp`);
                }
              }}
            >
              <Tooltip content="Whatsapp">
                <BsWhatsapp />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">Whatsapp</span>
              )}
              {!collapsed && waUnreadChat > 0 && (
                <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                  {waUnreadChat}
                </span>
              )}
            </span>
          </li>
        )}
        {process.env.REACT_APP_TELEGRAM_ENABLED && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={async () => {
                let sessionID = await asyncStorage.getItem(
                  LOCAL_STORAGE_DEFAULT_TELEGRAM_SESSION
                );
                if (sessionID) {
                  nav(`/telegram/${sessionID}`);
                } else {
                  nav(`/telegram`);
                }
              }}
            >
              <Tooltip content="telegram">
                <BsTelegram />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">Telegram</span>
              )}
              {!collapsed && waUnreadChat > 0 && (
                <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                  {waUnreadChat}
                </span>
              )}
            </span>
          </li>
        )}
        {process.env.REACT_APP_INSTAGRAM_ENABLED && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={async () => {
                let sessionID = await asyncStorage.getItem(
                  LOCAL_STORAGE_DEFAULT_INSTAGRAM_SESSION
                );
                if (sessionID) {
                  nav(`/instagram/${sessionID}`);
                } else {
                  nav(`/instagram`);
                }
              }}
            >
              <Tooltip content="instagram">
                <BsInstagram />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">Instagram</span>
              )}
              {!collapsed && waUnreadChat > 0 && (
                <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                  {waUnreadChat}
                </span>
              )}
            </span>
          </li>
        )}

        <li className="" style={{}}>
          <span
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
            onClick={async () => {
              nav(`/broadcast`);
            }}
          >
            <Tooltip content="Broadcast">
              <PiBroadcast />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Broadcast</span>
            )}
          </span>
        </li>
        <HR />
        <li className="text-xs text-gray-300" style={{}}>
          Preferences
        </li>
        <li className="" style={{}}>
          <span
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
            onClick={handleNavigation("/member")}
          >
            <Tooltip content="Member">
              <BsPeople />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Member</span>
            )}
          </span>
        </li>
        {checkPermission("customer_relationship:form:read") && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={handleNavigation("/form")}
            >
              <Tooltip content="Form">
                <SiGoogleforms />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">Form</span>
              )}
            </span>
          </li>
        )}
        {checkPermission("project_management:project:update") && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={handleNavigation("/task-attribute")}
            >
              <Tooltip content="Task Attribute">
                <BsAsterisk />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">
                  Task Attribute
                </span>
              )}
            </span>
          </li>
        )}
        {member?.role?.is_super_admin && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={handleNavigation("/gemini-agent")}
            >
              <Tooltip content="Gemini Agent">
                <MdOutlineAssistant />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">
                  Gemini Agent
                </span>
              )}
            </span>
          </li>
        )}
        {checkPermission("contact:customer:read") && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={handleNavigation("/contact")}
            >
              <Tooltip content="Contact">
                <LuContact2 />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">Contact</span>
              )}
            </span>
          </li>
        )}
        {member?.role?.is_super_admin && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={handleNavigation("/connection")}
            >
              <Tooltip content="Connection">
                <LuLink2 />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">
                  Connection
                </span>
              )}
            </span>
          </li>
        )}
        <li className="" style={{}}>
          <span
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
            onClick={handleNavigation("/tag")}
          >
            <Tooltip content="Tag">
              <BsTag />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Tag</span>
            )}
          </span>
        </li>
        <li className="" style={{}}>
          <span
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
            onClick={handleNavigation("/template")}
          >
            <Tooltip content="Template">
              <TbTemplate />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Template</span>
            )}
          </span>
        </li>
        {member?.role?.is_super_admin && (
          <li className="" style={{}}>
            <span
              className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group cursor-pointer"
              onClick={handleNavigation("/setting")}
            >
              <Tooltip content="Setting">
                <BsGear />
              </Tooltip>
              {!collapsed && (
                <span className="flex-1 ms-3 whitespace-nowrap">Setting</span>
              )}
            </span>
          </li>
        )}
      </ul>
      <div
        className="flex flex-row gap-2 items-center cursor-pointer hover:font-bold px-2"
        onClick={async () => {
          await asyncStorage.removeItem(LOCAL_STORAGE_TOKEN);
          await asyncStorage.removeItem(LOCAL_STORAGE_COMPANIES);
          await asyncStorage.removeItem(LOCAL_STORAGE_COMPANY_ID);
          window.location.reload();
        }}
      >
        <Tooltip content="Logout">
          <LuPowerOff />
        </Tooltip>
        {!collapsed && <span>Logout</span>}
      </div>
    </div>
  );
};
export default Sidebar;
