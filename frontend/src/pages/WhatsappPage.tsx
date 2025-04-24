import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Badge,
  Button,
  Drawer,
  Dropdown,
  Label,
  Modal,
  Textarea,
  TextInput,
} from "flowbite-react";
import {
  WhatsappMessageModel,
  WhatsappMessageSessionModel,
} from "../models/whatsapp_message";
import {
  LOCAL_STORAGE_DEFAULT_CHANNEL,
  LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION,
} from "../utils/constants";
import { useNavigate, useParams } from "react-router-dom";
import { asyncStorage } from "../utils/async_storage";
import { PaginationResponse } from "../objects/pagination";
import { getPagination } from "../utils/helper";
import {
  deleteWhatsappSession,
  getWhatsappSessions,
} from "../services/api/whatsappApi";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import toast from "react-hot-toast";
import { LoadingContext } from "../contexts/LoadingContext";
import WhatsappMessages from "../components/WhatsappMessages";
import { getContacts, sendContactMessage } from "../services/api/contactApi";
import { ContactModel } from "../models/contact";
import { FaMagnifyingGlass } from "react-icons/fa6";
import { ConnectionModel } from "../models/connection";
import { getConnections } from "../services/api/connectionApi";
import Select, { InputActionMeta } from "react-select";
interface WhatsappPageProps {}

const WhatsappPage: FC<WhatsappPageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);

  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [messages, setMessages] = useState<WhatsappMessageModel[]>([]);
  const [sessions, setSessions] = useState<WhatsappMessageSessionModel[]>([]);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const { sessionId } = useParams();
  const nav = useNavigate();
  const [openChannelForm, setOpenChannelForm] = useState(false);
  const [contacts, setContacts] = useState<ContactModel[]>([]);
  const [searchContact, setSearchContact] = useState("");
  const [modalSendMessage, setModalSendMessage] = useState(false);
  const [selectedContact, setSelectedContact] = useState<ContactModel>();
  const [message, setMessage] = useState("");
  const [connections, setConnections] = useState<ConnectionModel[]>([]);
  const [selectedConnection, setSelectedConnection] =
    useState<ConnectionModel>();

  useEffect(() => {
    setMounted(true);
    getConnections({ page: 1, size: 50 }).then((resp: any) => {
      setConnections(resp.data);
    });
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllSessions();
    }
  }, [mounted, sessionId]);

  useEffect(() => {
    if (wsMsg?.command == "WHATSAPP_RECEIVED" || wsMsg?.command == "UPDATE_SESSION") {
      getAllSessions();
    }
  }, [wsMsg, profile, sessionId]);

  const getAllContact = async (s: string) => {
    try {
      const resp: any = await getContacts({
        page: 1,
        size: 30,
        search: s,
      });
      setContacts(resp.data.items);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
    }
  };

  const getAllSessions = async () => {
    try {
      // setLoading(true);
      const resp: any = await getWhatsappSessions(sessionId ?? "", {
        page,
        size,
        search,
      });
      setSessions(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      // setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-4 flex flex-col h-full ">
        <div className="flex justify-between items-center mb-2 border-b pb-4">
          <h1 className="text-3xl font-bold ">Whatsapp</h1>
          <div className="flex gap-2">
            <Button
              gradientDuoTone="purpleToBlue"
              pill
              onClick={() => {
                setOpenChannelForm(true);
                getAllContact("");
              }}
            >
              + Chat
            </Button>
          </div>
        </div>
        <div className="flex flex-row w-full h-full flex-1 gap-2">
          <div className="w-[300px] h-full">
            <ul className="space-y-2">
              {sessions.map((e) => (
                <li
                  className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold group/item"
                  key={e.id}
                  style={{ background: sessionId == e.id ? "#e5e7eb" : "" }}
                >
                  <div className="flex justify-between w-full items-center">
                    <div
                      className="flex gap-2 items-center"
                      onClick={() => {
                        nav(`/whatsapp/${e.id}`);
                        asyncStorage.setItem(
                          LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION,
                          e.id
                        );
                      }}
                    >
                      {e.contact?.avatar && (
                        <img
                          src={e.contact?.avatar.url}
                          className=" aspect-square rounded-full object-cover w-8 h-8"
                        />
                      )}
                      <div className="flex flex-col">
                        <span className="font-semibold">{e.contact?.name}</span>
                        <small className="line-clamp-2 overflow-hidden text-ellipsis">
                          {e.last_message}
                        </small>
                      </div>
                    </div>
                    <div className="flex flex-col items-end">
                      {(e.count_unread ?? 0) > 0 && (
                        <div  className=" aspect-square w-4 text-xs h-4  rounded-full flex justify-center items-center bg-red-400 text-white"  color="red">{e.count_unread}</div>
                      )}
                      <div className="group/edit invisible group-hover/item:visible">
                        <Dropdown label="" inline>
                          <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {}}
                          >
                            Clear Chat
                          </Dropdown.Item>
                          <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {
                              if (
                                window.confirm(
                                  "Are you sure you want to delete this chat?"
                                )
                              ) {
                                deleteWhatsappSession(e.id!).then(() => {
                                  toast.success("Chat deleted");
                                  getAllSessions();
                                  window.location.href = "/whatsapp";
                                });
                              }
                            }}
                          >
                            Delete Chat
                          </Dropdown.Item>
                        </Dropdown>
                      </div>
                    </div>
                  </div>
                </li>
              ))}
            </ul>
          </div>
          <div className="w-full border-l relative">
            {sessionId && <WhatsappMessages sessionId={sessionId} />}
          </div>
        </div>
      </div>
      <Drawer
        position="right"
        open={openChannelForm}
        onClose={() => setOpenChannelForm(false)}
      >
        <div className="mt-16">
          <div className="relative">
            <TextInput
              placeholder="Search"
              value={searchContact}
              onChange={(e) => {
                getAllContact(e.target.value);
                setSearchContact(e.target.value);
              }}
              className="w-full"
            />
            <FaMagnifyingGlass className=" absolute top-3 right-2" />
          </div>

          <ul>
            {contacts.map((e) => (
              <li
                className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold"
                key={e.id}
                onClick={() => {
                  setModalSendMessage(true);
                  setSelectedContact(e);
                }}
                style={{ background: sessionId == e.id ? "#e5e7eb" : "" }}
              >
                <div className="flex gap-2 items-center">{e.name}</div>
                <small>{e.phone}</small>
              </li>
            ))}
          </ul>
        </div>
      </Drawer>
      <Modal show={modalSendMessage} onClose={() => setModalSendMessage(false)}>
        <Modal.Header>Send Message</Modal.Header>
        <Modal.Body>
          <div className="flex gap-2 space-y-4 flex-col w-full">
            <div className="w-full flex flex-col">
              <Label className=" font-semibold">Connection</Label>
              <Select
                formatOptionLabel={(option: ConnectionModel) => (
                  <div className="flex flex-row gap-2  items-center">
                    <span>{option.name}</span>
                  </div>
                )}
                options={connections}
                value={selectedConnection}
                onChange={(e) => setSelectedConnection(e!)}
              />
            </div>
            <div className="w-full flex flex-col">
              <Label className=" font-semibold">To</Label>
              <div>{selectedContact?.name}</div>
              <small>{selectedContact?.phone}</small>
            </div>
            <div className="w-full">
              <Label className=" font-semibold">Message</Label>
              <Textarea
                rows={7}
                placeholder="Message"
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                className="w-full"
              />
            </div>
          </div>
          <div className="flex gap-2"></div>
        </Modal.Body>
        <Modal.Footer>
          <Button
            gradientDuoTone="purpleToBlue"
            onClick={() => {
              setLoading(true);
              sendContactMessage(selectedContact!.id!, {
                message,
                type: "whatsapp",
                connection_id: selectedConnection!.id,
              })
                .then(() => {
                  setSelectedConnection(undefined);
                  setModalSendMessage(false);
                  setMessage("");
                  setSelectedContact(undefined);
                })
                .catch((err) => {
                  toast.error(`${err}`);
                })
                .finally(() => {
                  setLoading(false);
                });
            }}
          >
            Send Message
          </Button>
        </Modal.Footer>
      </Modal>
    </AdminLayout>
  );
};
export default WhatsappPage;
