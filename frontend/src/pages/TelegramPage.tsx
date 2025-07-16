import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { Avatar, Button, Drawer, Dropdown, Label } from "flowbite-react";
import { PiExport } from "react-icons/pi";
import { useNavigate, useParams } from "react-router-dom";
import { TelegramMessageSessionModel } from "../models/telegram";
import { LOCAL_STORAGE_DEFAULT_TELEGRAM_SESSION } from "../utils/constants";
import { asyncStorage } from "../utils/async_storage";
import { getContrastColor, getPagination, initial } from "../utils/helper";
import Moment from "react-moment";
import { PaginationRequest, PaginationResponse } from "../objects/pagination";
import {
  clearTelegramSession,
  deleteTelegramSession,
  getTelegramSessions,
} from "../services/api/telegramApi";
import TelegramMessages from "../components/TelegramMessages";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import ModalSession from "../components/ModalSession";
import { updateContact } from "../services/api/contactApi";
import { updateWhatsappSession } from "../services/api/whatsappApi";
import { LuFilter } from "react-icons/lu";
import Select, { InputActionMeta } from "react-select";
import { FaXmark } from "react-icons/fa6";
import { ConnectionModel } from "../models/connection";
import { TagModel } from "../models/tag";
import { getConnections } from "../services/api/connectionApi";
import { getTags } from "../services/api/tagApi";
import toast from "react-hot-toast";
import { LoadingContext } from "../contexts/LoadingContext";
interface TelegramPageProps {}

const TelegramPage: FC<TelegramPageProps> = ({}) => {
  const { profile, setProfile } = useContext(ProfileContext);
  const [downloadModal, setDownloadModal] = useState(false);
  const { loading, setLoading } = useContext(LoadingContext);
  const nav = useNavigate();
  const [sessions, setSessions] = useState<TelegramMessageSessionModel[]>([]);
  const { sessionId } = useParams();
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const [modalInfo, setModalInfo] = useState(false);
  const [drawerFilter, setDrawerFilter] = useState(false);
  const [selectedSession, setSelectedSession] =
    useState<TelegramMessageSessionModel>();
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const [tags, setTags] = useState<TagModel[]>([]);
  const [selectedConnection, setSelectedConnection] =
    useState<ConnectionModel>();
  const [selectedFilterConnection, setSelectedFilterConnection] =
    useState<ConnectionModel>();
  const [connections, setConnections] = useState<ConnectionModel[]>([]);
  const [selectedTags, setSelectedTags] = useState<TagModel[]>([]);
  useEffect(() => {
    setMounted(true);
    //   getConnections({ page: 1, size: 50 }).then((resp: any) => {
    //     setConnections(resp.data);
    //   });
  }, []);

  useEffect(() => {
    setMounted(true);
    getConnections({ page: 1, size: 50 }).then((resp: any) => {
      setConnections(resp.data);
    });
  }, []);

  useEffect(() => {
    if (
      wsMsg?.command == "TELEGRAM_RECEIVED" ||
      wsMsg?.command == "UPDATE_SESSION" ||
      wsMsg?.command == "TELEGRAM_MESSAGE_READ" ||
      wsMsg?.command == "TELEGRAM_CLEAR_MESSAGE"
    ) {
      getAllSessions();
    }
  }, [wsMsg, profile, sessionId]);

  useEffect(() => {
    if (mounted) {
      getAllSessions();
    }
  }, [
    mounted,
    sessionId,
    page,
    size,
    search,
    selectedTags,
    selectedFilterConnection,
  ]);

  useEffect(() => {
    if (mounted) {
      getAllTags();
    }
  }, [mounted]);
  const getAllSessions = async (p?: number) => {
    let params: PaginationRequest = {
      page: p ?? (search == "" ? 1 : page),
      size,
      search,
      tag_ids: selectedTags.map((t) => t.id).join(","),
    };
    if (selectedFilterConnection) {
      params["connection_session"] = selectedFilterConnection.id;
    }
    try {
      const resp: any = await getTelegramSessions(sessionId ?? "", params);
      // setSessions(resp.data.items);
      setSessions(p ? [...sessions, ...resp.data.items] : resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    }
  };

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

  return (
    <AdminLayout>
      <div className="p-4 flex flex-col h-full ">
        <div className="flex justify-between items-center mb-2 border-b pb-4">
          <h1 className="text-3xl font-bold ">Telegram</h1>
          <div className="flex gap-2">
            {/* <Button
                  pill
                  className="flex items-center gap-2"
                  color="gray"
                  onClick={() => {
                    setDownloadModal(true);
                  }}
                >
                  <PiExport className="" />
                  <span>Export</span>
                </Button> */}
           
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
            {/* <Button
                  gradientDuoTone="purpleToBlue"
                  pill
                  onClick={() => {
                    setOpenChannelForm(true);
                    getAllContact("");
                  }}
                >
                  + Chat
                </Button> */}
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
            <ul className="space-y-2">
              {sessions.map((e) => (
                <li
                  className="flex justify-between items-center p-2 hover:bg-gray-50 cursor-pointer hover:font-semibold group/item"
                  key={e.id}
                  style={{ background: sessionId == e.id ? "#e5e7eb" : "" }}
                >
                  <div className="flex justify-between w-full items-start">
                    <div
                      className="flex gap-2 items-center"
                      onClick={() => {
                        nav(`/telegram/${e.id}`);
                        asyncStorage.setItem(
                          LOCAL_STORAGE_DEFAULT_TELEGRAM_SESSION,
                          e.id
                        );
                      }}
                    >
                      <div className="flex flex-col">
                        <div className="flex flex-wrap gap-2">
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
                              <span className="font-semibold">
                                {e.contact?.name}
                              </span>
                              <small className="line-clamp-2 overflow-hidden text-ellipsis">
                                {e.last_message}
                              </small>
                              <small className="line-clamp-2 overflow-hidden text-ellipsis text-[8pt]">
                                <Moment fromNow>{e.last_online_at}</Moment>
                              </small>
                              <div className="flex flex-wrap gap-2">
                                {(e.contact?.tags ?? []).map((el) => (
                                  <div
                                    className="flex text-[8pt] text-white  px-2 rounded-full w-fit"
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
                          {e.count_unread}
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
                                clearTelegramSession(e.id!).then(() => {
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
                                deleteTelegramSession(e.id!).then(() => {
                                  toast.success("Chat deleted");
                                  getAllSessions();
                                  window.location.href = "/telegram";
                                });
                              }
                            }}
                          >
                            Delete Chat
                          </Dropdown.Item>
                          <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {
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
            {sessionId && <TelegramMessages sessionId={sessionId} />}
          </div>
        </div>
      </div>
      <ModalSession
        type="telegram"
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
                    (e) => e.type === "telegram" && e.status === "ACTIVE"
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
    </AdminLayout>
  );
};
export default TelegramPage;
