import React, { useContext } from "react";
import { CollapsedContext } from "../contexts/CollapsedContext";
import Logo from "./logo";
import { CompaniesContext, CompanyIDContext } from "../contexts/CompanyContext";
import { Avatar, Dropdown } from "flowbite-react";
import { ProfileContext } from "../contexts/ProfileContext";
import { initial } from "../utils/helper";

interface TopnavProps {}

const Topnav: React.FC<TopnavProps> = () => {
  const { companies, setCompanies } = useContext(CompaniesContext);
  const { companyID, setCompanyID } = useContext(CompanyIDContext);
  const { collapsed, setCollapsed } = useContext(CollapsedContext);
  const { profile, setProfile } = useContext(ProfileContext);

  return (
    <nav className="fixed top-0 z-20 w-full bg-white border-b border-gray-200 dark:bg-gray-800 dark:border-gray-700">
      <div className="px-3 py-3 lg:px-5 lg:pl-3">
        <div className="flex items-center justify-between">
          <div className="flex items-center justify-start rtl:justify-end ">
            <button
              onClick={() => setCollapsed(!collapsed)}
              data-drawer-target="logo-sidebar"
              data-drawer-toggle="logo-sidebar"
              aria-controls="logo-sidebar"
              type="button"
              className="inline-flex items-center p-2 text-sm mr-4 text-gray-500 rounded-lg  hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
            >
              <span className="sr-only">Open sidebar</span>
              <svg
                className="w-6 h-6"
                aria-hidden="true"
                fill="currentColor"
                viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  clipRule="evenodd"
                  fillRule="evenodd"
                  d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"
                />
              </svg>
            </button>
            <Logo />
          </div>
          <div className="flex items-center">
            <div className="flex items-center ms-3 gap-4">
              <Dropdown
                label={
                  companies?.find((c) => c.id === companyID)?.name ??
                  "no company"
                }
                color="gray"
              >
                {companies?.map((c) => (
                  <Dropdown.Item
                    as={"button"}
                    key={c.id}
                    onClick={() => {
                      setCompanyID(c.id!);
                      window.location.href = "/"
                    }}
                  >
                    {c.name}
                  </Dropdown.Item>
                ))}
              </Dropdown>
              <Avatar
                size="xs"
                img={profile?.profile_picture?.url}
                rounded
                stacked
                placeholderInitials={initial(profile?.full_name)}
              />
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Topnav;
