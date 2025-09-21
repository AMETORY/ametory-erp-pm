import { useContext, useEffect, useRef, useState, type FC } from "react";
import { useNavigate } from "react-router-dom";
import { WhatsappMessageSessionModel } from "../models/whatsapp_message";
import { PaginationRequest, PaginationResponse } from "../objects/pagination";
import { SearchContext } from "../contexts/SearchContext";
import { TagModel } from "../models/tag";
import { ConnectionModel } from "../models/connection";
import {
  clearWhatsappSession,
  deleteWhatsappSession,
  getWhatsappSessionDetail,
  getWhatsappSessions,
  updateWhatsappSession,
} from "../services/api/whatsappApi";
import { getContrastColor, getPagination, initial } from "../utils/helper";
import toast from "react-hot-toast";
import { asyncStorage } from "../utils/async_storage";
import { LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION } from "../utils/constants";
import { Avatar, Dropdown, Tooltip } from "flowbite-react";
import { AiOutlineLoading3Quarters } from "react-icons/ai";
import Moment from "react-moment";
import { WebsocketContext } from "../contexts/WebsocketContext";
import moment from "moment";
import ModalSession from "./ModalSession";
import { Tabs } from "flowbite-react";
import { updateContact } from "../services/api/contactApi";

interface WhatsappSessionListProps {
  sessionId?: string;
  selectedTags: TagModel[];
  selectedFilters: { value: string; label: string }[];
  sessionTypeSelected: string;
  selectedFilterConnection?: ConnectionModel;
}

const WhatsappSessionList: FC<WhatsappSessionListProps> = ({
  sessionId,
  selectedTags,
  selectedFilters,
  sessionTypeSelected,
  selectedFilterConnection,
}) => {
  const { wsMsg, setWsMsg } = useContext(WebsocketContext);
  const [sessions, setSessions] = useState<WhatsappMessageSessionModel[]>([]);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(20);
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [mounted, setMounted] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const { search, setSearch } = useContext(SearchContext);
  const timeout = useRef<number | null>(null);
  const [modalInfo, setModalInfo] = useState(false);
  const [selectedSession, setSelectedSession] =
    useState<WhatsappMessageSessionModel>();
  const [activeTab, setActiveTab] = useState(0);
  const [countPersonal, setCountPersonal] = useState(0);
  const [countGroup, setCountGroup] = useState(0);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (!mounted) return;
    getAllSessions();
    return () => {};
  }, [
    mounted,
    sessionId,
    selectedTags,
    selectedFilters,
    sessionTypeSelected,
    selectedFilterConnection,
    search,
    activeTab,
  ]);

  useEffect(() => {
    if (!wsMsg) return;
    // if (timeout.current) {
    //   window.clearTimeout(timeout.current);
    // }
    if (wsMsg?.command === "WHATSAPP_CLEAR_MESSAGE") {
      getAllSessions();
    }

    if (
      wsMsg?.command === "WHATSAPP_RECEIVED" ||
      wsMsg?.command === "WHATSAPP_MESSAGE_READ" ||
      wsMsg?.command === "UPDATE_SESSION"
    ) {
      if (wsMsg?.session_id) {
        let indexSession = sessions.findIndex((s) => s.id == wsMsg.session_id);

        if (indexSession > -1) {
          let newSessions = [...sessions];
          newSessions[indexSession] = {
            ...newSessions[indexSession],
            last_online_at: wsMsg.created_at,
            last_message: wsMsg.message,
            count_unread: (newSessions[indexSession].count_unread ?? 0) + 1,
          };
          newSessions.sort(
            (a, b) =>
              moment(b.last_online_at ?? new Date()).unix() -
              moment(a.last_online_at ?? new Date()).unix()
          );
          setSessions(newSessions);
        } else {
          getWhatsappSessionDetail(wsMsg?.session_id).then((resp: any) => {
            if (sessions.length === 0) {
              setSessions([resp.data]);
              return;
            } else {
              let newSessions = [...sessions];
              newSessions.push(resp.data);
              newSessions.sort(
                (a, b) =>
                  moment(b.last_online_at ?? new Date()).unix() -
                  moment(a.last_online_at ?? new Date()).unix()
              );
              setSessions(newSessions);
            }
          });
        }
      }
    }
    return () => {};
  }, [wsMsg]);

  useEffect(() => {
    setCountPersonal(
      sessions
        ?.filter((e) => !e.is_group)
        .reduce((e, acc) => e + (acc.count_unread ?? 0), 0)
    );
    setCountGroup(
      sessions
        ?.filter((e) => e.is_group)
        .reduce((e, acc) => e + (acc.count_unread ?? 0), 0)
    );
  }, [sessions]);

  const getAllSessions = async (p?: number) => {
    try {
      setIsLoading(true);

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
      if (activeTab == 1) {
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
      setIsLoading(false);
    }
  };
  const nav = useNavigate();

  const contactSession = (e: WhatsappMessageSessionModel) => (
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
            asyncStorage.setItem(LOCAL_STORAGE_DEFAULT_WHATSAPP_SESSION, e.id);
          }}
        >
          <div className="flex flex-col">
            <div className="flex flex-wrap gap-2 items-center">
              <div className="flex flex-row gap-2 items-center justify-center">
                <small>
                  {e.ref?.connected || e.ref?.type == "whatsapp-api" ? (
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
                    <span className="font-semibold">{e.contact?.name}</span>
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
        <div>
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
      </div>
    </li>
  );

  return (
    <>
      <div className="w-[300px] relative">
        <div className="w-full flex flex-row ">
          <button
            className="w-full text-sm items-center px-1 py-0.5 rounded-lg hover:bg-gray-50 cursor-pointer hover:font-semibold group/item text-center"
            style={{
              background: activeTab == 0 ? "#e5e7eb" : "",
              color: activeTab == 0 ? "black" : "",
            }}
            onClick={() => {
              setActiveTab(0);
            }}
          >
            Personal
            {countPersonal > 0 && (<span className="ml-2 text-xs font-bold bg-red-400 h-fit w-2 px-1 rounded-full text-white">
              {countPersonal > 99 ? "99+" : countPersonal}
            </span>)}
            
          </button>
          <button
            className="w-full text-sm items-center px-1 py-0.5 rounded-lg hover:bg-gray-50 cursor-pointer hover:font-semibold group/item  text-center"
            style={{
              background: activeTab == 1 ? "#e5e7eb" : "",
              color: activeTab == 1 ? "black" : "",
            }}
            onClick={() => {
              setActiveTab(1);
            }}
          >
            Group
            {countGroup > 0 && (<span className="ml-2 text-xs font-bold bg-red-400 h-fit w-2 px-1 rounded-full text-white">
              {countGroup > 99 ? "99+" : countGroup}
            </span>)}
          </button>
        </div>
        <div
          className="w-full relative"
          style={{
            height: "calc(100vh - 220px)",
            overflowY: "auto",
          }}
        >
          {isLoading && (
            <div className="flex flex-col items-center justify-center w-full p-8 absolute top-0 left-0 right-0 bottom-0 bg-[rgba(255,255,255,0.2)]">
              <div className="animate-spin h-5 w-5 border-b-2 border-gray-900 rounded-full">
                <AiOutlineLoading3Quarters />
              </div>
            </div>
          )}
          <ul className="">
            {sessions
              .filter((e) => !e.is_group && activeTab == 0)
              .map((e) => contactSession(e))}
            {sessions
              .filter((e) => e.is_group && activeTab == 1)
              .map((e) => contactSession(e))}

            {!pagination?.last && !isLoading && (
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
      </div>
      <ModalSession
        type="whatsapp"
        show={modalInfo}
        onClose={() => setModalInfo(false)}
        onSave={async (val) => {
          // console.log(val);
          try {
            await updateContact(val?.contact!.id!, val?.contact);
            await updateWhatsappSession(val?.id!, val);
            getAllSessions();
            setModalInfo(false);
          } catch (error) {}
        }}
        session={selectedSession}
      />
    </>
  );
};
export default WhatsappSessionList;
