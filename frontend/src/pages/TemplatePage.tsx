import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Button,
  Label,
  Modal,
  Pagination,
  Table,
  Textarea,
  TextInput,
} from "flowbite-react";
import { LoadingContext } from "../contexts/LoadingContext";
import { PaginationResponse } from "../objects/pagination";
import {
  createTemplate,
  deleteTemplate,
  getTemplates,
  updateTemplate,
} from "../services/api/templateApi";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { TemplateModel } from "../models/template";
import { getContrastColor, getPagination } from "../utils/helper";

interface TemplatePageProps {}

const TemplatePage: FC<TemplatePageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [showModal, setShowModal] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const nav = useNavigate();
  const [selectedTemplate, setSelectedTemplate] = useState<TemplateModel>();
  const [mounted, setMounted] = useState(false);
  const [templates, setTemplates] = useState<TemplateModel[]>([]);

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
      setLoading(true);
      let resp: any = await getTemplates({ page, size, search });
      setTemplates(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const save = async () => {
    setLoading(true);
    try {
      // if (selectedTemplate) {
      //   await updateTemplate(selectedTemplate!.id, {
      //     title,
      //   });
      // } else {
      let resp: any = await createTemplate({
        title,
        description,
      });
      // }
      nav(`/template/${resp.data.id}`);
      setShowModal(false);
      setTitle("");
      setDescription("");
      setSelectedTemplate(undefined);
      toast.success("Save successfully");
      getAllTemplates();
    } catch (error) {
      console.log(error);
      toast.error("Save failed");
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Template</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setShowModal(true);
            }}
          >
            + Create new template
          </Button>
        </div>
        <div className="h-[calc(100vh-240px)] overflow-y-auto">
          <Table striped>
            <Table.Head>
              <Table.HeadCell>Title</Table.HeadCell>
              <Table.HeadCell>Description</Table.HeadCell>
              <Table.HeadCell>Type</Table.HeadCell>
              <Table.HeadCell></Table.HeadCell>
            </Table.Head>
            <Table.Body className="bg-white">
              {templates.length === 0 && (
                <Table.Row>
                  <Table.Cell colSpan={3} className="text-center">
                    No template found.
                  </Table.Cell>
                </Table.Row>
              )}
              {templates.map((template) => (
                <Table.Row
                  key={template.id}
                  className="bg-white dark:border-gray-700 dark:bg-gray-800"
                >
                  <Table.Cell>
                    <span className="font-medium">{template.title}</span>
                  </Table.Cell>
                  <Table.Cell>
                    <div className="flex flex-row gap-4">
                      {(template.messages ?? []).length > 0 &&
                        (template?.messages?.[0]?.files ?? []).filter((file) =>
                          file.mime_type.includes("image")
                        ).length > 0 && (
                          <div className="w-16 h-16 aspect-square border rounded-lg">
                            <img
                              src={
                                template?.messages?.[0]?.files?.find((file) =>
                                  file.mime_type.includes("image")
                                )?.url
                              }
                              alt=""
                              className="object-cover w-full h-full rounded-lg"
                            />
                          </div>
                        )}

                      <span className="font-medium">
                        {template.description}
                      </span>
                    </div>
                  </Table.Cell>
                  <Table.Cell>
                    <span className="font-medium">{template.messages?.length && template.messages?.[0]?.type}</span>
                  </Table.Cell>

                  <Table.Cell>
                    <a
                      href="#"
                      className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                      onClick={() => {
                        nav(`/template/${template.id}`);
                      }}
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
                          deleteTemplate(template?.id!).then(() => {
                            getAllTemplates();
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
        </div>
        <Pagination
          className="mt-4"
          currentPage={page}
          totalPages={pagination?.total_pages ?? 0}
          onPageChange={(val) => {
            setPage(val);
          }}
          showIcons
        />
      </div>
      <Modal show={showModal} onClose={() => setShowModal(false)}>
        <Modal.Header>Create new template</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col space-y-4">
            <div>
              <Label htmlFor="name" value="Name" className="mb-1" />
              <TextInput
                id="name"
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Name"
                className="mb-4"
              />
            </div>

            <div>
              <Label
                htmlFor="description"
                value="Description"
                className="mb-1"
              />
              <Textarea
                rows={7}
                id="description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="Description"
                className="mb-4"
              />
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex flex-row justify-end w-full">
            <Button onClick={save}>Save</Button>
          </div>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default TemplatePage;
