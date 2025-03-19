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
  createTaskAttribute,
  deleteTaskAttribute,
  getTaskAttributes,
} from "../services/api/taskAttributeApi";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { TaskAttributeModel } from "../models/task_attribute";
import { getPagination } from "../utils/helper";

interface TaskAttributePageProps {}

const TaskAttributePage: FC<TaskAttributePageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [showModal, setShowModal] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const nav = useNavigate();
  const [mounted, setMounted] = useState(false);
  const [taskAttributes, setTaskAttributes] = useState<TaskAttributeModel[]>(
    []
  );

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllTaskAttributes();
    }
  }, [mounted, page, size, search]);

  const getAllTaskAttributes = async () => {
    try {
      setLoading(true);
      let resp: any = await getTaskAttributes({ page, size, search });
      setTaskAttributes(resp.data.items);
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
      const res: any = await createTaskAttribute({
        title,
        description,
        data: "[]",
      });
      setShowModal(false);
      setTitle("");
      setDescription("");
      toast.success("Save successfully");
      nav(`/task-attribute/${res.id}`);
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
          <h1 className="text-3xl font-bold ">Task Attribute</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setShowModal(true);
            }}
          >
            + Create new attribute
          </Button>
        </div>
        <Table hoverable={true}>
          <Table.Head>
            <Table.HeadCell>Title</Table.HeadCell>
            <Table.HeadCell>Description</Table.HeadCell>
            <Table.HeadCell>Action</Table.HeadCell>
          </Table.Head>
          <Table.Body className="bg-white">
            {taskAttributes.map((attribute) => (
              <Table.Row
                key={attribute.id}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                <Table.Cell>
                  <span className="font-medium">{attribute.title}</span>
                </Table.Cell>
                <Table.Cell>{attribute.description}</Table.Cell>
                <Table.Cell>
                  <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                    onClick={() => nav(`/task-attribute/${attribute.id}`)}
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
                          `Are you sure you want to delete  ${attribute.title}?`
                        )
                      ) {
                        deleteTaskAttribute(attribute.id);
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
        <Modal.Header>Create new attribute</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col space-y-4">
            <div>
              <Label htmlFor="title" value="Title" className="mb-1" />
              <TextInput
                id="title"
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Title"
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
            <Button onClick={save}>Create</Button>
          </div>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default TaskAttributePage;
