import {
  Button,
  FileInput,
  Label,
  Modal,
  Table,
  Tabs,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { SiGoogleforms } from "react-icons/si";
import { FaWpforms } from "react-icons/fa6";
import toast from "react-hot-toast";
import {
  createForm,
  createFormTemplate,
  deleteForm,
  deleteFormTemplate,
  getForms,
  getFormTemplates,
} from "../services/api/formApi";
import { LoadingContext } from "../contexts/LoadingContext";
import { PaginationResponse } from "../objects/pagination";
import { FormModel, FormTemplateModel } from "../models/form";
import { getPagination } from "../utils/helper";
import { useNavigate } from "react-router-dom";
import { FileModel } from "../models/file";
import { uploadFile } from "../services/api/commonApi";
import Select, { InputActionMeta } from "react-select";
import { getProjects } from "../services/api/projectApi";
import { ProjectModel } from "../models/project";
import { ColumnModel } from "../models/column";

interface FormPageProps {}

const FormPage: FC<FormPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [showModal, setShowModal] = useState(false);
  const [showModalForm, setShowModalForm] = useState(false);
  const [activeTab, setActiveTab] = useState(0);
  const [formTemplateTitle, setFormTemplateTitle] = useState("");
  const [formTemplateDesc, setFormTemplateDesc] = useState("");
  const [formTemplates, setFormTemplates] = useState<FormTemplateModel[]>([]);
  const [mounted, setMounted] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [pageForm, setPageForm] = useState(1);
  const [sizeForm, setSizeForm] = useState(20);
  const [searchForm, setSearchForm] = useState("");
  const [paginationForm, setPaginationForm] = useState<PaginationResponse>();
  const [formTitle, setFormTitle] = useState("");
  const [formDescription, setFormDescription] = useState("");
  const [fileCover, setFileCover] = useState<FileModel>();
  const [filePicture, setFilePicture] = useState<FileModel>();
  const [formPublic, setFormPublic] = useState(false);
  const [formWebhook, setFormWebhook] = useState(false);
  const [formTemplate, setFormTemplate] = useState<FormTemplateModel>();
  const [formMethod, setFormMethod] = useState("");
  const [formUrl, setFormUrl] = useState("");
  const [formHeaders, setFormHeaders] = useState("");
  const [projects, setProjects] = useState<ProjectModel[]>([]);
  const [formProject, setFormProject] = useState<ProjectModel>();
  const [formColumn, setFormColumn] = useState<ColumnModel>();
  const [forms, setForms] = useState<FormModel[]>([]);
  const nav = useNavigate();

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllTemplates();
      getAllForms();
      getAllProjects("");
    }
  }, [mounted, page, size, search]);

  const getAllTemplates = async () => {
    try {
      const resp: any = await getFormTemplates({ page, size, search });
      setFormTemplates(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const getAllForms = async () => {
    try {
      const resp: any = await getForms({
        page: pageForm,
        size: sizeForm,
        search: searchForm,
      });
      setForms(resp.data.items);
      setPaginationForm(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const getAllProjects = (s: string) => {
    getProjects({ page: 1, size: 10, search: s })
      .then((e: any) => setProjects(e.data.items))
      .catch(toast.error);
  };

  const createFormTemplateProcess = async () => {
    try {
      if (!formTemplateTitle) {
        toast.error("Please fill all the fields");
        return;
      }
      setLoading(true);
      const response: any = await createFormTemplate({
        title: formTemplateTitle,
        description: formTemplateDesc,
      });
      toast.success("Form template created successfully");
      nav(`/form-template/${response.id}`);
      setShowModal(false);
      setFormTemplateTitle("");
      setFormTemplateDesc("");
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const createFormProcess = async () => {
    try {
      if (!formTitle) {
        toast.error("Please fill all the fields");
        return;
      }
      if (!formTemplate) {
        toast.error("Please fill all the fields");
        return;
      }
      let data = {
        title: formTitle,
        description: formDescription,
        cover: fileCover,
        form_template_id: formTemplate?.id,
        submit_url: formUrl,
        method: formMethod,
        headers: formHeaders,
        is_public: formPublic,
        project_id: formProject?.id,
        column_id: formColumn?.id,
      };

      setLoading(true);
      const response: any = await createForm(data);
      toast.success("Form created successfully");
      nav(`/form/${response.id}`);
      setShowModalForm(false);
      setFormTitle("");
      setFormDescription("");
      setFileCover(undefined);
      setFormUrl("");
      setFormHeaders("");
      setFormPublic(false);
      setFormProject(undefined);
      setFormColumn(undefined);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-4">
        <Tabs
          aria-label="Default tabs"
          variant="default"
          onActiveTabChange={(tab) => {
            setActiveTab(tab);
            // console.log(tab);
          }}
        >
          <Tabs.Item
            title="Form Template"
            active={activeTab == 0}
            icon={SiGoogleforms}
          >
            <div className="p-4">
              <div className="flex justify-between items-center mb-4">
                <h1 className="text-3xl font-bold ">Form Template</h1>
                <Button
                  gradientDuoTone="purpleToBlue"
                  pill
                  onClick={() => {
                    setShowModal(true);
                  }}
                >
                  + Create Form Template
                </Button>
              </div>
              <Table hoverable={true}>
                <Table.Head>
                  <Table.HeadCell>Title</Table.HeadCell>
                  <Table.HeadCell>Created By</Table.HeadCell>
                  <Table.HeadCell>Action</Table.HeadCell>
                </Table.Head>
                <Table.Body className="bg-white">
                  {formTemplates.map((template) => (
                    <Table.Row
                      key={template.id}
                      className="bg-white dark:border-gray-700 dark:bg-gray-800"
                    >
                      <Table.Cell>
                        <span className="font-medium">{template.title}</span>
                      </Table.Cell>
                      <Table.Cell>
                        {template.created_by_member?.user?.full_name}
                      </Table.Cell>
                      <Table.Cell>
                        <a
                          href="#"
                          className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                          onClick={() => nav(`/form-template/${template.id}`)}
                        >
                          View
                        </a>
                        <a
                          href="#"
                          className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                          onClick={(e) => {
                            e.preventDefault();
                            if (
                              window.confirm(
                                `Are you sure you want to delete  ${template.title}?`
                              )
                            ) {
                              setLoading(true);
                              deleteFormTemplate(template.id!).then(() => {
                                getAllTemplates();
                              }).catch((e) => {
                                toast.error(`${e}`);
                              }).then(() => {
                                setLoading(false);
                              })

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
            </div>
          </Tabs.Item>
          <Tabs.Item title="Form" active={activeTab == 1} icon={FaWpforms}>
            <div className="p-4">
              <div className="flex justify-between items-center mb-4">
                <h1 className="text-3xl font-bold ">Form </h1>
                <Button
                  gradientDuoTone="purpleToBlue"
                  pill
                  onClick={() => {
                    setShowModalForm(true);
                  }}
                >
                  + Create Form
                </Button>
              </div>
              <Table hoverable={true}>
                <Table.Head>
                  <Table.HeadCell>Title</Table.HeadCell>
                  <Table.HeadCell>Created By</Table.HeadCell>
                  <Table.HeadCell>Template</Table.HeadCell>
                  <Table.HeadCell>Project</Table.HeadCell>
                  <Table.HeadCell>Action</Table.HeadCell>
                </Table.Head>
                <Table.Body className="bg-white">
                  {forms.map((form) => (
                    <Table.Row
                      key={form.id}
                      className="bg-white dark:border-gray-700 dark:bg-gray-800"
                    >
                      <Table.Cell>
                        <span className="font-medium">{form.title}</span>
                      </Table.Cell>
                      <Table.Cell>
                        {form.created_by_member?.user?.full_name}
                      </Table.Cell>
                      <Table.Cell>{form.form_template?.title}</Table.Cell>
                      <Table.Cell>{form.project?.name}</Table.Cell>
                      <Table.Cell>
                        <a
                          href="#"
                          className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                          onClick={() => nav(`/form/${form.id}`)}
                        >
                          View
                        </a>
                        <a
                          href="#"
                          className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                          onClick={(e) => {
                            e.preventDefault();
                            if (
                              window.confirm(
                                `Are you sure you want to delete  ${form.title}?`
                              )
                            ) {
                              setLoading(true);
                              deleteForm(form.id!).then(() => {
                                getAllForms();
                              }).catch((e) => {
                                toast.error(`${e}`);
                              }).then(() => {
                                setLoading(false);
                              })
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
            </div>
          </Tabs.Item>
        </Tabs>
      </div>
      <Modal
        size="xl"
        show={showModalForm}
        onClose={() => setShowModalForm(false)}
      >
        <Modal.Header>Create Form</Modal.Header>
        <Modal.Body className="flex flex-col space-y-8">
          <div>
            <div className=" block">
              <Label htmlFor="name">Cover</Label>
            </div>
            <FileInput
              id="name"
              placeholder="cover"
              accept="image/*"
              onChange={(val) => {
                const files = val?.target.files;
                if (files) {
                  try {
                    uploadFile(files[0], {}, (val) => console.log).then(
                      (v: any) => {
                        setFileCover(v.data);
                      }
                    );
                  } catch (error) {
                    console.log(error);
                  }
                }
              }}
            />
            {fileCover && (
              <img
                src={fileCover.url}
                className=" aspect-video w-full object-cover rounded-lg mt-4 bg-gray-50"
              />
            )}
          </div>
          <div>
            <div className=" block">
              <Label htmlFor="name">Template</Label>
            </div>
            <Select
              value={
                formTemplate
                  ? { label: formTemplate?.title, value: formTemplate?.id }
                  : null
              }
              options={formTemplates.map((e) => ({
                label: e.title,
                value: e.id,
              }))}
              onChange={(val) => {
                setFormTemplate(formTemplates.find((e) => e.id == val?.value));
              }}
            />
          </div>
          <div>
            <div className=" block">
              <Label htmlFor="name">Title</Label>
            </div>
            <TextInput
              id="name"
              type="text"
              placeholder="Title"
              required={true}
              value={formTitle}
              onChange={(e) => {
                setFormTitle(e.target.value);
              }}
            />
          </div>
          <div>
            <div className=" block">
              <Label htmlFor="name">Description</Label>
            </div>
            <Textarea
              id="description"
              placeholder="Description"
              required={true}
              value={formDescription}
              onChange={(e) => {
                setFormDescription(e.target.value);
              }}
            />
          </div>
          <div>
            <div className=" block">
              <Label htmlFor="name">Project</Label>
            </div>
            <Select
              value={
                formProject
                  ? { label: formProject?.name, value: formProject?.id }
                  : null
              }
              options={[
                { label: "Select Project", value: "" },
                ...projects.map((e) => ({
                  label: e.name,
                  value: e.id,
                })),
              ]}
              onChange={(val) => {
                setFormProject(projects.find((e) => e.id == val?.value));
              }}
              onInputChange={(val) => {
                getAllProjects(val);
              }}
            />
          </div>
          {formProject && (
            <div>
              <div className=" block">
                <Label htmlFor="name">Column</Label>
              </div>
              <Select
                value={
                  formColumn
                    ? { label: formColumn?.name, value: formColumn?.id }
                    : null
                }
                options={[
                  { label: "Select Column", value: "" },
                  ...formProject!.columns!.map((e) => ({
                    label: e.name!,
                    value: e.id!,
                  })),
                ]}
                onChange={(val) => {
                  setFormColumn(
                    formProject!.columns!.find((e) => e.id == val?.value)
                  );
                }}
              />
            </div>
          )}
          <div>
            <div className=" grid grid-cols-2">
              <ToggleSwitch
                checked={formPublic}
                onChange={(val) => setFormPublic(val)}
                label={"Is Public"}
              />
              <ToggleSwitch
                checked={formWebhook}
                onChange={(val) => setFormWebhook(val)}
                label={"Web Hook"}
              />
            </div>
          </div>
          {formWebhook && (
            <div>
              <div className=" block">
                <Label htmlFor="name">Webhook Submit URL</Label>
              </div>
              <TextInput
                name="url"
                id="url"
                placeholder="Webhook Submit URL"
                value={formUrl}
                onChange={(e) => {
                  setFormUrl(e.target.value);
                }}
              />
            </div>
          )}
          {formWebhook && (
            <div>
              <div className=" block">
                <Label htmlFor="name">Webhook Method</Label>
              </div>
              <Select
                options={[
                  { label: "POST", value: "POST" },
                  { label: "GET", value: "GET" },
                  { label: "PUT", value: "PUT" },
                  { label: "DELETE", value: "DELETE" },
                ]}
                value={{ value: formMethod, label: formMethod }}
                onChange={(val) => setFormMethod(val?.value || "POST")}
              />
            </div>
          )}
          {formWebhook && (
            <div>
              <div className=" block">
                <Label htmlFor="name">Webhook Headers</Label>
              </div>
              <Textarea
                id="header"
                placeholder={`{
        "token":"abcd"
}`}
                value={formHeaders}
                onChange={(e) => {
                  setFormHeaders(e.target.value);
                }}
              />
            </div>
          )}
          <div className="h-[100px] w-full"></div>
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button type="submit" onClick={createFormProcess}>
            Create
          </Button>
        </Modal.Footer>
      </Modal>
      <Modal
        show={showModal}
        onClose={() => {
          setShowModal(false);
        }}
      >
        <Modal.Header>Create Form Template</Modal.Header>
        <Modal.Body className="space-y-6">
          <div>
            <div className="mb-2 block">
              <Label htmlFor="name">Title</Label>
            </div>
            <TextInput
              id="name"
              type="text"
              placeholder="Title"
              required={true}
              value={formTemplateTitle}
              onChange={(e) => {
                setFormTemplateTitle(e.target.value);
              }}
            />
          </div>
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button type="submit" onClick={createFormTemplateProcess}>
            Create
          </Button>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default FormPage;
