import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Badge,
  Button,
  Label,
  Modal,
  Pagination,
  Table,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import { PaginationResponse } from "../objects/pagination";
import {
  createAiAgent,
  deleteAiAgent,
  getAiAgents,
} from "../services/api/aiAgentApi";
import { AgentModel } from "../models/agent";
import Select, { InputActionMeta } from "react-select";
import { ActiveCompanyContext } from "../contexts/CompanyContext";
import toast from "react-hot-toast";
import { getPagination } from "../utils/helper";
import { useNavigate } from "react-router-dom";

interface AiAgentPageProps {}

const AiAgentPage: FC<AiAgentPageProps> = ({}) => {
  const { activeCompany, setActiveCompany } = useContext(ActiveCompanyContext);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [selectedAgent, setSelectedAgent] = useState<AgentModel>();
  const [agents, setAgents] = useState<AgentModel[]>([]);
  const navigate = useNavigate();


  useEffect(() => {
    setMounted(true);
  }, []);
  useEffect(() => {
    if (mounted) {
      getAllAgents();
    }
  }, [mounted]);

  useEffect(() => {}, [activeCompany]);
  const getAllAgents = () => {
    getAiAgents({ page, size, search }).then((resp: any) => {
      setAgents(resp.data.items);
      setPagination(getPagination(resp.data));
    });
  };
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Agent</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setSelectedAgent({
                name: "",
                api_key: activeCompany?.setting?.gemini_api_key,
                active: true,
                system_instruction: "",
                model: "gemini-2.0-flash-exp",
              });
              setShowModal(true);
            }}
          >
            + Create new agent
          </Button>
        </div>
        <Table hoverable>
          <Table.Head>
            <Table.HeadCell>Agent Name</Table.HeadCell>
            <Table.HeadCell style={{ width: 500 }}>Instructon</Table.HeadCell>
            <Table.HeadCell>Status</Table.HeadCell>
            <Table.HeadCell></Table.HeadCell>
          </Table.Head>

          <Table.Body className="divide-y">
            {agents.length === 0 && (
              <Table.Row>
                <Table.Cell colSpan={5} className="text-center">
                  No agents found.
                </Table.Cell>
              </Table.Row>
            )}
            {agents.map((agent) => (
              <Table.Row
                key={agent.id}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                <Table.Cell
                  className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold"
                  onClick={() => navigate(`/agent/${agent.id}`)}
                >
                  {agent.name}
                </Table.Cell>
                <Table.Cell>
                  <div className="max-h-[4.5rem] overflow-hidden text-ellipsis whitespace-normal">
                    {agent.system_instruction}
                  </div>
                </Table.Cell>
                <Table.Cell>
                  {agent.active && (
                    <div className="w-fit">
                      <Badge>Active</Badge>
                    </div>
                  )}
                </Table.Cell>

                <Table.Cell>
                  <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                    onClick={() => {
                      navigate(`/ai-agent/${agent.id}`);
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
                          `Are you sure you want to delete agent ${agent.name}?`
                        )
                      ) {
                        deleteAiAgent(agent?.id!).then(() => {
                          getAllAgents();
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
      <Modal
        show={showModal}
        onClose={() => {
          setShowModal(false);
        }}
      >
        <Modal.Header>Create new agent</Modal.Header>
        <Modal.Body>
          <form>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Name" />
              <TextInput
                id="name"
                type="text"
                value={selectedAgent?.name ?? ""}
                onChange={(e) =>
                  setSelectedAgent((prev) => ({
                    ...prev,
                    name: e.target.value,
                  }))
                }
                placeholder="Name"
                required={true}
              />
            </div>
            <div className="mb-2 block">
              <Label htmlFor="system-instruction" value="System Instructon" />
              <Textarea
                id="system-instruction"
                value={selectedAgent?.system_instruction ?? ""}
                onChange={(e) =>
                  setSelectedAgent((prev) => ({
                    ...prev,
                    system_instruction: e.target.value,
                  }))
                }
                placeholder="System Instructon"
                rows={5}
                required={true}
              />
            </div>
           
          
           
          </form>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex justify-end gap-2 w-full">
            <Button color="gray" onClick={() => {
              setShowModal(false);
            }}>
              Cancel
            </Button>
            <Button
              onClick={() => {
                createAiAgent(selectedAgent)
                  .then(() => {
                    getAllAgents();
                    setShowModal(false);
                  })
                  .catch(toast.error).finally(() => {
                    
                  })
              }}
            >
              Create
            </Button>
          </div>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default AiAgentPage;
