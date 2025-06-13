import { useContext, useEffect, useRef, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Button,
  FileInput,
  HR,
  Label,
  Modal,
  ModalBody,
  ModalHeader,
  Table,
  Tabs,
  TabsRef,
  Textarea,
  TextInput,
} from "flowbite-react";
import { BsInfoCircle, BsPlugin } from "react-icons/bs";
import {
  addRapidAPIPlugins,
  createCompany,
  deleteCompanyRapidAPIPlugin,
  getCompanyRapidAPIPlugins,
  getRapidAPIPlugins,
  getSetting,
  updateSetting,
  uploadFile,
} from "../services/api/commonApi";
import { CompanyModel } from "../models/company";
import toast from "react-hot-toast";
import { LoadingContext } from "../contexts/LoadingContext";
import { FileModel } from "../models/file";
import {
  CompanyRapidApiPluginModel,
  RapidApiPluginModel,
} from "../models/rapid_api";
import Select, { InputActionMeta } from "react-select";
import { LuLink, LuLink2 } from "react-icons/lu";
import { Link } from "react-router-dom";
import {
  ActiveCompanyContext,
  CompaniesContext,
} from "../contexts/CompanyContext";
import { getProfile } from "../services/api/authApi";

interface SettingPageProps {}

const SettingPage: FC<SettingPageProps> = ({}) => {
  const { companies, setCompanies } = useContext(CompaniesContext);

  const { activeCompany, setActiveCompany } = useContext(ActiveCompanyContext);
  const tabsRef = useRef<TabsRef>(null);
  const [activeTab, setActiveTab] = useState(0);
  const [company, setCompany] = useState<CompanyModel>();
  const [mounted, setMounted] = useState(false);
  const { loading, setLoading } = useContext(LoadingContext);
  const [file, setFile] = useState<FileModel>();
  const [fileNewCompany, setNewCompanyFile] = useState<FileModel>();
  const [plugins, setPlugins] = useState<RapidApiPluginModel[]>([]);
  const [modalPluginOpen, setModalPluginOpen] = useState(false);
  const [selectedPlugin, setSelectedPlugin] = useState<RapidApiPluginModel>();
  const [modalCompany, setModalCompany] = useState(false);
  const [companyPlugins, setCompanyPlugins] = useState<
    CompanyRapidApiPluginModel[]
  >([]);
  const [pluginKey, setPluginKey] = useState("");
  const [pluginHost, setPluginHost] = useState("");
  const [newCompany, setNewCompany] = useState<CompanyModel>();

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      setLoading(true);
      getAllSetting();
      getAllPlugins();
      getAllCompanyPlugins();
    }
  }, [mounted]);

  const getAllSetting = async () => {
    try {
      setLoading(true);
      const resp: any = await getSetting();
      setCompany(resp.data);
      setActiveCompany(resp.data);
    } catch (error: any) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const getAllPlugins = async () => {
    try {
      setLoading(true);
      const resp: any = await getRapidAPIPlugins();
      setPlugins(resp.data);
    } catch (error: any) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const getAllCompanyPlugins = async () => {
    try {
      setLoading(true);
      const resp: any = await getCompanyRapidAPIPlugins();
      setCompanyPlugins(resp.data);
    } catch (error: any) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const updateCompanySetting = async () => {
    try {
      setLoading(true);
      //   await updateProfile(profile!);
      await updateSetting(company!);
      toast.success("Company updated successfully");
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const renderInfo = () => (
    <div className="flex flex-col gap-4 overflow-y-auto h-[calc(100vh-160px)] p-2">
      <h1 className="text-3xl font-bold">Edit Company</h1>
      <div className="bg-white rounded-lg p-4">
        <div className="flex flex-col gap-2 space-y-4">
          {company?.logo && (
            <div className="flex justify-center py-4 items-center">
              <img
                className="w-64 h-64 aspect-square object-cover rounded-full"
                src={company?.logo}
                alt="profile"
              />
            </div>
          )}

          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Logo</label>
            <FileInput
              accept="image/*"
              id="file-upload"
              onChange={(el) => {
                if (el.target.files) {
                  let f = el.target.files[0];
                  if (!f) return;
                  uploadFile(f, {}, (val) => {
                    console.log(val);
                  }).then((v: any) => {
                    setFile(v.data);
                    setCompany({
                      ...company!,
                      logo: v.data.url,
                    });
                  });
                }
              }}
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Name</label>
            <TextInput
              type="text"
              value={company?.name}
              name="company_name"
              onChange={(e) =>
                setCompany({ ...company!, name: e.target.value })
              }
              placeholder="Enter company name"
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Address</label>
            <Textarea
              value={company?.address}
              name="address"
              onChange={(e) =>
                setCompany({ ...company!, address: e.target.value })
              }
              placeholder="Enter company address"
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Email</label>
            <TextInput
              type="email"
              value={company?.email}
              name="email"
              onChange={(e) =>
                setCompany({ ...company!, email: e.target.value })
              }
              placeholder="Enter company email"
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Phone</label>
            <TextInput
              type="text"
              name="phone"
              value={company?.phone}
              onChange={(e) =>
                setCompany({ ...company!, phone: e.target.value })
              }
              placeholder="Enter company phone"
            />
          </div>
          <div className="flex justify-end items-center gap-2">
            <Button
              className="mt-8  bg-green-600"
              onClick={async () => {
                setModalCompany(true);
                setNewCompany({});
              }}
            >
              Create New Company
            </Button>
            <Button
              type="submit"
              className="mt-8 w-32"
              onClick={updateCompanySetting}
            >
              Save
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
  const renderPlugin = () => (
    <div className="flex flex-col gap-4 overflow-y-auto h-[calc(100vh-160px)] p-2">
      <div className="flex justify-between">
        <h1 className="text-3xl font-bold">Rapid API Plugin</h1>
        <Button
          gradientDuoTone="purpleToBlue"
          pill
          onClick={() => {
            setModalPluginOpen(true);
          }}
        >
          + Add Plugin
        </Button>
      </div>
      <Table>
        <Table.Head>
          <Table.HeadCell>Plugin Name</Table.HeadCell>
          <Table.HeadCell>Plugin URL</Table.HeadCell>
          <Table.HeadCell>Key</Table.HeadCell>
          <Table.HeadCell>Host</Table.HeadCell>
          <Table.HeadCell>Actions</Table.HeadCell>
        </Table.Head>
        <Table.Body className="divide-y">
          {companyPlugins.length === 0 && (
            <Table.Row>
              <Table.Cell colSpan={5} className="text-center">
                No plugins found.
              </Table.Cell>
            </Table.Row>
          )}
          {companyPlugins.map((plugin, i) => (
            <Table.Row
              key={i}
              className="bg-white dark:border-gray-700 dark:bg-gray-800"
            >
              <Table.Cell>{plugin.rapid_api_plugin?.name}</Table.Cell>
              <Table.Cell>
                <Link
                  className=" hover:text-blue-400 hover:underline"
                  target="_blank"
                  to={plugin.rapid_api_plugin?.url}
                >
                  {plugin.rapid_api_plugin?.url}
                </Link>
              </Table.Cell>
              <Table.Cell>{plugin.rapid_api_key}</Table.Cell>
              <Table.Cell>{plugin.rapid_api_host}</Table.Cell>
              <Table.Cell>
                {/* <a
                                      href="#"
                                      className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                                    >
                                      Edit
                                    </a> */}
                <a
                  href="#"
                  className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                  onClick={(e) => {
                    e.preventDefault();
                    if (
                      window.confirm(
                        `Are you sure you want to delete  ${plugin?.rapid_api_plugin?.name}?`
                      )
                    ) {
                      deleteCompanyRapidAPIPlugin(
                        plugin?.rapid_api_plugin_id
                      ).then(() => {
                        getAllCompanyPlugins();
                      });
                    }
                  }}
                >
                  Delete
                </a>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
      <HR />
      <h1 className="text-3xl font-bold">Gemini</h1>
      <div className="flex flex-row  items-center">
        <label className="text-sm font-medium w-1/4">API Key</label>

        <TextInput
          type="text"
          placeholder="Gemini API Key"
          className="w-3/4"
          value={company?.setting?.gemini_api_key ?? ""}
          onChange={(el) => {
            setCompany({
              ...company!,
              setting: {
                ...company!.setting!,
                gemini_api_key: el.target.value,
                company_id: company?.id!,
              },
            });
          }}
        />
      </div>

      <div>
        <Button type="submit" className="mt-6" onClick={updateCompanySetting}>
          Submit
        </Button>
      </div>

      <HR />
    </div>
  );

  const addPlugin = async () => {
    try {
      setLoading(true);
      await addRapidAPIPlugins({
        id: selectedPlugin?.id,
        key: pluginKey,
        host: pluginHost,
      });
      getAllCompanyPlugins();
      toast.success("Plugin added successfully");
      setModalPluginOpen(false);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="w-full h-full flex flex-col gap-4 px-8">
        <Tabs
          aria-label="Default tabs"
          variant="default"
          ref={tabsRef}
          onActiveTabChange={(tab) => {
            setActiveTab(tab);
            // console.log(tab);
          }}
          className="mt-4"
        >
          <Tabs.Item
            active={activeTab === 0}
            title="Basic Info"
            icon={BsInfoCircle}
          >
            {renderInfo()}
          </Tabs.Item>
          <Tabs.Item
            active={activeTab === 1}
            title="Plugin"
            icon={BsPlugin}
            className=""
          >
            {renderPlugin()}
          </Tabs.Item>
        </Tabs>
      </div>
      <Modal show={modalPluginOpen} onClose={() => setModalPluginOpen(false)}>
        <Modal.Header>Add Plugin</Modal.Header>
        <Modal.Body>
          <div className="space-y-4 flex flex-col">
            <div className="flex flex-col gap-1 ">
              <Label htmlFor="plugin-name">Name</Label>
              <Select
                id="plugin-name"
                value={
                  selectedPlugin
                    ? {
                        label: selectedPlugin.name,
                        value: selectedPlugin.id,
                      }
                    : null
                }
                onChange={(option) => {
                  setSelectedPlugin(plugins.find((e) => e.id == option?.value));
                }}
                options={plugins.map((plugin) => ({
                  label: plugin.name,
                  value: plugin.id,
                }))}
                placeholder="Select a plugin"
              />
            </div>
            {selectedPlugin?.url && (
              <div className="flex flex-col gap-1 ">
                <Label htmlFor="plugin-name">URL</Label>
                <div
                  className="flex gap-2 items-center hover:underline cursor-pointer hover:text-blue-600"
                  onClick={() => window.open(selectedPlugin?.url)}
                >
                  {selectedPlugin?.url}{" "}
                  <LuLink2 className="cursor-pointer text-gray-400 hover:text-gray-600" />
                </div>
              </div>
            )}
            <div className="flex flex-col gap-1">
              <Label htmlFor="plugin-key">Key</Label>
              <TextInput
                id="plugin-key"
                type="text"
                value={pluginKey}
                onChange={(e) => setPluginKey(e.target.value)}
                placeholder="Key"
              />
            </div>

            <div className="flex flex-col gap-1">
              <Label htmlFor="plugin-host">Host</Label>
              <TextInput
                id="plugin-host"
                type="text"
                value={pluginHost}
                onChange={(e) => setPluginHost(e.target.value)}
                placeholder="Host"
              />
            </div>
            <div className="min-h-[160px]"></div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex gap-2 justify-end">
            <Button onClick={addPlugin}>Save</Button>
          </div>
        </Modal.Footer>
      </Modal>
      <Modal show={modalCompany} onClose={() => setModalCompany(false)}>
        <ModalHeader>Create Company</ModalHeader>
        <ModalBody className="space-y-4 flex flex-col">
          {newCompany?.logo && (
            <div className="flex justify-center py-4 items-center">
              <img
                className="w-64 h-64 aspect-square object-cover rounded-full"
                src={newCompany?.logo}
                alt="profile"
              />
            </div>
          )}

          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Logo</label>
            <FileInput
              accept="image/*"
              id="file-upload"
              onChange={(el) => {
                if (el.target.files) {
                  let f = el.target.files[0];
                  if (!f) return;
                  uploadFile(f, {}, (val) => {
                    console.log(val);
                  }).then((v: any) => {
                    setNewCompanyFile(v.data);
                    setNewCompany({
                      ...newCompany!,
                      logo: v.data.url,
                    });
                  });
                }
              }}
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Name</label>
            <TextInput
              type="text"
              value={newCompany?.name}
              name="company_name"
              onChange={(e) =>
                setNewCompany({ ...newCompany!, name: e.target.value })
              }
              placeholder="Enter company name"
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Address</label>
            <Textarea
              value={newCompany?.address}
              name="address"
              onChange={(e) =>
                setNewCompany({ ...newCompany!, address: e.target.value })
              }
              placeholder="Enter company address"
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Email</label>
            <TextInput
              type="email"
              value={newCompany?.email}
              name="email"
              onChange={(e) =>
                setNewCompany({ ...newCompany!, email: e.target.value })
              }
              placeholder="Enter company email"
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Phone</label>
            <TextInput
              type="text"
              name="phone"
              value={newCompany?.phone}
              onChange={(e) =>
                setNewCompany({ ...newCompany!, phone: e.target.value })
              }
              placeholder="Enter company phone"
            />
          </div>
        </ModalBody>
        <Modal.Footer>
          <div className="flex gap-2 justify-end">
            <Button
              onClick={async () => {
                try {
                  setLoading(true);
                  // console.log(newCompany)
                  const response = await createCompany(newCompany);
                  setModalCompany(false);
                  toast.success("Company created successfully");
                  getProfile().then((res: any) => {
                    setCompanies(res.user.companies);
                  });
                } catch (error) {
                  toast.error(`${error}`);
                } finally {
                  setLoading(false);
                }
              }}
            >
              Save
            </Button>
          </div>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default SettingPage;
