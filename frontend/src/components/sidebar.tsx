import { useContext, useEffect, useState, type FC } from "react";
import { AiOutlineDashboard } from "react-icons/ai";
import { BsKanban, BsPeople } from "react-icons/bs";
import { HiOutlineInboxArrowDown } from "react-icons/hi2";
import { SlPeople } from "react-icons/sl";
import { asyncStorage } from "../utils/async_storage";
import { LOCAL_STORAGE_COMPANIES, LOCAL_STORAGE_COMPANY_ID, LOCAL_STORAGE_TOKEN } from "../utils/constants";
import { LuPowerOff } from "react-icons/lu";
import { CollapsedContext } from "../contexts/CollapsedContext";
import { useNavigate } from "react-router-dom";
import { Tooltip } from "flowbite-react";
import { GoTasklist } from "react-icons/go";
import { getInboxMessagesCount } from "../services/api/inboxApi";

interface SidebarProps {}

const Sidebar: FC<SidebarProps> = ({}) => {
  const { collapsed, setCollapsed } = useContext(CollapsedContext);
  const [mounted, setMounted] = useState(false);
  const [inboxUnreadCount, setInboxUnreadCount] = useState(0);

  useEffect(() => {
    setMounted(true)
  }, []);
  
  const nav = useNavigate();

  useEffect(() => {
    if (mounted) {
      getInboxMessagesCount().then((resp: any) => setInboxUnreadCount(resp.data)).catch(console.error)
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
            {!collapsed && inboxUnreadCount > 0 && (
              <span className="inline-flex items-center justify-center w-3 h-3 p-3 ms-3 text-sm font-medium text-blue-800 bg-blue-100 rounded-full dark:bg-blue-900 dark:text-blue-300">
                {inboxUnreadCount}
              </span>
            )}
          </a>
        </li>
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

