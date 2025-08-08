import { useContext, useEffect, useRef, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Avatar,
  Badge,
  Button,
  Datepicker,
  Drawer,
  Dropdown,
  Label,
  Modal,
  Textarea,
  TextInput,
  Tooltip,
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
import { PaginationRequest, PaginationResponse } from "../objects/pagination";
import { getContrastColor, getPagination, initial } from "../utils/helper";
import {
  clearWhatsappSession,
  deleteWhatsappSession,
  exportXls,
  getWhatsappSessionDetail,
  getWhatsappSessions,
  updateWhatsappSession,
} from "../services/api/whatsappApi";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import toast from "react-hot-toast";
import { LoadingContext } from "../contexts/LoadingContext";
import WhatsappMessages from "../components/WhatsappMessages";
import {
  getContacts,
  sendContactMessage,
  updateContact,
} from "../services/api/contactApi";
import { ContactModel } from "../models/contact";
import { FaDownload, FaMagnifyingGlass, FaXmark } from "react-icons/fa6";
import { ConnectionModel } from "../models/connection";
import { getConnections } from "../services/api/connectionApi";
import Select, { InputActionMeta } from "react-select";
import Moment from "react-moment";
import { LuDownload, LuFilter } from "react-icons/lu";
import { TagModel } from "../models/tag";
import { getTags } from "../services/api/tagApi";
import ModalSession from "../components/ModalSession";
import { getMembers } from "../services/api/commonApi";
import { MemberModel } from "../models/member";
import moment from "moment";
import { PiExport } from "react-icons/pi";
import { SearchContext } from "../contexts/SearchContext";
interface WhatsappPageProps {}

const WhatsappPage: FC<WhatsappPageProps> = ({}) => {
  const timeout = useRef<number | null>(null);
  const [members, setMembers] = useState<MemberModel[]>([]);
  const { loading, setLoading } = useContext(LoadingContext);
  const [selectedMembers, setSelectedMembers] = useState<
    { value: string; label: string }[]
  >([]);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [messages, setMessages] = useState<WhatsappMessageModel[]>([]);
  const [sessions, setSessions] = useState<WhatsappMessageSessionModel[]>([]);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
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
  const [tags, setTags] = useState<TagModel[]>([]);
  const [selectedSession, setSelectedSession] =
    useState<WhatsappMessageSessionModel>();
  const [modalInfo, setModalInfo] = useState(false);
  const [downloadModal, setDownloadModal] = useState(false);
  const [modalDateOpen, setModalDateOpen] = useState(false);
  const { search, setSearch } = useContext(SearchContext);
  const [sessionTypeSelected, setSessionTypeSelected] = useState("");
  const [filters, setFilters] = useState([
    {
      label: "All Chat",
      value: "all",
    },
    {
      label: "Unreplied",
      value: "unreplied",
    },
    {
      label: "Unread",
      value: "unread",
    },
  ]);

  const sessionType = [
    { label: "All", value: "" },
    { label: "Group", value: "group" },
    { label: "Personal", value: "personal" },
  ];

  const today = new Date();
  const start = new Date(
    today.getFullYear(),
    today.getMonth(),
    today.getDate()
  );
  const end = new Date(
    today.getFullYear(),
    today.getMonth(),
    today.getDate(),
    23,
    59,
    59
  );
  const [dateRange, setDateRange] = useState([start, end]);

  const [selectedFilters, setSelectedFilters] = useState<
    { value: string; label: string }[]
  >([]);

  const [selectedTags, setSelectedTags] = useState<TagModel[]>([]);
  const [selectedDownloadTags, setSelectedDownloadTags] = useState<TagModel[]>(
    []
  );
  const [selectedDownloadConnections, setSelectedDownloadConnections] =
    useState<ConnectionModel[]>([]);
  const [selectedConnection, setSelectedConnection] =
    useState<ConnectionModel>();
  const [selectedFilterConnection, setSelectedFilterConnection] =
    useState<ConnectionModel>();
  const [drawerFilter, setDrawerFilter] = useState(false);
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
  }, [
    mounted,
    sessionId,
    size,
    search,
    selectedTags,
    selectedFilters,
    selectedFilterConnection,
    sessionTypeSelected,
  ]);

  useEffect(() => {
    if (mounted) {
      getAllTags();
    }
  }, [mounted]);
  useEffect(() => {
    if (mounted) {
      getMembers({ page: 1, size: 10 })
        .then((res: any) => {
          setMembers(res.data.items);
        })
        .catch(toast.error);
    }
  }, [mounted]);
  const getAllTags = async () => {
    try {
      setLoading(true);
      let resp: any = await getTags({ page: 1, size: 100 });
      setTags(resp.data.items);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (timeout.current) {
      window.clearTimeout(timeout.current);
    }

    timeout.current = window.setTimeout(() => {
      if (wsMsg?.command === "WHATSAPP_CLEAR_MESSAGE") {
        getAllSessions();
      } else if (
        wsMsg?.command === "WHATSAPP_RECEIVED" ||
        wsMsg?.command === "WHATSAPP_MESSAGE_READ" ||
        wsMsg?.command === "UPDATE_SESSION"
      ) {
        if (wsMsg?.session_id) {
          getWhatsappSessionDetail(wsMsg?.session_id).then((resp: any) => {
            // console.log("SESSION DATA", resp.data);
            if (sessions.length === 0) {
              setSessions([resp.data]);
              return;
            }

            let sessionData: WhatsappMessageSessionModel[] = sessions.map(
              (s) => {
                if (s.id === resp.data.id) {
                  return resp.data;
                }
                return s;
              }
            );
            sessionData.sort(
              (a, b) =>
                moment(b.last_online_at ?? new Date()).unix() -
                moment(a.last_online_at ?? new Date()).unix()
            );

            // for (const element of sessionData) {
            //   console.log(
            //     element.contact?.name,
            //     element.contact?.phone,
            //     element.last_message
            //   );
            // }
            setSessions(sessionData);
          });
        } else {
          getAllSessions();
        }
      }
    }, 500);
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

  const getAllSessions = async (p?: number) => {
    try {
      // console.log("unread", selectedFilters.some((f) => f.value == "unread") && null);
      // console.log("unreplied", selectedFilters.some((f) => f.value == "unreplied") && null);
      // setLoading(true);
      let params: PaginationRequest = {
        page: p ?? (search == "" ? 1 : page),
        size,
        search,
        tag_ids: selectedTags.map((t) => t.id).join(","),
        is_unread: selectedFilters.some((f) => f.value == "unread")
          ? true
          : null,
        is_unreplied: selectedFilters.some((f) => f.value == "unreplied")
          ? true
          : null,
      };
      if (sessionTypeSelected != "") {
        params["type"] = sessionTypeSelected;
      }
      if (selectedFilterConnection) {
        params["connection_session"] = selectedFilterConnection.session_name;
      }

      const resp: any = await getWhatsappSessions(sessionId ?? "", params);
      setSessions(p ? [...sessions, ...resp.data.items] : resp.data.items);
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
              pill
              className="flex items-center gap-2"
              color="gray"
              onClick={() => {
                setDownloadModal(true);
              }}
            >
              <PiExport className="" />
              <span>Export</span>
            </Button>
            <Button
              pill
              className="flex items-center gap-2"
              color="gray"
              onClick={() => {
                setDrawerFilter(true);
              }}
            >
              <LuFilter className="" />
              <span>Filter</span>
            </Button>
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
          <div
            className="w-[300px]"
            style={{
              height: "calc(100vh - 160px)",
              overflowY: "auto",
            }}
          >
            <ul className="">
              {sessions.map((e) => (
                <li
                  className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold group/item border-b last:border-b-0"
                  key={e.id}
                  style={{ background: sessionId == e.id ? "#e5e7eb" : "" }}
                >
                  <div className="flex justify-between w-full items-start">
                    <div
                      className="flex gap-2 items-center  w-full"
                      onClick={() => {
                        nav(`/whatsapp/${e.id}`);
                        asyncStorage.setItem(
                          LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION,
                          e.id
                        );
                      }}
                    >
                      <div className="flex flex-col">
                        <div className="flex flex-wrap gap-2 items-center">
                          <div className="flex flex-row gap-2 items-center justify-center">
                            <small>
                              {e.ref?.connected ? (
                                <Tooltip content="Connected">🟢</Tooltip>
                              ) : (
                                <Tooltip content="Not Connected">🔴</Tooltip>
                              )}
                            </small>
                            <div
                              className="flex text-[8pt] text-white  px-2 rounded-full w-fit"
                              style={{
                                background: e.ref?.color,
                                color: getContrastColor(e.ref?.color),
                              }}
                            >
                              {e.ref?.name}
                            </div>
                          </div>
                        </div>
                        <div className="flex gap-1">
                          <div className="flex flex-row gap-2 items-start">
                            <Avatar
                              size="md"
                              img={e.contact?.profile_picture?.url}
                              rounded
                              stacked
                              placeholderInitials={initial(e.contact?.name)}
                              className="cursor-pointer mt-2"
                              // onClick={() => nav("/profile")}
                            />
                            <div className="w-3/4">
                              <div className="flex flex-col">
                                <span className="font-semibold">
                                  {e.contact?.name}
                                </span>
                                <small>({e.contact?.phone})</small>
                                <small className="line-clamp-2 overflow-hidden text-ellipsis">
                                  {e.last_message}
                                </small>
                              </div>
                              <small className="line-clamp-2 overflow-hidden text-ellipsis text-[8pt]">
                                <Moment fromNow>{e.last_online_at}</Moment>
                              </small>
                              <div className="flex flex-wrap gap-2">
                                {(e.contact?.tags ?? []).map((el) => (
                                  <div
                                    className="flex text-[8pt] text-white  px-2 rounded-full w-fit"
                                    key={el.id}
                                    style={{
                                      background: el.color,
                                      color: getContrastColor(el.color),
                                    }}
                                  >
                                    {el.name}
                                  </div>
                                ))}
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div className="flex flex-col items-end">
                      {(e.count_unread ?? 0) > 0 && (
                        <div
                          className=" aspect-square w-4 text-xs h-4  rounded-full flex justify-center items-center bg-red-400 text-white"
                          color="red"
                        >
                          {(e.count_unread ?? 0) > 99 ? "99+" : e.count_unread}
                        </div>
                      )}
                      <div className="group/edit invisible group-hover/item:visible">
                        <Dropdown label="" inline>
                          <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {
                              if (
                                window.confirm(
                                  "Are you sure you want to clear this session?"
                                )
                              ) {
                                clearWhatsappSession(e.id!).then(() => {
                                  toast.success("Chat cleared");
                                  getAllSessions();
                                });
                              }
                            }}
                          >
                            Clear Chat
                          </Dropdown.Item>
                          <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {
                              if (
                                window.confirm(
                                  "Are you sure you want to delete this session?"
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
                          <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {
                              console.log(e);
                              setSelectedSession(e);
                              setModalInfo(true);
                            }}
                          >
                            Info
                          </Dropdown.Item>
                        </Dropdown>
                      </div>
                    </div>
                  </div>
                </li>
              ))}
              {!pagination?.last && (
                <div className="w-full flex justify-center p-4">
                  <button
                    className="btn btn-primary hover:bg-gray-200 px-4 py-2 rounded-lg"
                    onClick={() => {
                      getAllSessions((pagination?.page ?? 0) + 1);
                      setPage((pagination?.page ?? 0) + 1);
                    }}
                  >
                    Load More
                  </button>
                </div>
              )}
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
                options={connections.filter(
                  (e) => e.type === "whatsapp" && e.status === "ACTIVE"
                )}
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
      <Drawer
        open={drawerFilter}
        onClose={function (): void {
          setDrawerFilter(false);
        }}
        position="right"
        style={{ width: "400px" }}
      >
        <Drawer.Header>Filter</Drawer.Header>
        <Drawer.Items>
          <div className="mt-8 flex flex-col space-y-4">
            <div>
              <Label htmlFor="name" value="Filter" />
              <Select
                isMulti
                value={selectedFilters}
                options={filters}
                onChange={(e) => {
                  setSelectedFilters(
                    e.map((e) => e.value).includes("all") ? [] : e.map((e) => e)
                  );
                }}
              />
            </div>
            <div>
              <Label htmlFor="connection" value="Connection" />
              <Select
                formatOptionLabel={(option) => (
                  <div className=" flex flex-row justify-between">
                    <p className="text-sm font-medium text-gray-900 truncate">
                      {option.label}
                    </p>
                    {option.value == "all" ? (
                      <FaXmark />
                    ) : (
                      <small> {option.session}</small>
                    )}
                  </div>
                )}
                options={[
                  ...connections.filter(
                    (e) => e.type === "whatsapp" && e.status === "ACTIVE"
                  ),
                  { id: "all", name: "Clear", session: "" },
                ].map((e) => ({
                  value: e.id!,
                  label: e.name!,
                  session: e.session_name!,
                }))}
                value={{
                  value: selectedFilterConnection?.id!,
                  label: selectedFilterConnection?.name!,
                  session: selectedFilterConnection?.session_name!,
                }}
                onChange={(e) => {
                  if (e!.value === "all") {
                    setSelectedFilterConnection(undefined);
                  } else {
                    setSelectedFilterConnection(
                      connections.find((c) => c.id === e!.value!)
                    );
                  }
                }}
              />
            </div>
            <div>
              <Label htmlFor="name" value="Session Type" />
              <Select
                value={sessionType.find((e) => e.value === sessionTypeSelected)}
                options={sessionType}
                onChange={(e) => {
                  setSessionTypeSelected(e!.value);
                }}
              />
            </div>
            <div>
              <Label htmlFor="name" value="Tag" />
              <Select
                isMulti
                value={selectedTags.map((tag) => ({
                  value: tag.id,
                  label: tag.name,
                  color: tag.color,
                }))}
                options={tags.map((tag) => ({
                  value: tag.id,
                  label: tag.name,
                  color: tag.color,
                }))}
                onChange={(e) => {
                  setSelectedTags(
                    e.map((tag) => ({
                      id: tag.value,
                      name: tag.label,
                      color: tag.color,
                    }))
                  );
                }}
                formatOptionLabel={(option) => (
                  <div
                    className="w-fit px-2 py-1 rounded-lg"
                    style={{
                      backgroundColor: option.color,
                      color: getContrastColor(option.color),
                    }}
                  >
                    <span>{option.label}</span>
                  </div>
                )}
              />
            </div>
            <div>
              <Label htmlFor="Total List" value="List Contact" />
              <Select
                value={{ value: `${size}`, label: `${size}` }}
                options={[
                  { value: "20", label: "20" },
                  { value: "50", label: "50" },
                  { value: "100", label: "100" },
                  { value: "200", label: "200" },
                  { value: "500", label: "500" },
                ]}
                onChange={(e) => {
                  setSize(parseInt(e!.value));
                }}
              />
            </div>
          </div>
        </Drawer.Items>
      </Drawer>
      <ModalSession
        type="whatsapp"
        show={modalInfo}
        onClose={() => setModalInfo(false)}
        onSave={async (val) => {
          console.log(val);
          try {
            await updateContact(val?.contact!.id!, val?.contact);
            await updateWhatsappSession(val?.id!, val);
            getAllSessions();
            setModalInfo(false);
          } catch (error) {}
        }}
        session={selectedSession}
      />
      <Modal show={downloadModal} onClose={() => setDownloadModal(false)}>
        <Modal.Header>
          <h3 className="text-lg font-semibold">Export</h3>
        </Modal.Header>
        <Modal.Body>
          <div className="flex flex-col pb-32 space-y-4">
            <div>
              <Label>Connections</Label>
              <Select
                options={connections.filter(
                  (item: any) =>
                    item.status === "ACTIVE" && item.type == "whatsapp"
                )}
                value={selectedDownloadConnections}
                isMulti
                onChange={(val) => {
                  setSelectedDownloadConnections(val.map((e) => e));
                }}
                formatOptionLabel={(option) => (
                  <div className="flex items-center space-x-3">
                    <div className="flex-1 min-w-0">
                      <p className="text-sm font-medium text-gray-900 truncate">
                        {option.name}
                      </p>
                      <p className="text-xs text-gray-500 truncate">
                        {option.session_name}
                      </p>
                    </div>
                  </div>
                )}
              />
            </div>
            {/* <div>
              <Label>Member</Label>
              <Select
                className="w-full"
                isMulti
                placeholder="Select Members"
                options={members.map((member) => ({
                  label: member.user?.full_name ?? "",
                  value: member.id ?? "",
                }))}
                value={selectedMembers}
                onChange={(selectedOptions) => {
                  setSelectedMembers(selectedOptions.map((e) => e));
                  // setSelectedMembers(selectedOptions.map((option) => ({ id: option.value })));
                }}
              />
            </div> */}
            <div>
              <Label className=" font-semibold">Date Range</Label>
              <div
                className="p-2 bg-white rounded-lg min-w-[240px] cursor-pointer border border-gray-400"
                onClick={() => setModalDateOpen(true)}
              >
                {moment(dateRange[0]).format("DD MMM YYYY")}{" "}
                {moment(dateRange[0]).format("HH:mm")} -{" "}
                {moment(dateRange[1]).format("DD MMM YYYY")}{" "}
                {moment(dateRange[1]).format("HH:mm")}
              </div>
            </div>
            <div>
              <Label className=" font-semibold">Tag</Label>
              <Select
                isMulti
                value={selectedDownloadTags.map((tag) => ({
                  value: tag.id,
                  label: tag.name,
                  color: tag.color,
                }))}
                options={tags.map((tag) => ({
                  value: tag.id,
                  label: tag.name,
                  color: tag.color,
                }))}
                onChange={(e) => {
                  setSelectedDownloadTags(
                    e.map((tag) => ({
                      id: tag.value,
                      name: tag.label,
                      color: tag.color,
                    }))
                  );
                }}
                formatOptionLabel={(option) => (
                  <div
                    className="w-fit px-2 py-1 rounded-lg"
                    style={{
                      backgroundColor: option.color,
                      color: getContrastColor(option.color),
                    }}
                  >
                    <span>{option.label}</span>
                  </div>
                )}
              />
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="w-full flex justify-end">
            <Button
              onClick={async () => {
                try {
                  // if (selectedDownloadConnections.length === 0) {
                  //   toast.error("Please select at least one connection");
                  //   return;
                  // }
                  setLoading(true);
                  var data = {
                    start_date: dateRange[0],
                    end_date: dateRange[1],
                    member_ids: selectedMembers.map((e) => e.value),
                    tag_ids: selectedDownloadTags.map((e) => e.id),
                    sessions: selectedDownloadConnections.map((e) => e.session),
                  };

                  const response: any = await exportXls(data);
                  const blob = new Blob([response], {
                    type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
                  });
                  const url = window.URL.createObjectURL(blob);
                  const link = document.createElement("a");
                  link.href = url;
                  link.download = `whatsapp_${new Date().toISOString()}.xlsx`;
                  link.style.display = "none";
                  document.body.appendChild(link);
                  link.click();
                  setTimeout(() => {
                    document.body.removeChild(link);
                    window.URL.revokeObjectURL(url);
                  }, 0);
                  setDownloadModal(false);
                  setSelectedTags([]);
                } catch (error) {
                  toast.error(`${error}`);
                } finally {
                  setLoading(false);
                }
              }}
            >
              <PiExport className="mr-1" /> Export Now
            </Button>
          </div>
        </Modal.Footer>
      </Modal>
      <Modal
        size="4xl"
        show={modalDateOpen}
        onClose={() => setModalDateOpen(false)}
        dismissible
      >
        <Modal.Header>Date Range</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col pb-32">
            <div className="grid grid-cols-2 gap-2 ">
              <div className="flex daterange">
                <Datepicker
                  value={dateRange[0]}
                  onChange={(v) => setDateRange([v!, dateRange[1]])}
                  className="min-w-[200px]"
                />
                <div className="flex w-full">
                  <input
                    type="time"
                    id="time"
                    className="rounded-none rounded-s-lg bg-gray-50 border text-gray-900 leading-none focus:ring-blue-500 focus:border-blue-500 block flex-1 w-full text-sm border-gray-300 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={moment(dateRange[0]).format("HH:mm")}
                    onChange={(e) => {
                      const newDate = new Date(dateRange[0]);
                      newDate.setHours(parseInt(e.target.value.split(":")[0]));
                      newDate.setMinutes(
                        parseInt(e.target.value.split(":")[1])
                      );
                      setDateRange([newDate, dateRange[1]]);
                    }}
                  />
                  <span className="inline-flex  items-center px-3 text-sm text-gray-900 bg-gray-200  rounded-s-0 border-0 border-gray-300 rounded-e-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                    <svg
                      className="w-4 h-4 text-gray-500 dark:text-gray-400"
                      aria-hidden="true"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        fillRule="evenodd"
                        d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10-4.477 10-10 10S2 17.523 2 12Zm11-4a1 1 0 1 0-2 0v4a1 1 0 0 0 .293.707l3 3a1 1 0 0 0 1.414-1.414L13 11.586V8Z"
                        clipRule="evenodd"
                      />
                    </svg>
                  </span>
                </div>
              </div>
              <div className="flex daterange">
                <Datepicker
                  value={dateRange[1]}
                  onChange={(v) => setDateRange([dateRange[0], v!])}
                  className="min-w-[200px]"
                />
                <div className="flex w-full">
                  <input
                    type="time"
                    id="time"
                    className="rounded-none rounded-s-lg bg-gray-50 border text-gray-900 leading-none focus:ring-blue-500 focus:border-blue-500 block flex-1 w-full text-sm border-gray-300 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={moment(dateRange[1]).format("HH:mm")}
                    onChange={(e) => {
                      const newDate = new Date(dateRange[1]);
                      newDate.setHours(parseInt(e.target.value.split(":")[0]));
                      newDate.setMinutes(
                        parseInt(e.target.value.split(":")[1])
                      );
                      setDateRange([dateRange[0], newDate]);
                    }}
                  />
                  <span className="inline-flex items-center px-3 text-sm text-gray-900 bg-gray-200  rounded-s-0 border-0 border-gray-300 rounded-e-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                    <svg
                      className="w-4 h-4 text-gray-500 dark:text-gray-400"
                      aria-hidden="true"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        fillRule="evenodd"
                        d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10-4.477 10-10 10S2 17.523 2 12Zm11-4a1 1 0 1 0-2 0v4a1 1 0 0 0 .293.707l3 3a1 1 0 0 0 1.414-1.414L13 11.586V8Z"
                        clipRule="evenodd"
                      />
                    </svg>
                  </span>
                </div>
              </div>
            </div>
            <div className="mt-4">
              <ul className="grid grid-cols-2 gap-4">
                <li>
                  <button
                    className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                    onClick={() => {
                      const start = new Date();
                      start.setHours(0, 0, 0, 0);
                      const end = new Date();
                      end.setHours(23, 59, 59, 999);
                      setDateRange([start, end]);
                    }}
                  >
                    Today
                  </button>
                </li>
                <li>
                  <button
                    className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                    onClick={() => {
                      const yesterday = new Date();
                      yesterday.setDate(yesterday.getDate() - 1);
                      yesterday.setHours(0, 0, 0, 0);
                      const end = new Date();
                      end.setDate(end.getDate() - 1);
                      end.setHours(23, 59, 59, 999);
                      setDateRange([yesterday, end]);
                    }}
                  >
                    Yesterday
                  </button>
                </li>
                <li>
                  <button
                    className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                    onClick={() => {
                      const lastWeek = new Date();
                      lastWeek.setDate(lastWeek.getDate() - 7);
                      lastWeek.setHours(0, 0, 0, 0);
                      const end = new Date();
                      end.setHours(23, 59, 59, 999);
                      setDateRange([lastWeek, end]);
                    }}
                  >
                    Last Week
                  </button>
                </li>
                <li>
                  <button
                    className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                    onClick={() => {
                      const firstDayOfMonth = new Date(
                        new Date().getFullYear(),
                        new Date().getMonth(),
                        1
                      );
                      firstDayOfMonth.setHours(0, 0, 0, 0);
                      const end = moment(firstDayOfMonth)
                        .endOf("month")
                        .toDate();
                      end.setHours(23, 59, 59, 999);
                      setDateRange([firstDayOfMonth, end]);
                    }}
                  >
                    This Month
                  </button>
                </li>
                {[...Array(4)].map((_, i) => {
                  const quarter = i + 1;
                  const start = new Date(
                    new Date().getFullYear(),
                    (quarter - 1) * 3,
                    1
                  );
                  const end = new Date(
                    new Date().getFullYear(),
                    quarter * 3,
                    0
                  );
                  return (
                    <li key={i}>
                      <button
                        key={i}
                        className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                        onClick={() => {
                          const endCopy = new Date(end);
                          endCopy.setHours(23, 59, 59, 999);
                          setDateRange([start, endCopy]);
                        }}
                      >
                        Q{quarter}
                      </button>
                    </li>
                  );
                })}

                {[...Array(4)].map((_, i) => {
                  const year = new Date().getFullYear() - (i === 0 ? 0 : i);
                  return (
                    <li key={i}>
                      <button
                        key={i}
                        className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                        onClick={() => {
                          const endCopy = new Date(year, 11, 31);
                          endCopy.setHours(23, 59, 59, 999);
                          setDateRange([new Date(year, 0, 1), endCopy]);
                        }}
                      >
                        {year == new Date().getFullYear() ? "This Year" : year}
                      </button>
                    </li>
                  );
                })}
              </ul>
            </div>
          </div>
        </Modal.Body>
      </Modal>
    </AdminLayout>
  );
};
export default WhatsappPage;
