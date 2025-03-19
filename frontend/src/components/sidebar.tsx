import { HR, Tooltip } from "flowbite-react";
import { useContext, useEffect, useState, type FC } from "react";
import { AiOutlineDashboard } from "react-icons/ai";
import { BsAsterisk, BsGear, BsKanban, BsPeople } from "react-icons/bs";
import { GoTasklist } from "react-icons/go";
import { HiOutlineChat } from "react-icons/hi";
import { HiOutlineInboxArrowDown } from "react-icons/hi2";
import { LuContact2, LuPowerOff } from "react-icons/lu";
import { SiGoogleforms } from "react-icons/si";
import { useNavigate } from "react-router-dom";
import { CollapsedContext } from "../contexts/CollapsedContext";
import { getInboxMessagesCount, getSentMessagesCount } from "../services/api/inboxApi";
import { asyncStorage } from "../utils/async_storage";
import { LOCAL_STORAGE_COMPANIES, LOCAL_STORAGE_COMPANY_ID, LOCAL_STORAGE_DEFAULT_CHANNEL, LOCAL_STORAGE_TOKEN } from "../utils/constants";
import { MdOutlineAssistant } from "react-icons/md";

interface SidebarProps {}

const Sidebar: FC<SidebarProps> = ({}) => {
  const { collapsed, setCollapsed } = useContext(CollapsedContext);
  const [mounted, setMounted] = useState(false);
  const [inboxUnreadCount, setInboxUnreadCount] = useState(0);
  const [sentUnreadCount, setSentUnreadCount] = useState(0);
  const [indexUnreadChat, setIndexUnreadChat] = useState(0);

  useEffect(() => {
    setMounted(true)
  }, []);
  
  const nav = useNavigate();

  useEffect(() => {
    if (mounted) {
      getInboxMessagesCount().then((resp: any) => setInboxUnreadCount(resp.data)).catch(console.error)
      getSentMessagesCount().then((resp: any) => setSentUnreadCount(resp.data)).catch(console.error)
    }
  
  }, [mounted]);

  const handleNavigation =
    (path: string) => (event: React.MouseEvent<HTMLAnchorElement>) => {
      event.preventDefault();
      nav(path);
    };

  return (
    <div className="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800 flex flex-col">
      <ul className="space-y-2 font-medium flex-1">
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/")}
          >
            <Tooltip content="Dashboard">
              <AiOutlineDashboard />
            </Tooltip>
            {!collapsed && <span className="ms-3">Dashboard</span>}
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/task")}
          >
            <Tooltip content="Task">
              <GoTasklist />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Task</span>
            )}
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/project")}
          >
            <Tooltip content="Project">
              <BsKanban />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Project</span>
            )}
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
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
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={async () => {
              let channelID = await asyncStorage.getItem(LOCAL_STORAGE_DEFAULT_CHANNEL)
              if (channelID) {
                nav(`/chat/${channelID}`)
              } else {
                nav(`/chat`)

              }
            }}
          >
            <Tooltip content="Chat">
              <HiOutlineChat />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Chat</span>
            )}
            {!collapsed && indexUnreadChat  > 0 && (
              <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                {indexUnreadChat }
              </span>
            )}
          </a>
        </li>
        <HR />
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/member")}
          >
            <Tooltip content="Member">
              <BsPeople />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Member</span>
            )}
           
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/form")}
          >
            <Tooltip content="Form">
              <SiGoogleforms />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Form</span>
            )}
           
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/task-attribute")}
          >
            <Tooltip content="Task Attribute">
              <BsAsterisk />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Task Attribute</span>
            )}
           
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/gemini-agent")}
          >
            <Tooltip content="Gemini Agent">
              <MdOutlineAssistant />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Gemini Agent</span>
            )}
           
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/contact")}
          >
            <Tooltip content="Contact">
              <LuContact2 />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Contact</span>
            )}
           
          </a>
        </li>
        <li className="" style={{ }}>
          <a
            href="#"
            className="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group"
            onClick={handleNavigation("/setting")}
          >
            <Tooltip content="Setting">
              <BsGear />
            </Tooltip>
            {!collapsed && (
              <span className="flex-1 ms-3 whitespace-nowrap">Setting</span>
            )}
           
          </a>
        </li>
       
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

