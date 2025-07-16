import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { Button, Dropdown } from "flowbite-react";
import { PiExport } from "react-icons/pi";
import { useNavigate, useParams } from "react-router-dom";
import { InstagramMessageSessionModel } from "../models/instagram";
import { LOCAL_STORAGE_DEFAULT_INSTAGRAM_SESSION } from "../utils/constants";
import { asyncStorage } from "../utils/async_storage";
import { getContrastColor, getPagination } from "../utils/helper";
import Moment from "react-moment";
import { PaginationResponse } from "../objects/pagination";
import { deleteInstagramSession, getInstagramSessions } from "../services/api/instagramApi";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import InstagramMessages from "../components/InstagramMessages";
import toast from "react-hot-toast";

interface InstagramPageProps {}

const InstagramPage: FC<InstagramPageProps> = ({}) => {
    const { profile, setProfile } = useContext(ProfileContext);
  const [downloadModal, setDownloadModal] = useState(false);
  const nav = useNavigate();
  const [sessions, setSessions] = useState<InstagramMessageSessionModel[]>([]);
  const { sessionId } = useParams();
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
      useContext(WebsocketContext);

  useEffect(() => {
    setMounted(true);
    //   getConnections({ page: 1, size: 50 }).then((resp: any) => {
    //     setConnections(resp.data);
    //   });
  }, []);


  useEffect(() => {
    if (
      wsMsg?.command == "INSTAGRAM_RECEIVED" ||
      wsMsg?.command == "UPDATE_SESSION" ||
      wsMsg?.command == "INSTAGRAM_MESSAGE_READ" ||
      wsMsg?.command == "INSTAGRAM_CLEAR_MESSAGE"
    ) {
      getAllSessions();
    }
  }, [wsMsg, profile, sessionId]);


  useEffect(() => {
    if (mounted) {
      getAllSessions();
    }
  }, [mounted, sessionId, page, size, search]);
  const getAllSessions = () => {
    getInstagramSessions(sessionId ?? "", { page, size, search }).then(
      (resp: any) => {
        setSessions(resp.data.items);
        setPagination(getPagination(resp.data));
      }
    );
  };

  return (
    <AdminLayout>
      <div className="p-4 flex flex-col h-full ">
        <div className="flex justify-between items-center mb-2 border-b pb-4">
          <h1 className="text-3xl font-bold ">Instagram</h1>
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
            {/* <Button
                  pill
                  className="flex items-center gap-2"
                  color="gray"
                  onClick={() => {
                    setDrawerFilter(true);
                  }}
                >
                  <LuFilter className="" />
                  <span>Filter</span>
                </Button> */}
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
                  <div className="flex justify-between w-full items-center">
                    <div
                      className="flex gap-2 items-center"
                      onClick={() => {
                        nav(`/instagram/${e.id}`);
                        asyncStorage.setItem(
                          LOCAL_STORAGE_DEFAULT_INSTAGRAM_SESSION,
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
                          {e.contact?.avatar && (
                            <img
                              src={e.contact?.avatar.url}
                              className=" aspect-square rounded-full object-cover w-8 h-8"
                            />
                          )}
                          <span className="font-semibold">
                            {e.contact?.name}
                          </span>
                        </div>
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
                          {/* <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {
                              if (
                                window.confirm(
                                  "Are you sure you want to clear this session?"
                                )
                              ) {
                                // clearWhatsappSession(e.id!).then(() => {
                                //   toast.success("Chat cleared");
                                //   getAllSessions();
                                // });
                              }
                            }}
                          >
                            Clear Chat
                          </Dropdown.Item> */}
                          <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {
                              if (
                                window.confirm(
                                  "Are you sure you want to delete this session?"
                                )
                              ) {
                                deleteInstagramSession(e.id!).then(() => {
                                  toast.success("Instagram session deleted");
                                  getAllSessions();
                                  window.location.href = "/instagram";
                                });
                              }
                            }}
                          >
                            Delete Chat
                          </Dropdown.Item>
                          <Dropdown.Item
                            className="flex gap-2"
                            onClick={() => {
                              //   setSelectedSession(e);
                              //   setModalInfo(true);
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
            </ul>
          </div>
          <div className="w-full border-l relative">
            {sessionId && <InstagramMessages sessionId={sessionId} />}
          </div>
        </div>
      </div>
    </AdminLayout>
  );
};
export default InstagramPage;
