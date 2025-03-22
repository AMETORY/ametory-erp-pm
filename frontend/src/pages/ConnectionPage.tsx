import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Badge,
  Button,
  Drawer,
  Label,
  Modal,
  Pagination,
  Table,
  Textarea,
  TextInput,
} from "flowbite-react";
import Select, { InputActionMeta } from "react-select";
import {
  createConnection,
  deleteConnection,
  getConnection,
  getConnections,
  updateConnection,
} from "../services/api/connectionApi";
import { useNavigate } from "react-router-dom";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import { ConnectionModel } from "../models/connection";
import { PaginationResponse } from "../objects/pagination";
import ConnectionDrawer from "../components/ConnectionDrawer";
interface ConnectionPageProps {}

const ConnectionPage: FC<ConnectionPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [showModal, setShowModal] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [name, setName] = useState("");
  const [connections, setConnections] = useState<ConnectionModel[]>([]);
  const [mounted, setMounted] = useState(false);
  const [showDetailOpen, setShowDetailOpen] = useState(false);
  const [description, setDescription] = useState("");
  const [sessionName, setSessionName] = useState("");
  const [activeConnection, setActiveConnection] = useState<ConnectionModel>();
  const [selectedConnection, setSelectedConnection] = useState<{
    label: string;
    value: string;
  }>({ label: "WHATSAPP", value: "whatsapp" });
  const nav = useNavigate();
  var connectionType = [{ label: "WHATSAPP", value: "whatsapp" }];

  useEffect(() => {
    setMounted(true);
  }, []);

  const createNewConnection = async () => {
    try {
      if (!name || !description) {
        throw new Error("Name and description are required");
      }
      if (selectedConnection.value == "whatsapp") {
        if (!sessionName) {
          throw new Error("Phone Number is required");
        }
      }
      setLoading(true);
      let resp: any = await createConnection({
        name,
        description,
        type: selectedConnection.value,
        session_name: sessionName,
      });
      getConnection(resp.id).then((resp: any) => {
        setActiveConnection(resp.data);
        setShowDetailOpen(true);
        setShowModal(false);
        getAllConnections();
      });
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const getAllConnections = async () => {
    try {
      setLoading(true);
      let resp: any = await getConnections({ page, size });
      setConnections(resp.data);
      setPagination({
        page,
        size,
        total_pages: resp.pagination.total_pages,
        total: resp.pagination.total_rows,
      });
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (mounted) {
      getAllConnections();
    }
  }, [mounted]);
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Connection</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setShowModal(true);
            }}
          >
            + Create new Connection
          </Button>
        </div>
        <Table>
          <Table.Head>
            <Table.HeadCell>Name</Table.HeadCell>
            <Table.HeadCell>Description</Table.HeadCell>
            <Table.HeadCell>Type</Table.HeadCell>
            <Table.HeadCell>Status</Table.HeadCell>
            <Table.HeadCell></Table.HeadCell>
          </Table.Head>

          <Table.Body className="divide-y">
            {connections.length === 0 && (
              <Table.Row>
                <Table.Cell colSpan={5} className="text-center">
                  No connections found.
                </Table.Cell>
              </Table.Row>
            )}
            {connections.map((connection) => (
              <Table.Row
                key={connection.id}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                <Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold">
                  {connection.name}
                </Table.Cell>
                <Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold">
                  {connection.description}
                </Table.Cell>

                <Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold">
                  {connection.type}
                </Table.Cell>
                <Table.Cell className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold">
                  <div className="flex w-fit">
                    {connection?.status == "PENDING" && (
                      <Badge color="warning">{connection?.status}</Badge>
                    )}
                    {connection?.status == "ACTIVE" && (
                      <Badge color="success">{connection?.status}</Badge>
                    )}
                  </div>
                </Table.Cell>

                <Table.Cell>
                  <a
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500 cursor-pointer"
                    onClick={() => {
                      // nav(`/connection/${connection.id}`);

                      getConnection(connection.id!).then((res: any) => {
                        setShowDetailOpen(true);
                        setActiveConnection(res.data);
                      });
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
                          `Are you sure you want to delete connection ${connection.name}?`
                        )
                      ) {
                        deleteConnection(connection?.id!).then(() => {
                          getAllConnections();
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
        <Modal.Header>Create Connection</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col space-y-4">
            <div className="mb-2 block">
              <Label htmlFor="type" value="Type" />
              <Select
                options={connectionType}
                value={selectedConnection}
                onChange={(val) => setSelectedConnection(val as any)}
              />
            </div>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Name" />
              <TextInput
                id="name"
                placeholder="Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </div>
            {selectedConnection.value === "whatsapp" && (
              <div className="mb-2 block">
                <Label htmlFor="session_name" value="Phone Number" />
                <TextInput
                  id="session_name"
                  placeholder="Phone Number"
                  value={sessionName}
                  onChange={(e) => setSessionName(e.target.value)}
                />
              </div>
            )}

            <div className="mb-2 block">
              <Label htmlFor="description" value="Description" />
              <Textarea
                id="description"
                placeholder="Description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
              />
            </div>
            <div className="h-40"></div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex justify-end w-full">
            <Button onClick={createNewConnection}>Create</Button>
          </div>
        </Modal.Footer>
      </Modal>
      {activeConnection && (
        <Drawer
          style={{ width: "600px" }}
          position="right"
          open={showDetailOpen}
          onClose={() => {
            setShowDetailOpen(false);
            setActiveConnection(undefined);
          }}
        >
          <ConnectionDrawer
            connection={activeConnection}
            onUpdate={(connection) => {
              setActiveConnection(connection);
              // console.log(connection);
              // updateConnection(connection.id!, connection).then(() => {
              //   getAllConnections();
              // })
            }}
            onFinish={() => {
              // setShowDetailOpen(false);
              getAllConnections();
              setActiveConnection(undefined);
            }}
            onSave={() => {
              updateConnection(activeConnection.id!, activeConnection).then(
                () => {
                  getAllConnections();
                }
              );
            }}
          />
        </Drawer>
      )}
    </AdminLayout>
  );
};
export default ConnectionPage;
