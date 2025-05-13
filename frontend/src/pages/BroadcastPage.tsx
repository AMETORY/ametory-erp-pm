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
import { Mention, MentionsInput } from "react-mentions";
import { PaginationResponse } from "../objects/pagination";
import {
  createBroadcast,
  deleteBroadcast,
  getBroadcasts,
} from "../services/api/broadcastApi";
import { SearchContext } from "../contexts/SearchContext";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";
import { BroadcastModel } from "../models/broadcast";
import { parseMentions } from "../utils/helper-ui";
import { TemplateModel } from "../models/template";
import { getTemplates } from "../services/api/templateApi";
import Select from "react-select";

interface BroadcastPageProps {}
const neverMatchingRegex = /($a)/;
const BroadcastPage: FC<BroadcastPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [mounted, setMounted] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const [emojis, setEmojis] = useState([]);
  const [message, setMessage] = useState("");
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const { search, setSearch } = useContext(SearchContext);
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [description, setDescription] = useState("");
  const [broadcasts, setBroadcasts] = useState<BroadcastModel[]>([]);
  const [useTemplate, setUseTemplate] = useState(false);
  const [templates, setTemplates] = useState<TemplateModel[]>([]);
  const [selectedTemplate, setSelectedTemplate] = useState<TemplateModel>();
  const nav = useNavigate();
  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllBroadcast();
      getTemplates({ page: 1, size: 100 }).then((res: any) => {
        setTemplates(res.data.items);
      });
    }
  }, [mounted, page, size, search]);

  const getAllBroadcast = () => {
    getBroadcasts({ page, size, search }).then((res: any) => {
      setBroadcasts(res.data);
      setPagination(res.pagination);
    });
  };
  useEffect(() => {
    fetch(process.env.REACT_APP_BASE_URL + "/assets/static/emojis.json")
      .then((response) => {
        return response.json();
      })
      .then((jsonData) => {
        setEmojis(jsonData.emojis);
      });
  }, []);
  const queryEmojis = (query: any, callback: (emojis: any) => void) => {
    if (query.length === 0) return;

    const matches = emojis
      .filter((emoji: any) => {
        return emoji.name.indexOf(query.toLowerCase()) > -1;
      })
      .slice(0, 10);
    return matches.map(({ emoji }) => ({ id: emoji }));
  };

  const send = async () => {
    try {
      if (description.trim().length === 0) {
        toast.error("description is required");
        return;
      }
      if (message.trim().length === 0 && selectedTemplate === undefined) {
        toast.error("Message is required");
        return;
      }
      setLoading(true);
      let resp: any = await createBroadcast({
        message,
        description,
        template_id: selectedTemplate?.id,
      });
      nav(`/broadcast/${resp.data.id}`);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Broadcast</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setShowModal(true);
            }}
          >
            + Create new broadcast
          </Button>
        </div>
        <div className="h-[calc(100vh-240px)] overflow-y-auto">
        <Table>
          <Table.Head>
            <Table.HeadCell>ID</Table.HeadCell>
            <Table.HeadCell>Description</Table.HeadCell>
            <Table.HeadCell>Message</Table.HeadCell>
            <Table.HeadCell>Actions</Table.HeadCell>
          </Table.Head>
          <Table.Body>
            {broadcasts?.length === 0 && (
              <Table.Row>
                <Table.Cell colSpan={4} className="text-center">
                  No broadcast found
                </Table.Cell>
              </Table.Row>
            )}
            {broadcasts?.map((broadcast: any, index) => (
              <Table.Row key={broadcast.id}>
                <Table.Cell>{index + 1}</Table.Cell>
                <Table.Cell>{broadcast.description}</Table.Cell>
                <Table.Cell>
                  {parseMentions(broadcast.message, () => {})}
                </Table.Cell>
                <Table.Cell className="flex items-center justify-center gap-2">
                  <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                    onClick={() => nav(`/broadcast/${broadcast.id}`)}
                  >
                    View
                  </a>
                  <a
                    href="#"
                    className="font-medium text-red-600 hover:underline dark:text-red-500"
                    onClick={() => {
                      if (
                        window.confirm(
                          "Are you sure you want to delete this broadcast?"
                        )
                      ) {
                        setLoading(true);
                        deleteBroadcast(broadcast.id)
                          .then(() => {
                            getAllBroadcast();
                          })
                          .catch(() => {
                            toast.error("Delete failed");
                          })
                          .finally(() => setLoading(false));
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
      {showModal && (
        <Modal show={showModal} onClose={() => setShowModal(false)}>
          <Modal.Header>Create new broadcast</Modal.Header>
          <Modal.Body>
            <form className="pb-32 space-y-4">
              <div className="mb-2 block">
                <Label htmlFor="description" value="Description" />
                <Textarea
                  id="description"
                  placeholder="Input description"
                  className="mt-1 input-white"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                />
              </div>
              {!useTemplate && (
                <div className="mb-2 block">
                  <Label htmlFor="message" value="Message" />

                  <MentionsInput
                    value={message}
                    onChange={(val) => {
                      setMessage(val.target.value);
                    }}
                    style={emojiStyle}
                    placeholder={
                      "Press ':' for emojis, and template using '@' and shift+enter to send"
                    }
                    autoFocus
                  >
                    <Mention
                      trigger="@"
                      data={[
                        { id: "{{user}}", display: "Full Name" },
                        { id: "{{phone}}", display: "Phone Number" },
                      ]}
                      style={{
                        backgroundColor: "#cee4e5",
                      }}
                      appendSpaceOnAdd
                    />
                    <Mention
                      trigger=":"
                      markup="__id__"
                      regex={neverMatchingRegex}
                      data={queryEmojis}
                    />
                  </MentionsInput>
                </div>
              )}

              <div>
                <ToggleSwitch
                  checked={useTemplate}
                  onChange={(e) => {
                    setUseTemplate(e);
                    if (e) {
                      setMessage("");
                    }
                  }}
                  label="Use Template"
                />
              </div>
              {useTemplate && (
                <div>
                  <Label>Template</Label>
                  <Select
                    options={templates.map((t: any) => ({
                      value: t.id,
                      label: t.title,
                    }))}
                    value={{
                      value: selectedTemplate?.id,
                      label: selectedTemplate?.title,
                    }}
                    onChange={(val) =>
                      setSelectedTemplate(
                        templates.find((t: any) => t.id === val?.value)
                      )
                    }
                  />
                </div>
              )}
            </form>
          </Modal.Body>
          <Modal.Footer>
            <div className="flex w-full justify-end">
              <Button type="submit" color="success" onClick={send}>
                Save
              </Button>
            </div>
          </Modal.Footer>
        </Modal>
      )}
    </AdminLayout>
  );
};
export default BroadcastPage;

const emojiStyle = {
  control: {
    fontSize: 16,
    lineHeight: 1.2,
    minHeight: 160,
  },

  highlighter: {
    padding: 9,
    border: "1px solid transparent",
  },

  input: {
    fontSize: 16,
    lineHeight: 1.2,
    padding: 9,
    border: "1px solid silver",
    borderRadius: 10,
  },

  suggestions: {
    list: {
      backgroundColor: "white",
      border: "1px solid rgba(0,0,0,0.15)",
      fontSize: 16,
    },

    item: {
      padding: "5px 15px",
      borderBottom: "1px solid rgba(0,0,0,0.15)",

      "&focused": {
        backgroundColor: "#cee4e5",
      },
    },
  },
};
