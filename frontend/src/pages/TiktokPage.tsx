import { useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { useNavigate, useParams } from "react-router-dom";
import TiktokMessages from "../components/TiktokMessages";
import { PaginationResponse } from "../objects/pagination";
import { ConnectionModel } from "../models/connection";
import { getConnections } from "../services/api/connectionApi";
import Select, { InputActionMeta } from "react-select";
import { getTiktokSessions } from "../services/api/tiktokApi";
import { TiktokMessageSession } from "../models/tiktok";
import { asyncStorage } from "../utils/async_storage";
import { LOCAL_STORAGE_DEFAULT_TIKTOK_SESSION } from "../utils/constants";
import { getContrastColor, initial } from "../utils/helper";
import { Avatar } from "flowbite-react";
import Moment from "react-moment";
import moment from "moment";
interface TiktokPageProps {}

const TiktokPage: FC<TiktokPageProps> = ({}) => {
  const { sessionId } = useParams();
  const [sessions, setSessions] = useState<TiktokMessageSession[]>([]);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const [selectedConnection, setSelectedConnection] =
    useState<ConnectionModel>();
  const [selectedFilterConnection, setSelectedFilterConnection] =
    useState<ConnectionModel>();
  const [connections, setConnections] = useState<ConnectionModel[]>([]);
  const nav = useNavigate();
  useEffect(() => {
    setMounted(true);
    getConnections({ page: 1, size: 50, type: "tiktok" }).then((resp: any) => {
      setConnections(resp.data);
      if (resp.data.length > 0) {
        setSelectedConnection(resp.data[0]);
      }
    });
  }, []);

  useEffect(() => {
    if (!selectedConnection) return;
    getTiktokSessions(sessionId || "", {
      page: page,
      size: size,
      search: search,
      tag_ids: "",
      connection_session: selectedConnection?.id,
    }).then((resp: any) => {
      setSessions(resp.data.conversations);
      setPagination(resp.data);
    });
    return () => {};
  }, [selectedConnection]);

  const sessionCard = (e: TiktokMessageSession) => {
    let contact = e.participants.find((p) => p.role == "BUYER");
    let content = JSON.parse(e.latest_message.content)
    return (
      <div className="flex justify-between w-full items-start">
        <div
          className="flex gap-2 items-center"
          onClick={() => {
            nav(`/tikrok/${e.id}`);
            asyncStorage.setItem(LOCAL_STORAGE_DEFAULT_TIKTOK_SESSION, e.id);
          }}
        >
          <div className="flex flex-col">
            {/* <div className="flex flex-wrap gap-2">
              <div
                className="flex text-[8pt] text-white  px-2 rounded-full w-fit"
                style={
                  {
                    // background: e.ref?.color,
                    // color: getContrastColor(e.ref?.color),
                  }
                }
              >
                {contact?.nickname}
              </div>
            </div> */}
            <div className="flex gap-1">
              <div className="flex flex-row gap-2 items-start">
                <Avatar
                  size="md"
                  img={contact?.avatar}
                  rounded
                  stacked
                  placeholderInitials={initial(contact?.nickname)}
                  className="cursor-pointer mt-2"
                />
                <div className="w-3/4">
                  <span className="font-semibold">{contact?.nickname}</span>
                  <small className="line-clamp-2 overflow-hidden text-ellipsis">
                    {content.content}
                  </small>
                  <small className="line-clamp-2 overflow-hidden text-ellipsis text-[8pt]">
                    {/* {moment(e.latest_message.create_time * 1000).format("DD/MM/YYYY")} */}
                    <Moment fromNow>{e.latest_message.create_time * 1000}</Moment>
                  </small>
                  {/* <div className="flex flex-wrap gap-2">
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
                </div> */}
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="flex flex-col items-end">
          {/* {(e.count_unread ?? 0) > 0 && (
          <div
            className=" aspect-square w-4 text-xs h-4  rounded-full flex justify-center items-center bg-red-400 text-white"
            color="red"
          >
            {e.count_unread}
          </div>
        )} */}
          <div className="group/edit invisible group-hover/item:visible">
            {/* <Dropdown label="" inline>
            <Dropdown.Item
              className="flex gap-2"
              onClick={() => {
                if (window.confirm(
                  "Are you sure you want to clear this session?"
                )) {
                  clearTelegramSession(e.id!).then(() => {
                    toast.success("Chat cleared");
                    getAllSessions();
                  });
                }
              } }
            >
              Clear Chat
            </Dropdown.Item>
            <Dropdown.Item
              className="flex gap-2"
              onClick={() => {
                if (window.confirm(
                  "Are you sure you want to delete this session?"
                )) {
                  deleteTelegramSession(e.id!).then(() => {
                    toast.success("Chat deleted");
                    getAllSessions();
                    window.location.href = "/telegram";
                  });
                }
              } }
            >
              Delete Chat
            </Dropdown.Item>
            <Dropdown.Item
              className="flex gap-2"
              onClick={() => {
                setSelectedSession(e);
                setModalInfo(true);
              } }
            >
              Info
            </Dropdown.Item>
          </Dropdown> */}
          </div>
        </div>
      </div>
    );
  };

  return (
    <AdminLayout>
      <div className="p-4 flex flex-col h-full ">
        <div className="flex justify-between items-center mb-2 border-b pb-4">
          <h1 className="text-3xl font-bold ">Tiktok Shop</h1>
          <Select
            formatOptionLabel={(option: ConnectionModel) => (
              <div className="flex flex-row gap-2  items-center">
                <span>{option.name}</span>
              </div>
            )}
            options={connections.filter(
              (e) => e.type === "tiktok" && e.status === "ACTIVE"
            )}
            value={selectedConnection}
            onChange={(e) => setSelectedConnection(e!)}
          />
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
                  {sessionCard(e)}
                </li>
              ))}
            </ul>
          </div>
          <div className="w-full border-l relative">
            {sessionId && <TiktokMessages sessionId={sessionId} />}
          </div>
        </div>
      </div>
    </AdminLayout>
  );
};
export default TiktokPage;
