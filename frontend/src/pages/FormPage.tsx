import {
  Button,
  Label,
  Modal,
  Table,
  Tabs,
  Textarea,
  TextInput,
} from "flowbite-react";
import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { SiGoogleforms } from "react-icons/si";
import { FaWpforms } from "react-icons/fa6";
import toast from "react-hot-toast";
import { createFormTemplate, getFormTemplates } from "../services/api/formApi";
import { LoadingContext } from "../contexts/LoadingContext";
import { PaginationResponse } from "../objects/pagination";
import { FormTemplateModel } from "../models/form";
import { getPagination } from "../utils/helper";
import { useNavigate } from "react-router-dom";

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
  const [size, setsize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const nav = useNavigate();

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllTemplates();
    }
  }, [mounted, page, size, search]);

  const getAllTemplates = async () => {
    try {
      const resp: any = await getFormTemplates({ page, size, search });
      setFormTemplates(resp.data.items);
      setPagination(getPagination(resp.data));
      toast.success("Templates fetched successfully");
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
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
            </div>
          </Tabs.Item>
        </Tabs>
      </div>
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
