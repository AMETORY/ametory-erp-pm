import { useContext, useEffect, useRef, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Badge,
  Button,
  Dropdown,
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
import { BsPlusCircle, BsTag, BsFileEarmark } from "react-icons/bs";
import { getProducts } from "../services/api/productApi";
import { ProductModel } from "../models/product";
import { FileModel } from "../models/file";
import { uploadFile } from "../services/api/commonApi";
import MessageMention from "../components/MessageMention";
import { TbTemplate } from "react-icons/tb";

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
  const [modaProduct, setModalProduct] = useState(false);
  const [products, setProducts] = useState<ProductModel[]>([]);
  const [selectedProducts, setSelectedProducts] = useState<ProductModel[]>([]);
  const [files, setFiles] = useState<FileModel[]>([]);
  const fileRef = useRef<HTMLInputElement>(null);
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
          <Table striped>
            <Table.Head>
              <Table.HeadCell>ID</Table.HeadCell>
              <Table.HeadCell>Description</Table.HeadCell>
              <Table.HeadCell>Message</Table.HeadCell>
              <Table.HeadCell>Status</Table.HeadCell>
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
                    {broadcast?.template_id ? (
                      <div className="flex flex-row gap-2">
                        {" "}
                        <TbTemplate />
                        {broadcast?.template?.description}
                      </div>
                    ) : (
                      parseMentions(broadcast.message, () => {})
                    )}
                  </Table.Cell>
                  <Table.Cell>
                    <div className="w-fit">
                      <Badge
                        color={
                          broadcast?.status === "DRAFT"
                            ? "warning"
                            : broadcast?.status === "EXPIRED"
                            ? "red"
                            : broadcast?.status === "STOPPED"
                            ? "yellow"
                            : broadcast?.status === "NOT_STARTED"
                            ? "yellow"
                            : broadcast?.status === "PROCESSING"
                            ? "blue"
                            : broadcast?.status === "SCHEDULED"
                            ? "purple"
                            : broadcast?.status === "FAILED"
                            ? "danger"
                            : "success"
                        }
                      >
                        {broadcast?.status}
                      </Badge>
                    </div>
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
                <div className="mb-2 block position">
                  <Label htmlFor="message" value="Message" />
                  <MessageMention
                    msg={message}
                    onChange={(val: any) => {
                      setMessage(val.target.value);
                    }}
                    onClickEmoji={() => {
                      // setSelectedMessage(template?.messages?.[i]);
                    }}
                    onSelectEmoji={(emoji: string) => {
                      setMessage(message + emoji);
                    }}
                  />

                  {/* <div className="absolute bottom-2 right-2 z-50">
                    <Dropdown
                      label={<BsPlusCircle />}
                      inline
                      placement="top"
                      arrowIcon={false}
                    >
                      <Dropdown.Item
                        className="flex gap-2"
                        onClick={() => {
                          fileRef.current?.click();
                        }}
                        icon={BsFileEarmark}
                      >
                        File
                      </Dropdown.Item>
                      <Dropdown.Item
                        className="flex gap-2"
                        onClick={() => {
                          getProducts({ page: 1, size: 10 }).then(
                            (res: any) => {
                              setProducts(res.data.items);
                            }
                          );
                          setModalProduct(true);
                        }}
                        icon={BsTag}
                      >
                        Product
                      </Dropdown.Item>
                    </Dropdown>
                  </div> */}
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
      <input
        multiple
        accept=".png, .jpg, .jpeg, .doc, .docx, .xls, .xlsx, .pdf"
        type="file"
        name="file"
        id=""
        ref={fileRef}
        className="hidden"
        onChange={async (e) => {
          if ((e.target.files ?? []).length > 0) {
            for (
              let index = 0;
              index < (e.target.files ?? []).length;
              index++
            ) {
              const element = (e.target.files ?? [])[index];
              let resp: any = await uploadFile(element, {}, console.log);
              setFiles((prev) => [...prev, resp.data]);
            }
          }
        }}
      />
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
